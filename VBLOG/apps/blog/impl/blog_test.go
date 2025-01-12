package impl_test

import (
	"context"
	"fmt"
	"github.com/accoladexin/vblog/apps/blog"
	"github.com/accoladexin/vblog/conf"
	"os"
	"testing"
)

var (
	ctx = context.Background()
)

func TestBlog(t *testing.T) {
	t.Log(os.Getenv("MYSQL_HOST")) // 获取环境变量
	t.Log(os.Getenv("goroot"))     // D:\development_tool\goland1.23
	if os.Getenv("MYSQL_HOST") == "" {
		t.Error("MYSQL_HOST is empty")
	} else {
		fmt.Println("hav sql")
	}
	fmt.Println("--------------")
}

//func TestImpl_CreateBlog(t *testing.T) {
//	in := &blog.CreateBlogRequest{}
//	fmt.Println(in)
//}

func TestImpl_CreateBlog(t *testing.T) {
	// 创建一个对象
	in := &blog.CreateBlogRequest{} //type CreateBlogRequest struct
	in.Title = "test111"
	in.Author = "test"
	in.Content = "test"
	in.Tags = map[string]string{
		"tag1": "tag1",
		"tag2": "tag2",
	}

	// 调用方法,contoller在另外一个测试单元中
	// // sql语句：INSERT INTO `blogs`
	//	//(`created_at`,`updated_at`,`published_at`,`title`,`author`,`content`,`tags`,`status`)
	//	//VALUES (1736591397,1736591397,0,'test','test','test','{"tag1":"tag1","tag2":"tag2"}',0)
	createBlog, err := controller.CreateBlog(ctx, in)
	if err != nil {
		fmt.Println("=================")
		t.Error(err)
	}
	t.Log(createBlog)

}
func TestGorm(t *testing.T) {
	// 创建一个对象
	in := &blog.Blog{}
	conf.LoadConfigFromEnv()
	orm := conf.C().MySQL.ORM().Debug()
	err := orm.WithContext(context.Background()).Save(in).Error
	if err != nil {
		t.Error(err)
	}
	t.Log(in)

}

func Test_error(t *testing.T) {
	blog.NewCreateBlogRequest()

}
