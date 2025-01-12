package api

import (
	"github.com/accoladexin/vblog/apps/blog"
	"github.com/accoladexin/vblog/ioc"
	"github.com/gin-gonic/gin"
)

// 用来检查是否实现了ioc的接口
var _ ioc.ControllerObjIoc = &Handler{}

// 处理Http报文的处理器
// 又这个handler类来负责实现具体的API 接口
type Handler struct {
	// 需要一个业务的具体实现 所以继承了Service的接口特点
	svc blog.Service
}

// 为啥要用一个函数来构造
// 初始化一些基础数据
// 具体传递实现，在什么时候决定
// 典型的面向接口编程
// 使用ioc 替换掉 直接依赖
func NewHandler() *Handler {
	return &Handler{}
}

// 构造函数带有具体的对象
func NewHandlerWithObj(svc blog.Service) *Handler {
	//svcobj := ioc.GetServiceIocByName(blog.AppName)
	// //断言
	//service := svcobj.(blog.Service)
	//if service == nil {
	//	panic("service is nil")
	//}

	return &Handler{
		svc: svc, // 非ioc版本
	}
}

func (h *Handler) Name() string {
	// 使用不同的存储来保存对象 controller map, api map
	return blog.AppName
}

// 把你的请求和Http 路由对应上
// r 就是一个gin http 的路由器
// URL:  /vblog/api/v1/blogs --->  那个Handler(处理函数)
func (h *Handler) Registry(r gin.IRouter) {
	r.GET("", h.QueryBlog)
	r.GET("/:id", h.DescribeBlog)
	r.POST("", h.CreateBlog)
	r.PUT("/:id", h.UpdateBlog)
	r.DELETE("/:id", h.DeleteBlog)
}

// import _ "gitee.com/go-course/go11/vblog/apps/blog/api"
// 怎么控制加载哪些业务单元
func init() {
	//ioc.RegistryHttpApi(&Handler{})
	// 将serice对象注册到ioc容器
	//fmt.Println("========================================================")
	ioc.RegisterControllerIoc(&Handler{})
}

// 让他从Ioc中获取依赖的对象
func (h *Handler) Init() error {
	// 获取出来的这个对象是个Ioc对象
	// 使用类型断言 来把any --> 正在的业务对象
	// 内置逻辑: blog.AppName 名称 --> blog.Service 类型  object.(type struct<object>|interface)
	//  var a = any   a.(string)
	// svc = ioc.GetController(blog.AppName)
	// svc.(blog.Service)
	h.svc = ioc.GetServiceIocByName(blog.AppName).(blog.Service)
	return nil
}
