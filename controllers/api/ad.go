// 父控制器
package controllers

import (
	dto "synolux/dto"
	"synolux/service"
)

type AdController struct {
	BaseController
}

// 获取列表
func (c *AdController) Index() {
	limit, _ := c.GetInt("limit")

	//获取文章列表
	query := dto.AdQuery{}
	query.Limit = limit
	query.Status = 1
	service_Ad := new(service.AdService)
	list := service_Ad.All(query)

	//组装数据
	resp := make(map[string]interface{}) //创建1个空集合
	resp["list"] = list
	c.SuccessJson("success", resp)
}
