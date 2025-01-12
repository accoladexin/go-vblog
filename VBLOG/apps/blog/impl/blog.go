package impl

import (
	"context"
	"fmt"
	"github.com/accoladexin/vblog/apps/blog"
	"github.com/accoladexin/vblog/common/logger"
	"github.com/accoladexin/vblog/conf"
)

// ctx: Context上下文, 一个接口或者函数, 他的参数保护2类数据:
// 1. 用户传递过来的数据: in *blog.Blog
// 2. 还有1类数据 不是通过用户传递的, 比如Trace, RequestId
func (i *impl) CreateBlog(ctx context.Context, in *blog.CreateBlogRequest) (*blog.Blog, error) {
	// 校验数据

	logger.L().Debug().
		Str("Method", "CreateBlog").
		Str("Add", conf.C().Http.Address()).
		Interface("request", in).
		Msg(fmt.Sprintf("CreateBlog called, Context: %+v, Request:%+v", ctx, in))
	if err := in.Validate(); err != nil {
		return nil, err
	}

	// 构造插入数据，同时更新时间
	ins := blog.NewBlog(in)

	// 尽量不要用
	//i.db.Save(ins)

	// 链式调用，
	// sql：INSERT INTO `blogs` (`created_at`,`updated_at`,`published_at`,`title`,`author`,`content`,`tags`,`status`)
	//VALUES (1736591397,1736591397,0,'test','test','test','{"tag1":"tag1","tag2":"tag2"}',0)
	//
	//fmt.Println(ins) // updated_at时间没有，
	err := i.db.WithContext(ctx).Save(ins).Error
	// fmt.Println(ins)
	if err != nil {
		return nil, err
	}
	return ins, nil
}

// 查询文章列表
func (i *impl) QueryBlog(ctx context.Context, in *blog.QueryBlogRequest) (
	*blog.BlogSet, error) {
	// 提前构造一个BlogSet对象出来
	set := blog.NewBlogSet()

	//  需要构造SQL查询
	query := i.db.WithContext(ctx).Model(&blog.Blog{})

	// 模糊查询
	if in.Keywords != "" {
		// %xxxx% ---> *xxxx*
		// go项目教学  --> 项目教学     (%项目教学%)
		query = query.Where("content LIKE ?", "%"+in.Keywords+"%")
	}
	if in.Author != "" {
		query = query.Where("author = ?", in.Author)
	}

	// 分页是必传参数? offset 忽略前面多少个，获取多少个数据
	// 数据是一个[]*Blog
	err := query.
		// COUNT(*) WHERE xxx OFFSET LIMIT
		Count(&set.Total).
		Offset(int(in.Offset())).
		Limit(in.PageSize).
		// SELECT * FORM ....
		// SELECT * FROM `blogs` LIMIT 20
		// SELECT * FROM `blogs` WHERE content LIKE '%项目教学%' LIMIT 20
		// MySQL的教程对SQL基本使用有所了解
		Find(&set.Items).
		Error
	if err != nil {
		return nil, err
	}

	return set, nil
}

// 查询单个文章
// id = "xxx" 根据文章id的一个条件过滤查询
func (i *impl) DescribeBlog(ctx context.Context, in *blog.DescribeBlogRequest) (
	*blog.Blog, error) {
	//  需要构造SQL查询
	query := i.db.WithContext(ctx).Model(&blog.Blog{})

	ins := blog.NewBlog(blog.NewCreateBlogRequest())
	err := query.
		Where("id = ?", in.Id).
		First(ins).
		Error
	if err != nil {
		return nil, err
	}

	return ins, nil
}

// 更新文章
func (i *impl) UpdateBlog(ctx context.Context, in *blog.UpdateBlogRequest) (
	*blog.Blog, error) {

	return nil, nil
}

// 删除文章, 返回删除的对象, 用前端提升, 用于对象最终
func (i *impl) DeleteBlog(ctx context.Context, in *blog.DeleteBlogRequest) (
	*blog.Blog, error) {

	return nil, nil
}
