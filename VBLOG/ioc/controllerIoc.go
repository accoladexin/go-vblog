package ioc

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
)

var controllerIocList = map[string]ControllerObjIoc{}

// 初始化ioc
// 遍历controllerIoc, 调用初始化方法，注册路由
func InitController(pathPrefix string, r gin.IRouter) error {
	for k, v := range controllerIocList {
		err := v.Init() // 初始化方法，一般是v通过ioc的方式将自己的service从ioc给拿出来
		if err != nil {
			return fmt.Errorf("init %s, error %s\n", k, err)
		}
		// 注册
		// 每一模块生成一个子路由
		// 子路由的Prefix 如何决定? "/vblog/api/v1/blogs"
		// "/vblog/api/v1" + "blogs"
		//  project prefix +  app prefix
		// path.Join
		fmt.Printf("register:==========> %s\n", path.Join(pathPrefix, k))
		// 创建一个path.Join(pathPrefix, k)未开头的组
		v.Registry(r.Group(path.Join(pathPrefix, k)))
	}
	return nil
}

// 打印下当前已经托管的实例的名称
func ShowControllerIoc() (names []string) {
	for k := range controllerIocList {
		names = append(names, k)
	}
	return names
}

// 托管业务实现的类
func RegisterControllerIoc(obj ControllerObjIoc) {
	controllerIocList[obj.Name()] = obj
}

func GetControllerIoc(name string) ControllerObjIoc {
	return controllerIocList[name]
}
