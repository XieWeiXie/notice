package wx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
)

const (
	appId = "wx7fc2506259e3247f"
	secret = "f3923ac4e0b560ef79bfdeea8c30956f"
)

const (
	token = " https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	templateId = "fM-h6Z9HO7XBg_GfzHjzQpC2piH_h9ZyuHslcvu33WQ"
	openId = "o7vFp5po9jmILt4u4DUOo_5nFsPA"
	localToken = "50_KvtSoYwOFSKDY1kX_JIqfQsKxeBFra7ZPz7ytupim5dds3mXQ17qGwMvBrdD-VqBjVsA0qfaxdElGR1J5drOPDObwtCC5c4wgB5SPVwrqQXbSG0FC7wPqsoc8VrA7O_xwAO_xTfAQG1OAMozHSXhAGAOZE"
	templatePost = "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s"
)

type NoticeWX struct {

}

func (n NoticeWX) Do(v interface{}) error {
	params := v.([]TemplateData)
	data := TemplateNotice{
		ToUser:     openId,
		TemplateId: templateId,
		Data: struct {
			Platform TemplateData `json:"platform"`
			Date     TemplateData `json:"date"`
			Title    TemplateData `json:"title"`
		}{
			Platform: params[0],
			Date: params[1],
			Title: params[2],
		},
	}
	by, _ := json.Marshal(data)
	reader := bytes.NewReader(by)
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf(templatePost, localToken), reader)
	if err != nil {
		return err
	}
	response, err := http.DefaultClient.Do(request)
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
	return err
}

func (n NoticeWX) token() string {
	res, err := http.Get(fmt.Sprintf(token, appId, secret))
	if err != nil {
		return ""
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err.Error()
	}
	return gjson.Parse(doc.Text()).Get("access_token").String()
}

type TemplateData struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type TemplateNotice struct {
	ToUser string `json:"touser"`
	TemplateId string `json:"template_id"`
	Data struct{
		Platform TemplateData `json:"platform"`
		Date TemplateData `json:"date"`
		Title TemplateData `json:"title"`
	}`json:"data"`
}