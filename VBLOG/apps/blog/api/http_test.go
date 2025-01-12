package api_test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

type CreateBlogRequest1 struct {
	Title   string `json:"title" gorm:"column:title" validate:"required"`
	Author  string `json:"author" gorm:"column:author" validate:"required"`
	Content string `json:"content" validate:"required"`

	Tags map[string]string `json:"tags" gorm:"serializer:json"`
}

func dealfunc(c *gin.Context) {
	ins := &CreateBlogRequest1{}
	c.BindJSON(ins)
	fmt.Println(ins)
}
func TestHTTP1(t *testing.T) {
	r := gin.Default()
	r.RouterGroup.GET("/api", dealfunc)
	r.Run(":11220")
}
