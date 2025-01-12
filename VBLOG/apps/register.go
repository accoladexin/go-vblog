package apps

// 本包主要用于实现环境的初始加载

import (
	// 加载service层的ioc通过init()
	_ "github.com/accoladexin/vblog/apps/blog/api"
	// 加载controller层的ioc通过init()
	_ "github.com/accoladexin/vblog/apps/blog/impl"
)
