package task

import (
	"asoul-fan-support/app/service"
	"encoding/csv"
	"fmt"
	"os"
	"sync"
)

const (
	tempSaveFile = "temp/recommend_slice.csv"
)

type taskGetRecommendSlice struct {
	recommendSlice []Video
}

//bid:BV号
//title：标题
//image_url:封面链接
//auth:作者
//play_val:播放量
//time:时长
//url：原视频的跳转链接

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

func Register(svc *service.Context) {
	once.Do(func() {
		_tr = &taskGetRecommendSlice{recommendSlice: []Video{}}

		_tr.init()
	})
}

func (tr *taskGetRecommendSlice) init() {
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

	fmt.Println(tr.recommendSlice)
}

var csvLenMap = map[int]func(v *Video, val string){
	1: func(v *Video, val string) {
		v.PlayVal = val
	},
	2: func(v *Video, val string) {
		v.ImageUrl = val
	},
	3: func(v *Video, val string) {
		v.Url = val
	},
	4: func(v *Video, val string) {
		v.Desc = val
	},
	5: func(v *Video, val string) {
		v.Title = val
	},
	6: func(v *Video, val string) {
		v.Auth = val
	},
	9: func(v *Video, val string) {
		v.Time = val
	},
	10: func(v *Video, val string) {
		v.TimeSecond = val
	},
}

func buildVideo(csvLen []string) Video {
	video := Video{}

	for i, s := range csvLen {
		if call, ok := csvLenMap[i]; ok {
			call(&video, s)
		}
	}

	return video
}
