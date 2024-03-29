package task

import (
	"asoul-fan-support/app/service"
	"asoul-fan-support/lib/utility"
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	tempSaveFile = "temp/recommend_slice.csv"
)

type taskGetRecommendSlice struct {
	recommendSlice Videos
}

//bid:BV号
//title：标题
//image_url:封面链接
//auth:作者
//play_val:播放量
//time:时长
//url：原视频的跳转链接

type Videos []Video

func (v Videos) Len() int {
	return len(v)
}

func (v Videos) Less(i, j int) bool {
	iP, err := strconv.Atoi(v[i].PlayVal)
	if err != nil {
		panic(err)
	}

	jP, err := strconv.Atoi(v[j].PlayVal)
	if err != nil {
		panic(err)
	}

	return iP > jP
}

func (v Videos) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

type Video struct {
	Bid        string `json:"bid"`
	Title      string `json:"title"`
	Desc       string `json:"description"`
	ImageUrl   string `json:"image_url"`
	Auth       string `json:"auth"`
	PlayVal    string `json:"play_val"`
	Time       string `json:"time"`
	TimeSecond string `json:"time_second"`
	Url        string `json:"url"`
}

var (
	once sync.Once
	_tr  *taskGetRecommendSlice
)

func Hot(l int) []Video {
	temp := make([]Video, 0, l)
	for i, video := range _tr.recommendSlice {
		if i == l {
			break
		}

		temp = append(temp, video)
	}

	return temp
}

func Register(svc *service.Context) {
	once.Do(func() {
		_tr = &taskGetRecommendSlice{recommendSlice: []Video{}}

		_tr.init()
	})
}

func (tr *taskGetRecommendSlice) init() {
	tr.rearrange()
	ticker(tr)
}

func (tr *taskGetRecommendSlice) rearrange() {
	if _, err := os.Stat(tempSaveFile); err != nil {
		return
	}

	f, err := os.Open(tempSaveFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	cR := csv.NewReader(f)
	cAll, err := cR.ReadAll()

	if err != nil {
		panic(err)
	}

	temp := make([]Video, 0, len(cAll))
	for _, strings := range cAll {
		temp = append(temp, buildVideo(strings))
	}

	tr.recommendSlice = temp

	sort.Sort(tr.recommendSlice)
}

func buildVideo(csvLen []string) Video {
	video := Video{}

	for i, s := range csvLen {
		switch i {
		case 1:
			video.PlayVal = s
		case 2:
			video.ImageUrl = s
		case 3:
			video.Url = s
		case 4:
			video.Desc = s
		case 5:
			video.Title = s
		case 6:
			video.Auth = s
		case 9:
			video.Time = s
		case 10:
			video.TimeSecond = s
		case 11:
			video.Bid = s
			video.Url = fmt.Sprintf("https://www.bilibili.com/video/%s", s)
		}
	}

	return video
}

func ticker(tr *taskGetRecommendSlice) {
	tk := time.NewTicker(15 * time.Minute)

	stopChan := make(chan bool)
	go func(ticker *time.Ticker) {
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if err := search(); err != nil {
					fmt.Println(err)
					continue
				}
				tr.rearrange()
			case stop := <-stopChan:
				if stop {
					return
				}
			}
		}
	}(tk)
}

func search() error {
	var hitMap = map[string]int{}
	var mids = []string{
		"393396916", // 贾布
		"351609538", // 珈乐
		"672328094", // 嘉然
		"672342685", // 乃琳
		"672346917", // 向晚
		"672353429", // 贝拉
	}

	cvFile, err := os.Create(tempSaveFile)
	if err != nil {
		return err
	}
	defer cvFile.Close()
	_, _ = cvFile.WriteString("\xEF\xBB\xBF") // UTF-8 BOM
	cW := csv.NewWriter(cvFile)

	b := &utility.BiliBili{}
	for _, mid := range mids {
		var end = false
		for i := 1; !end; i++ {
			var res *utility.SpaceSearchResponse
			var err error
			for n := 0; n < 3; n++ {
				time.Sleep(time.Second * time.Duration(n+1))

				res, err = b.SpaceSearch(mid, 30, i)
				if err != nil {
					if n == 2 {
						fmt.Printf("读取失败, mid:%s ps:%d pn:%d \n %v \n", mid, 30, i, err)
						break
					}
				} else {
					break
				}
			}

			if len(res.List.Vlist) < 30 {
				end = true
			} else if len(res.List.Vlist) == 0 {
				end = true
				break
			}

			for _, v := range res.List.Vlist {
				lengthS := parseTime(v.Length)
				if lengthS > (15 * 60) {
					continue
				}

				// 出道日之前的
				if int64(v.Created) <= time.Date(2020, 12, 11, 0, 0, 0, 0, time.Local).Unix() {
					continue
				}

				// 14天之前的
				if int64(v.Created) <= time.Now().Add(-14*24*time.Hour).Unix() {
					continue
				}

				if _, ok := hitMap[v.Bvid]; ok {
					continue
				} else {
					hitMap[v.Bvid] = 1
				}

				err := cW.Write([]string{
					strconv.Itoa(v.Comment),
					strconv.Itoa(v.Play),
					v.Pic,
					v.Subtitle,
					v.Description,
					v.Title,
					v.Author,
					strconv.Itoa(v.Mid),
					strconv.Itoa(v.Created),
					v.Length,
					strconv.Itoa(lengthS),
					v.Bvid,
				})

				if err != nil {
					fmt.Printf("写入失败, %+v %+v \n", v, err)
				}
			}
		}
	}

	cW.Flush()

	err = cW.Error()
	if err != nil {
		return err
	}
	return nil
}

func parseTime(sTime string) int {
	l := strings.Split(sTime, ":")
	t := 0
	k := 1

	for i := len(l); i > 0; i-- {
		n, _ := strconv.Atoi(l[i-1])
		t += n * k

		k = k * 60
	}

	return t
}
