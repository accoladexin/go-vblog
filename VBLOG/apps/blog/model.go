package blog

import (
	"encoding/json"
	"github.com/accoladexin/vblog/common"
	"time"
)

// 本文件主要用于声明接口的数据结构
func NewBlogSet() *BlogSet {
	return &BlogSet{
		Items: []*Blog{},
	}
}

type BlogSet struct {
	// 总共多个篇文章
	Total int64   `json:"total"`
	Items []*Blog `json:"items"`
}

// Stringer
func (s *BlogSet) String() string {
	dj, _ := json.MarshalIndent(s, "", "  ")
	return string(dj)
}

// 通过一个构建函数来构建Blog
// 为什么要使用构建函数？为什么不直接使用struct
// 使用构造好按时保证兼容性: 把需要初始化的 在这个函数进行初始化,
// 有需要默认执行，补充默认值
func NewBlog(req *CreateBlogRequest) *Blog {
	//fmt.Println(NewMeta())
	return &Blog{
		Meta:              NewMeta(),
		CreateBlogRequest: req,
	}
}

// 数据
// 这个对象如何保存数据库里面 ?
// 通过匿名嵌套来组合产生新的结构体
type Blog struct {
	*Meta
	*CreateBlogRequest
}

// Stringer
func (s *Blog) String() string {
	dj, _ := json.MarshalIndent(s, "", "  ")
	return string(dj)
}

//	type Tabler interface {
//		TableName() string
//	}
//

// 定义gorm 存入数据时表的名称
func (i *Blog) TableName() string {
	return "blogs"
}

func NewMeta() *Meta {
	return &Meta{
		CreatedAt: time.Now().Unix(),
		//UpdatedAt: 12341234,
	}
}

// 文章的元数据:
// + 文章的Id
// + 创建时间
// + 修改时间
// + 发布时间
type Meta struct {
	Id int `json:"id"`
	// 直接用时间戳, 我选择用时间戳
	// 我们是做后端, 一般数据库的时间对象是又时区约束
	// 后端直接存储时间戳, 当需要展示的时候，由前端(Web,APP,...)负责带上用户的当前时区做展示
	//默认情况下，如果模型中有 created_at updated_at deleted_at 字段，并且类型为 time.Time 或者 int64, GORM 默认会自动维护这些时间戳。
	CreatedAt   int64 `json:"created_at"`
	UpdatedAt   int64 `json:"updated_at"`
	PublishedAt int64 `json:"pulished_at"`
}

func NewCreateBlogRequest() *CreateBlogRequest {
	return &CreateBlogRequest{
		Tags:   map[string]string{},
		Status: STATUS_DRAFT,
	}
}

// 用户传入的数据
// + 标题
// + 作者
// + 内容(Markdown)
// + 标签(map)
// GORM Object ---> Table Row
// Object struct Tag:  gorm:"column:title"
// 你不定义tag,默认使用你json的tag
// Insert （title) VALUE (?)
type CreateBlogRequest struct {
	//  validate:"required"` 官方的校验工具
	// CreateAt --> createAt
	// CreateAt --> create_at(通常我们json的tag也是采用蛇形缩写)
	// json, gorm, text
	// 文章标题
	Title   string `json:"title" gorm:"column:title" validate:"required"`
	Author  string `json:"author" gorm:"column:author" validate:"required"`
	Content string `json:"content" validate:"required"`
	// map[string]string orm是不知道如何入库的
	// "serializer:json"` 直接存成json数据到数据库里面
	Tags map[string]string `json:"tags" gorm:"serializer:json"`
	// 文章是由状态
	Status STATUS `json:"status"`

	// 新加了一个字段
	// Extra map[string]string
}

// 检查用户提交的参数是否合法
// 使用这个来做校验: https://github.com/go-playground/validator
func (req *CreateBlogRequest) Validate() error {
	// 调用的全局校验器，详见https://github.com/go-playground/validator
	return common.Validate(req)
}
