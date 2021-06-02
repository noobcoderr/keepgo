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
	"keepgo/register"
)

func main() {
	keepGO := mqant.CreateApp(
		module.Debug(true),
		)
	err := keepGO.Run(
		register.Module(),
		)
	if err!=nil{
		log.Error(err.Error())
	}
}