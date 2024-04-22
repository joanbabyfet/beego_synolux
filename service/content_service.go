package service

import (
	"errors"
	"synolux/models"

	"github.com/beego/beego/logs"
)

type ContentService struct {
}

// 根据编码获取详情
func (s *ContentService) GetByCode(code string) (*models.Content, error) {
	entity := new(models.Content)
	info, err := entity.GetByCode(code)
	if err != nil {
		logs.Error("内容不存在")
		return nil, errors.New("内容不存在")
	}
	return info, nil
}
