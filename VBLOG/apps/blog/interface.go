package blog

import "context"

const (
	// App指的一个业务单元
	// ioc.GetController(AppName) 获取到 这个业务模块的具体实现
	// 简单的一中抽象, (常量/变量) 都是一种引用, 不用硬编码: blogs, 出行了100次
	AppName = "blogs"
)

// 博客管理业务接口(CRUD)
// Blog 是CreateBlog这个接口的参数，是用户传递的数据
// 定义接口的时候, 你站在顶层来进行设计, 站在使用者的角度, 代码调用方
type Service interface {
	// 查询文章列表
	QueryBlog(context.Context, *QueryBlogRequest) (*BlogSet, error)
	// 查询单个文章
	DescribeBlog(context.Context, *DescribeBlogRequest) (*Blog, error)
	// 接口一定要保证很强一个兼容性
	CreateBlog(context.Context, *CreateBlogRequest) (*Blog, error)
	// 更新文章
	UpdateBlog(context.Context, *UpdateBlogRequest) (*Blog, error)
	// 删除文章, 返回删除的对象, 用前端提升, 用于对象最终
	DeleteBlog(context.Context, *DeleteBlogRequest) (*Blog, error)
}

func NewQueryBlogRequest() *QueryBlogRequest {
	return &QueryBlogRequest{
		PageSize:   20,
		PageNumber: 1,
	}
}

// 1. 列表查询请求的参数
// 服务端分页, 默认执: 1页 20个
// 关键字查询, 模糊搜索
// 条件过滤, 比如过滤作者是谁的文章
type QueryBlogRequest struct {
	// 一页多少个
	PageSize int
	// 当前是那页
	PageNumber int
	// 模糊搜索, 搜索文章内容
	Keywords string
	// 条件过滤
	Author string
}

func (r *QueryBlogRequest) Offset() int {
	return 1
}

func NewDescribeBlogRequest(id string) *DescribeBlogRequest {
	return &DescribeBlogRequest{
		Id: id,
	}
}

type DescribeBlogRequest struct {
	Id string
}

func NewUpdateBlogRequest(id string) *UpdateBlogRequest {
	return &UpdateBlogRequest{
		DescribeBlogRequest: NewDescribeBlogRequest(id),
		CreateBlogRequest:   NewCreateBlogRequest(),
	}
}

type UpdateBlogRequest struct {
	*DescribeBlogRequest
	*CreateBlogRequest
}

func NewDeleteBlogRequest(id string) *DeleteBlogRequest {
	return &DeleteBlogRequest{
		DescribeBlogRequest: NewDescribeBlogRequest(id),
	}
}

type DeleteBlogRequest struct {
	*DescribeBlogRequest
}
