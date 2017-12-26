package app

import (
	"github.com/domac/io_match/brand"
	"net"
	"os"
	"sync"
)

type Agentd struct {
	httpListener net.Listener //http监听器
	waitGroup    WaitGroupWrapper

	sync.RWMutex          //同步锁
	opts         *Options //配置参数选项

	exitChan chan int

	isExit bool //退出标识
}

//创建后台进程对象
func New(opts *Options) *Agentd {
	a := &Agentd{
		opts: opts,
	}
	return a
}

func (self *Agentd) GetOptions() *Options {
	return self.opts
}

func (self *Agentd) GetExitCh() chan int {
	return self.exitChan
}

//后台程序退出
func (self *Agentd) Exit() {
	self.opts.Logger.Warnf("agentd program is exiting ...")
	if self.httpListener != nil {
		self.httpListener.Close()
	}
	close(self.exitChan)
	self.isExit = true
	self.waitGroup.Wait()
}

//主程序入口
//Agent主要逻辑入口
func (self *Agentd) Main() {
	ctx := &Context{self}

	//http服务开关
	if self.opts.HTTPAddress != "" {
		httpListener, err := net.Listen("tcp", self.opts.HTTPAddress)
		if err != nil {
			self.opts.Logger.Errorf("listen (%s) failed - %s", self.opts.HTTPAddress, err)
			os.Exit(1)
		}
		self.Lock()
		self.httpListener = httpListener
		self.Unlock()
		//开启自身的 api 服务端
		apiServer := newAPIServer(ctx)
		//开启对外提供的api服务
		self.waitGroup.Wrap(func() {
			Serve(self.httpListener, apiServer, "HTTP", self.opts.Logger)
		})

		brand.InitKeys(self.GetOptions().Brand)
	}
}
