package ioc

import "github.com/gin-gonic/gin"

type ServiceObjIoc interface {
	// 用于初始化
	Init() error
	// 对象名称
	Name() string
}

type ControllerObjIoc interface {
	ServiceObjIoc
	// 路由注册的函数
	Registry(r gin.IRouter)
}
