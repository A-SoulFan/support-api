package utility

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"net/http"
)

const (
	webInterfaceView = "https://api.bilibili.com/x/web-interface/view?bvid="
	playerPageList   = "https://api.bilibili.com/x/player/pagelist?bvid="
)

type BiliBili struct {
}

type CommonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Ttl     int         `json:"ttl"`
	Data    interface{} `json:"data"`
}

type PlayerPageListResponse struct {
	Cid       int    `json:"cid"`
	Page      int    `json:"page"`
	From      string `json:"from"`
	Part      string `json:"part"`
	Duration  int    `json:"duration"`
	Vid       string `json:"vid"`
	Weblink   string `json:"weblink"`
	Dimension struct {
		Width  int `json:"width"`
		Height int `json:"height"`
		Rotate int `json:"rotate"`
	} `json:"dimension"`
	FirstFrame string `json:"first_frame"`
}

type WebInterfaceViewResponse struct {
	BvId      string `json:"bvid"`
	Aid       int    `json:"aid"`
	Videos    int    `json:"videos"`
	Tid       int    `json:"tid"`
	TName     string `json:"tname"`
	Copyright int    `json:"copyright"`
	Pic       string `json:"pic"`
	Title     string `json:"title"`
	Pubdate   int    `json:"pubdate"`
	Ctime     int    `json:"ctime"`
	Desc      string `json:"desc"`
	DescV2    []struct {
		RawText string `json:"raw_text"`
		Type    int    `json:"type"`
		BizId   int    `json:"biz_id"`
	} `json:"desc_v2"`
	State    int `json:"state"`
	Duration int `json:"duration"`
	Rights   struct {
		Bp            int `json:"bp"`
		Elec          int `json:"elec"`
		Download      int `json:"download"`
		Movie         int `json:"movie"`
		Pay           int `json:"pay"`
		Hd5           int `json:"hd5"`
		NoReprint     int `json:"no_reprint"`
		Autoplay      int `json:"autoplay"`
		UgcPay        int `json:"ugc_pay"`
		IsCooperation int `json:"is_cooperation"`
		UgcPayPreview int `json:"ugc_pay_preview"`
		NoBackground  int `json:"no_background"`
		CleanMode     int `json:"clean_mode"`
		IsSteinGate   int `json:"is_stein_gate"`
		Is360         int `json:"is_360"`
	} `json:"rights"`
	Owner struct {
		Mid  int    `json:"mid"`
		Name string `json:"name"`
		Face string `json:"face"`
	} `json:"owner"`
	Stat struct {
		Aid        int    `json:"aid"`
		View       int    `json:"view"`
		Danmaku    int    `json:"danmaku"`
		Reply      int    `json:"reply"`
		Favorite   int    `json:"favorite"`
		Coin       int    `json:"coin"`
		Share      int    `json:"share"`
		NowRank    int    `json:"now_rank"`
		HisRank    int    `json:"his_rank"`
		Like       int    `json:"like"`
		Dislike    int    `json:"dislike"`
		Evaluation string `json:"evaluation"`
		ArgueMsg   string `json:"argue_msg"`
	} `json:"stat"`
	Dynamic   string `json:"dynamic"`
	Cid       int    `json:"cid"`
	Dimension struct {
		Width  int `json:"width"`
		Height int `json:"height"`
		Rotate int `json:"rotate"`
	} `json:"dimension"`
	NoCache bool `json:"no_cache"`
	Pages   []struct {
		Cid       int    `json:"cid"`
		Page      int    `json:"page"`
		From      string `json:"from"`
		Part      string `json:"part"`
		Duration  int    `json:"duration"`
		Vid       string `json:"vid"`
		Weblink   string `json:"weblink"`
		Dimension struct {
			Width  int `json:"width"`
			Height int `json:"height"`
			Rotate int `json:"rotate"`
		} `json:"dimension"`
		FirstFrame string `json:"first_frame"`
	} `json:"pages"`
	Subtitle struct {
		AllowSubmit bool          `json:"allow_submit"`
		List        []interface{} `json:"list"`
	} `json:"subtitle"`
	UserGarb struct {
		UrlImageAniCut string `json:"url_image_ani_cut"`
	} `json:"user_garb"`
}

func (bb *BiliBili) PlayerPageList(bvId string) ([]PlayerPageListResponse, error) {
	url := playerPageList + bvId
	data := make([]PlayerPageListResponse, 0)

	if err := fastGet(url, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func (bb *BiliBili) WebInterfaceView(bvId string) (*WebInterfaceViewResponse, error) {
	url := webInterfaceView + bvId
	data := &WebInterfaceViewResponse{}

	if err := fastGet(url, data); err != nil {
		return nil, err
	}

	return data, nil
}

func fastGet(url string, dataInterface interface{}) error {
	result := &CommonResponse{Data: &dataInterface}
	client := resty.New()
	if resp, err := client.R().SetResult(result).Get(url); err != nil {
		return err
	} else {
		if resp.StatusCode() != http.StatusOK {
			return errors.New("request fail")
		}

		if result.Code != 0 {
			return errors.New(result.Message)
		}

		return nil
	}
}
