package app

import (
	"github.com/domac/io_match/brand"
	"github.com/domac/io_match/util"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"path/filepath"
	"strconv"
)

type ApiServer struct {
	ctx    *Context     //上下文
	router http.Handler //路由
}

//HTTP 服务
func newAPIServer(ctx *Context) *ApiServer {

	log := Log(ctx.Agentd.opts.Logger)

	router := httprouter.New()
	router.HandleMethodNotAllowed = true
	router.PanicHandler = LogPanicHandler(ctx.Agentd.opts.Logger)
	router.NotFound = LogNotFoundHandler(ctx.Agentd.opts.Logger)
	router.MethodNotAllowed = LogMethodNotAllowedHandler(ctx.Agentd.opts.Logger)

	s := &ApiServer{
		ctx:    ctx,
		router: router,
	}

	//在这里注册路由服务
	router.Handle("GET", "/io_match", Decorate(s.ioHandler, log, Default)) //json格式输出
	router.Handle("GET", "/test", Decorate(s.testHandler, log, Default))   //json格式输出
	return s
}

func (s *ApiServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.router.ServeHTTP(w, req)
}

func (s *ApiServer) testHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) (interface{}, error) {
	res := NewResult(RESULT_CODE_FAIL, true, "test ok", nil)
	return res, nil
}

//IO比赛官方专用API接口
func (s *ApiServer) ioHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) (interface{}, error) {
	brand_file := s.ctx.Agentd.opts.Brand
	brand_file_path, _ := filepath.Abs(brand_file)

	if !util.IsExist(brand_file_path) {
		res := NewResult(RESULT_CODE_FAIL, false, brand_file_path, "file not found")
		return res, nil
	}

	paramReq, err := NewReqParams(req)
	if err != nil {
		return nil, err
	}

	sign, _ := paramReq.Get("sign")
	dataDisk, _ := paramReq.Get("dataDisk")
	dataCheckequenceStr, _ := paramReq.Get("dataCheckequence")
	dataCheckequence, _ := strconv.Atoi(dataCheckequenceStr)
	s.ctx.Logger().Infof("sign:%s | dataDisk:%s | dataCheckequence:%d", sign, dataDisk, dataCheckequence)

	//测试输出结果
	res := brand.HandleDiskData(sign, dataDisk, dataCheckequence)
	return res, nil
}
