package service

import (
)

type IService interface {
	Notice(exp string) error
}

const (
	ZBi = iota
	XkCoin
	HuoBi
)

func NewIService(_ty int) IService {
	switch _ty {
	case ZBi:
		return zbi.NewZbi()
	case  XkCoin:
		return
	case HuoBi:
	}
	return nil
}
