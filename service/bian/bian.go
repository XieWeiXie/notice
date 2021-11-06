package bian

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
	"github.com/xiewei/notice/notice"
	"github.com/xiewei/notice/notice/wx"
	"log"
	"net/http"
	"time"
)

var (
	announcement = "https://www.binance.com/zh-CN/support/announcement/c-48?navId=48"
	prefix = "https://www.binance.com/zh-CN/support/announcement/"
)

type BiAn struct {

}

func (b BiAn) Notice(message string) (err error){
	res, err := http.Get(announcement)
	if err != nil {
		log.Println(fmt.Sprintf("http.Get fail, err: %v", err))
		return err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println(fmt.Sprintf("goquery.NewDocumentFromReader fail, err: %v", err))
		return err
	}

	json := doc.Find("#__APP_DATA").Text()
	jsonData := gjson.Parse(json)
	for _, one := range jsonData.Get("routeProps.b723.navDataResource").Array() {
		for _ , i := range one.Get("articles").Array() {
			var article = new(Catalog)
			article.Title = i.Get("title").String()
			article.ReleaseDate = i.Get("releaseDate").Int()
			article.Date = time.Unix(article.ReleaseDate/1000, 0)
			article.Link = fmt.Sprintf(prefix + "/%s", i.Get("code").String())
			fmt.Println(article)
			now := time.Now()
			if article.Date.Month() != now.Month() {
				continue
			}
			noticer := notice.NewNoticer(notice.WxNoticeType)
			var params = make([]wx.TemplateData, 0)
			params = append(params, wx.TemplateData{
				Value: "Binance",
				Color: "#173177",
			})
			params = append(params, wx.TemplateData{
				Value: article.Date.String(),
				Color: "#173177",
			})
			params = append(params, wx.TemplateData{
				Value: article.Title,
				Color: "#173177",
			})
			if err = noticer.Do(params); err != nil {
				log.Println("noticer", err)
			}
		}
	}
	return
}


type Catalog struct {
	Title string `json:"title"`
	ReleaseDate int64 `json:"releaseDate"`
	Date time.Time `json:"date"`
	Link string `json:"link"`
}

func NewBiAn() BiAn {
	return BiAn{}
}