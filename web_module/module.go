/**
 * @Author: zhangjie
 * @Description:
 * @File:  module
 * @Version: 1.0.0
 * @Date: 2021/6/3 11:49 下午
 */

package web_module

import (
	"context"
	"github.com/liangdas/mqant/conf"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
	basemodule "github.com/liangdas/mqant/module/base"
	mqrpc "github.com/liangdas/mqant/rpc"
	"io"
	"net/http"
)

var Module = func() module.Module {
	this := new(Web)
	return this
}

type Web struct {
	basemodule.BaseModule
}

func (self *Web) GetType() string {
	//很关键,需要与配置文件中的Module配置对应
	return "Web"
}
func (self *Web) Version() string {
	//可以在监控时了解代码版本
	return "1.0.0"
}
func (self *Web) OnInit(app module.App, settings *conf.ModuleSettings) {
	self.BaseModule.OnInit(self, app, settings)
}

func (self *Web) startHttpServer() *http.Server {
	srv := &http.Server{Addr: ":8080"}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm()
		rstr, err := mqrpc.String(
			self.Call(
				context.Background(),
				"regist",
				"/say/hi",
				mqrpc.Param(r.Form.Get("name")),
			),
		)
		log.Info("RpcCall %v , err %v", rstr, err)
		_, _ = io.WriteString(w, rstr)
	})
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			// cannot panic, because this probably is an intentional close
			log.Info("Httpserver: ListenAndServe() error: %s", err)
		}
	}()
	// returning reference so caller can call Shutdown()
	return srv
}

func (self *Web) Run(closeSig chan bool) {
	log.Info("web: starting HTTP server :8080")
	srv := self.startHttpServer()
	<-closeSig
	log.Info("web: stopping HTTP server")
	// now close the server gracefully ("shutdown")
	// timeout could be given instead of nil as a https://golang.org/pkg/context/
	if err := srv.Shutdown(nil); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}
	log.Info("web: done. exiting")
}

func (self *Web) OnDestroy() {
	//一定别忘了继承
	self.BaseModule.OnDestroy()
}
