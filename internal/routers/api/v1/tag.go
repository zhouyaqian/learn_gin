package v1

import (
	"github.com/gin-gonic/gin"
	"learn_gin/pkg/app"
	"learn_gin/pkg/errcode"
)

type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {

}

// List
// @Summary 获取多个标签
// @Produce json
// @Param name query string false "标签名称" maxlength(100)
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
}

// Create
// @Summary 新增标签
func (t Tag) Create(c *gin.Context) {

}

func (t Tag) Update(c *gin.Context) {

}

func (t Tag) Delete(c *gin.Context) {

}
