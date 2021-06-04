/**
 * @Author: zhangjie
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2021/6/3 12:25 上午
 */

package main

import (
	"github.com/liangdas/mqant"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
	"github.com/liangdas/mqant/registry"
	"github.com/liangdas/mqant/registry/consul"
	"github.com/nats-io/nats.go"
	"keepgo/register"
	"keepgo/web_module"
)

func main() {
	// docker pull consul # 默认拉取latest
	// docker run -d -p 8500:8500 --restart=always --name=consul consul:latest agent -server -bootstrap -ui -node=1 -client='0.0.0.0'
	rs := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:8500"}
	})
	// docker pull nats
	// docker run -d -p 5555:4444 nats -p 4444
	nc, err := nats.Connect("nats://127.0.0.1:5555", nats.MaxReconnects(10000))
	if err != nil {
		log.Error("nats error %v", err)
		return
	}
	keepGO := mqant.CreateApp(
		module.Debug(true),
		module.Nats(nc),     //指定nats rpc
		module.Registry(rs), //指定服务发现
	)
	err = keepGO.Run(
		register.Module(),
		web_module.Module(),
	)
	if err != nil {
		log.Error(err.Error())
	}
}
