业务的实际逻辑类似与java的controller
```go
// 使用web框架: Gin
func (h *Handler) CreateBlog(c *gin.Context) {
	// 从web框架中 获取用户的请求,请求参数放到Body
	// 原始的处理方式
	// io.ReadAll(c.Request.Body)
	// gin body --> json.Umarshal

	// 获取用户信息
	// 中间件的逻辑:
	// tkObj, err := usvc.ValiateToken(c.Request.Context(), user.NewValiateTokenRequest(tk))
	// if err != nil {
	// 	common.SendFaild(c, err)
	// 	return
	// }
	// // 3. 把User对象 放到Ctx中(请求的上下文), 后面CreateBlog需要从上下文中获取用户的身份信息
	// c.Set(user.REQUEST_CTX_USER_KEY, tkObj)
	// 注意这里设置的 *user.Token
	v, ok := c.Get(user.REQUEST_CTX_TOKEN_KEY)
	if !ok {
		common.SendFaild(c, common.NewUnauthorized())
		return
	}

	u := v.(*user.Token)
	logger.L().Debug().Msgf("login user: %s", u)

	// 初始化一个请求对象
	in := blog.NewCreateBlogRequest()
	err := c.BindJSON(in)
	if err != nil {
		// 如何同意处理异常？ 我们的接口如何进行统一处理:
		// 正常应该返回什么数据?
		// 异常应该返回什么数据?
		common.SendFaild(c, err)
		return
	}

	in.Author = u.Username

	// 处理用户的请求
	// 调用业务逻辑成来处理业务逻辑
	// err := controllers.CreateBlog(ins)

	ins, err := h.svc.CreateBlog(c.Request.Context(), in)
	if err != nil {
		// common.SendFaild(c, http.StatusBadRequest, http.StatusBadRequest, err.Error())
		common.SendFaild(c, err)
		return
	}

	// 把业务逻辑处理后的结果返回给用户
	c.JSON(http.StatusOK, ins)
}

// 使用web框架: Gin
func (h *Handler) QueryBlog(c *gin.Context) {
	// 认证用户
	// checkUser(c)
	// ...

	// 模拟网络延迟，前端最Loading处理
	// time.Sleep(3 * time.Second)

	// 初始化一个请求对象
	in := blog.NewQueryBlogRequest()

	// 请求的参数 就是HTTP QueryString, URL参数
	// page_size=1234&page_number=1&author=&keywords=
	ps := c.Query("page_size")
	if ps != "" {
		in.PageSize, _ = strconv.Atoi(c.Query("page_size"))
	}
	pn := c.Query("page_number")
	if pn != "" {
		in.PageNumber, _ = strconv.Atoi(c.Query("page_number"))
	}

	in.Author = c.Query("author")
	in.Keywords = c.Query("keywords")

	// 处理用户的请求
	// 调用业务逻辑成来处理业务逻辑
	// err := controllers.CreateBlog(ins)

	ins, err := h.svc.QueryBlog(c.Request.Context(), in)
	if err != nil {
		// common.SendFaild(c, http.StatusBadRequest, http.StatusBadRequest, err.Error())
		common.SendFaild(c, err)
		return
	}

	// 把业务逻辑处理后的结果返回给用户
	c.JSON(http.StatusOK, ins)
}

// 使用web框架: Gin
// func (h *Handler) CreateBlog(c *gin.Context) {
// 	// 从web框架中 获取用户的请求

// 	// 用户参数可能通过Http Header传递过来
// 	c.GetHeader("CROS_TOKEN")
// 	// URL路径参数: /blogs/:id  --> /blogs/12  id=12
// 	c.Param("id")
// 	// URL query参数
// 	// GET /path?id=1234&name=Manu&value=
// 	c.Query("id")
// 	// ...
// 	// HTTP Boby 的参数 {"title": "文章1", "content": "xxxx"}
// 	in := &blog.CreateBlogRequest{}
// 	// json.Unmarshal  json --> object
// 	c.BindJSON(in)

// 	// 处理用户的请求
// 	// 调用业务逻辑成来处理业务逻辑
// 	// err := controllers.CreateBlog(ins)

// 	ins, err := h.svc.CreateBlog(c.Request.Context(), in)
// 	if err != nil {
// 		c.AbortWithError(http.StatusOK, err)
// 		return
// 	}

// 	// 把业务逻辑处理后的结果返回给用户
// 	c.JSON(http.StatusOK, ins)
// }

func (h *Handler) DescribeBlog(c *gin.Context) {
	req := &blog.DescribeBlogRequest{
		Id: c.Param("id"),
	}
	ins, err := h.svc.DescribeBlog(c.Request.Context(), req)
	if err != nil {
		common.SendFaild(c, err)
		return
	}

	c.JSON(http.StatusOK, ins)
}

func (h *Handler) UpdateBlog(c *gin.Context) {
	req := blog.NewUpdateBlogRequest(c.Param("id"))
	// 保持作者
	v, ok := c.Get(user.REQUEST_CTX_TOKEN_KEY)
	if !ok {
		common.SendFaild(c, common.NewUnauthorized())
		return
	}

	u := v.(*user.Token)
	req.Author = u.Username

	err := c.BindJSON(req.CreateBlogRequest)
	if err != nil {
		common.SendFaild(c, err)
		return
	}

	ins, err := h.svc.UpdateBlog(c.Request.Context(), req)
	if err != nil {
		common.SendFaild(c, err)
		return
	}

	c.JSON(http.StatusOK, ins)
}

func (h *Handler) DeleteBlog(c *gin.Context) {
	req := blog.NewDeleteBlogRequest(c.Param("id"))
	ins, err := h.svc.DeleteBlog(c.Request.Context(), req)
	if err != nil {
		common.SendFaild(c, err)
		return
	}

	c.JSON(http.StatusOK, ins)
}

```