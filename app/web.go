package app

import (
	"fmt"
	lg "github.com/domac/io_match/log"
	"log"
	"net"
	"net/http"
	"strings"
)

type logWriter struct {
	lg.Logger
}

func (l logWriter) Write(p []byte) (int, error) {
	return len(p), nil

}

//简单清爽的Http服务
func Serve(listener net.Listener, handler http.Handler, proto string, l lg.Logger) {
	l.Infof("[%s]starting %s http listen", proto, listener.Addr())
	server := &http.Server{
		Handler:  handler,
		ErrorLog: log.New(logWriter{}, "", 0)}

	err := server.Serve(listener)
	if err != nil && !strings.Contains(err.Error(), "use of closed network connection") {
		l.Errorf(fmt.Sprintf("ERROR: http.Serve() - %s", err))
	}
	l.Warnf(fmt.Sprintf("%s: closing %s", proto, listener.Addr()))
}
