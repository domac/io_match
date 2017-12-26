package app

import (
	"github.com/domac/io_match/log"
)

//配置选项
type Options struct {
	// 基本参数
	HTTPAddress string `flag:"http_address"`
	Brand       string `flag:"brand"`
	Logger      log.Logger
}

func NewOptions() *Options {
	return &Options{
		Logger: log.DefaultLogger,
	}
}
