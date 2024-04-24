// 父控制器
package admin

import (
	dto "synolux/dto"
	"synolux/models"
	"synolux/service"

	"github.com/beego/beego/v2/core/validation"
)

type ArticleController struct {
	AdminBaseController
}

// 获取列表
func (c *ArticleController) Index() {
	catid, _ := c.GetInt("catid")
	page, _ := c.GetInt("page")
	page_size, _ := c.GetInt("page_size")
	if page < 1 {
		page = 1
	}
	if page_size < 1 {
		page_size = 10
	}

	//获取文章列表
	query := dto.ArticleQuery{}
	query.Catid = catid
	query.Page = page
	query.PageSize = page_size
	query.Status = 1
	service_article := new(service.ArticleService)
	list, total := service_article.PageList(query)

	//组装数据
	resp := make(map[string]interface{}) //创建1个空集合
	var next_page int                    //是否有下一页
	if page*page_size >= int(total) {
		next_page = 0
	} else {
		next_page = 1
	}
	resp["next_page"] = next_page
	resp["total"] = total
	resp["list"] = list
	c.SuccessJson("success", resp)
}

// 获取详情
func (c *ArticleController) Detail() {
	id, _ := c.GetInt("id")

	//参数验证
	entity := models.Article{Id: id}
	valid := validation.Validation{}
	valid.Required(entity.Id, "id")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.ErrorJson(-1, err.Key+err.Error(), nil)
		}
	}

	service_article := new(service.ArticleService)
	info, err := service_article.GetById(id)
	if err != nil {
		c.ErrorJson(-2, err.Error(), nil)
	}
	c.SuccessJson("success", info)
}

// 保存
func (c *ArticleController) Save() {
	id, _ := c.GetInt("id")
	catid, _ := c.GetInt("catid")
	title := c.GetString("title")
	info := c.GetString("info")
	content := c.GetString("content")
	author := c.GetString("author")
	status, _ := c.GetInt8("status")

	//参数验证
	entity := models.Article{
		Id:      id,
		Catid:   catid,
		Title:   title,
		Info:    info,
		Content: content,
		Author:  author,
		Status:  status,
	}
	valid := validation.Validation{}
	if entity.Id > 0 {
		valid.Required(entity.Id, "id")
	}
	valid.Required(entity.Title, "title")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.ErrorJson(-1, err.Key+err.Error(), nil)
		}
	}

	service_article := new(service.ArticleService)
	stat, err := service_article.Save(entity)
	if stat < 0 {
		c.ErrorJson(stat, err.Error(), nil)
	}
	c.SuccessJson("success", nil)
}

// 删除
func (c *ArticleController) Delete() {
	id, _ := c.GetInt("id")

	//参数验证
	entity := models.Article{Id: id}
	valid := validation.Validation{}
	valid.Required(entity.Id, "id")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.ErrorJson(-1, err.Key+err.Error(), nil)
		}
	}

	service_article := new(service.ArticleService)
	stat, err := service_article.DeleteById(id)
	if stat < 0 {
		c.ErrorJson(stat, err.Error(), nil)
	}
	c.SuccessJson("success", nil)
}

// 启用
func (c *ArticleController) Enable() {
	id, _ := c.GetInt("id")

	//参数验证
	entity := models.Article{Id: id}
	valid := validation.Validation{}
	valid.Required(entity.Id, "id")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.ErrorJson(-1, err.Key+err.Error(), nil)
		}
	}

	service_article := new(service.ArticleService)
	stat, err := service_article.EnableById(id)
	if stat < 0 {
		c.ErrorJson(stat, err.Error(), nil)
	}
	c.SuccessJson("success", nil)
}

// 禁用
func (c *ArticleController) Disable() {
	id, _ := c.GetInt("id")

	//参数验证
	entity := models.Article{Id: id}
	valid := validation.Validation{}
	valid.Required(entity.Id, "id")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.ErrorJson(-1, err.Key+err.Error(), nil)
		}
	}

	service_article := new(service.ArticleService)
	stat, err := service_article.DisableById(id)
	if stat < 0 {
		c.ErrorJson(stat, err.Error(), nil)
	}
	c.SuccessJson("success", nil)
}
