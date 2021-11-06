package notice

import "github.com/xiewei/notice/notice/wx"

type Noticer interface {
	Do(v interface{}) error
}

const (
	WxNoticeType = 1
)

func NewNoticer(_type int) Noticer {
	switch _type {
	case WxNoticeType:
		return wx.NoticeWX{}
	}
	return nil
}
