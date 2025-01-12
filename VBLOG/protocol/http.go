package protocol

import (
	"context"
	"github.com/accoladexin/vblog/common/logger"
	"github.com/accoladexin/vblog/conf"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 定义http server的一些属性
type Http struct {
	server *http.Server
}

func NewHttp(r *gin.Engine) *Http {
	return &Http{
		server: &http.Server{
			// 服务器监听地址
			Addr: conf.C().Http.Address(),
			// 配置http server使用的路由
			Handler: r,
			//  client ---> server
			ReadTimeout: 5 * time.Second,
			// server ---> client
			WriteTimeout: 5 * time.Second,
		},
	}
}

// 启动
func (h *Http) Start() error {
	logger.L().Debug().Msgf("http sever: %s", conf.C().Http.Address())

	// http
	return h.server.ListenAndServe()

	// https
	// h.server.ServeTLS()
}

// 停止
func (h *Http) Stop(ctx context.Context) {
	// 服务的优雅关闭, 先关闭监听,新的请求就进不来, 等待老的请求 处理完成
	// 自己介绍来自操作系统的信号量 来决定 你的服务是否需要关闭
	// nginx  reload, os reload sign,  config.reload()
	// os termnial sign, 过了terminal的超时时间会直接kill
	h.server.Shutdown(ctx)
}
