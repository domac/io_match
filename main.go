package main

import (
	"flag"
	"github.com/domac/io_match/app"
	"github.com/judwhite/go-svc/svc"
	"github.com/mreiferson/go-options"
	"log"
	"os"
	"path/filepath"
	"runtime/debug"
	"syscall"
)

var (
	flagSet     = flag.NewFlagSet("IOMATCH", flag.ExitOnError)
	brand       = flagSet.String("brand", "/home/apps/brand_name.txt", "path to brand file")
	data        = flagSet.String("data", "", "path to data file")
	httpAddress = flagSet.String("http_address", ":20711", "<addr>:<port> to listen on for HTTP clients")
)

//程序封装
type program struct {
	Agentd *app.Agentd
}

//框架初始化
func (p *program) Init(env svc.Environment) error {
	if env.IsWindowsService() {
		//切换工作目录
		dir := filepath.Dir(os.Args[0])
		return os.Chdir(dir)
	}
	return nil
}

//程序启动
func (p *program) Start() error {

	flagSet.Parse(os.Args[1:])

	var cfg map[string]interface{}
	opts := app.NewOptions()
	options.Resolve(opts, flagSet, cfg)
	//后台进程创建
	daemon := app.New(opts)
	daemon.Main()
	p.Agentd = daemon
	return nil
}

//程序停止
func (p *program) Stop() error {
	if p.Agentd != nil {
		p.Agentd.Exit()
	}
	return nil
}

//引导程序
func main() {
	debug.SetGCPercent(20)
	prg := &program{}
	if err := svc.Run(prg, syscall.SIGINT, syscall.SIGTERM); err != nil {
		log.Fatal(err)
	}
}
