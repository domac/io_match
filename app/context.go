package app

import (
	"github.com/domac/io_match/log"
)

//上下文环境
type Context struct {
	Agentd *Agentd
}

func (c *Context) Logger() log.Logger {
	return c.Agentd.opts.Logger
}
