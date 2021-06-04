/**
 * @Author: zhangjie
 * @Description:基础支付模块
 * @File:  module
 * @Version: 1.0.0
 * @Date: 2021/6/4 3:32 下午
 */

package pay

import (
	"fmt"
	"github.com/liangdas/mqant/conf"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
	basemodule "github.com/liangdas/mqant/module/base"
	"time"
)

var Module = func() module.Module {
	this := new(Pay)
	return this
}

type Pay struct {
	basemodule.BaseModule
}

func (m *Pay) GetType() string {
	return "Pay"
}

func (m *Pay) Version() string {
	return "0.0.1"
}

func (m *Pay) OnInit(app module.App, settings *conf.ModuleSettings) {
	m.BaseModule.OnInit(m, app, settings)
	log.Info("%v模块初始化完成...", m.GetType())
	m.GetServer().RegisterGO("/order/create", m.createOrderId)

}

//func (m *Regist)OnAppConfigurationLoaded(app module.App) {
//
//}

func (m *Pay) Run(closeSig chan bool) {
	log.Info("%v模块运行中...", m.GetType())
	log.Info("%v get a order id...", m.GetType())
	<-closeSig
	log.Info("%v模块已停止...", m.GetType())
}

func (m *Pay) OnDestroy() {
	m.BaseModule.OnDestroy()
	log.Info("%v模块已回收...", m.GetType())
}

func (m *Pay) createOrderId(app string) (r string, err error) {
	return fmt.Sprintf("o-%v", time.Now().Unix()), nil
}
