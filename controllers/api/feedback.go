// 父控制器
package controllers

import (
	"synolux/models"
	"synolux/service"

	"github.com/beego/beego/v2/core/validation"
)

type FeedbackController struct {
	BaseController
}

// 保存
func (c *FeedbackController) Save() {
	name := c.GetString("name")
	mobile := c.GetString("mobile")
	email := c.GetString("email")
	content := c.GetString("content")

	//参数验证
	entity := models.Feedback{
		Name:    name,
		Mobile:  mobile,
		Email:   email,
		Content: content,
	}
	valid := validation.Validation{}
	valid.Required(entity.Name, "name")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.ErrorJson(-1, err.Key+err.Error(), nil)
		}
	}

	service_feedback := new(service.FeedbackService)
	stat, err := service_feedback.Save(entity)
	if stat < 0 {
		c.ErrorJson(stat, err.Error(), nil)
	}
	c.SuccessJson("success", nil)
}
