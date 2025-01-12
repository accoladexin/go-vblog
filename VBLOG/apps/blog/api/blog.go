package api

import (
	"fmt"
	"github.com/accoladexin/vblog/apps/blog"
	"github.com/accoladexin/vblog/common"
	"github.com/accoladexin/vblog/common/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 使用web框架: Gin
func (h *Handler) CreateBlog(c *gin.Context) {
	// 从web框架中 获取用户的请求,请求参数放到Body
	// 原始的处理方式 读取了数据就没了

	//all, err2 := io.ReadAll(c.Request.Body)
	//if err2 != nil {
	//	panic(err2)
	//}
	//
	//fmt.Println("原始获取数据method: ", all)

	// gin封装后处理逻辑
	fmt.Println("gin封装后的处理逻辑: ============================================")
	//logger.L().Info().Msg("gin收到数据")
	//fmt.Println("===== Request Info =====")
	//fmt.Println("Method:", c.Request.Method)
	//fmt.Println("URL:", c.Request.URL)
	//fmt.Println("Header:", c.Request.Header)
	//fmt.Println("Body:", c.Request.Body) // 注意这里是 io.ReadCloser 类型，需要读取才能看到内容
	//body, err := io.ReadAll(c.Request.Body)
	//if err != nil {
	//	fmt.Println("Error read body", err)
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//fmt.Println("Request Body:", string(body)) // 读取并打印内容

	ins := blog.NewCreateBlogRequest()
	//c.ShouldBindBodyWithJSON(ins)
	c.BindJSON(ins)
	fmt.Println("=============>", ins)
	logger.L().Info().Interface("request", ins)
	createBlog, err2 := h.svc.CreateBlog(c.Request.Context(), ins)
	if err2 != nil {
		common.SendFaild(c, err2)
		return
	}
	createBlog.String()

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

	// 具体的Serice处理逻辑
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
