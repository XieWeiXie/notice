package service

import (
	"github.com/xiewei/notice/service/bian"
	"github.com/xiewei/notice/service/coinbase"
	"github.com/xiewei/notice/service/huobi"
	"github.com/xiewei/notice/service/xkcoin"
	"github.com/xiewei/notice/service/zbi"
)

type IService interface {
	Notice(exp string) error
}

const (
	ZBi = iota
	XkCoin
	HuoBi
	CoinBase
	BiAn
)

func NewIService(_ty int) IService {
	switch _ty {
	case ZBi:
		return zbi.NewZbi()
	case XkCoin:
		return xkcoin.NewXKCoin()
	case HuoBi:
		return huobi.NewHuoBi()
	case CoinBase:
		return coinbase.NewCoinBase()
	case BiAn:
		return bian.NewBiAn()
	}
	return nil
}
