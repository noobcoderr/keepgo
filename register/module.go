/**
 * @Author: zhangjie
 * @Description:
 * @File:  module
 * @Version: 1.0.0
 * @Date: 2021/6/3 12:31 上午
 */

package register

import (
	"fmt"
	"github.com/liangdas/mqant/conf"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
	basemodule "github.com/liangdas/mqant/module/base"
)

var Module = func() module.Module {
	this := new(Regist)
	return this
}

type Regist struct {
	basemodule.BaseModule
}

func (m *Regist) GetType() string {
	return "regist"
}

func (m *Regist) Version() string {
	return "0.0.1"
}

func (m *Regist) OnInit(app module.App, settings *conf.ModuleSettings) {
	m.BaseModule.OnInit(m, app, settings)
	log.Info("%v模块初始化完成...", m.GetType())
	m.GetServer().RegisterGO("/say/hi", m.say)

}

//func (m *Regist)OnAppConfigurationLoaded(app module.App) {
//
//}

func (m *Regist) Run(closeSig chan bool) {
	log.Info("%v模块运行中...", m.GetType())
	log.Info("%v say hello world...", m.GetType())
	<-closeSig
	log.Info("%v模块已停止...", m.GetType())
}

func (m *Regist) OnDestroy() {
	m.BaseModule.OnDestroy()
	log.Info("%v模块已回收...", m.GetType())
}

func (m *Regist) say(name string) (r string, err error) {
	return fmt.Sprintf("hi %v", name), nil
}
