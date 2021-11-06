package main

import "github.com/xiewei/notice/service"

func main()  {
	BiAn := service.NewIService(service.BiAn)
	BiAn.Notice("")
}
