package service

import (
	"errors"
	"synolux/models"
	"synolux/utils"

	"github.com/beego/beego/v2/core/logs"
)

type FeedbackService struct {
}

// 保存
func (s *FeedbackService) Save(entity models.Feedback) (int, error) {
	stat := 1

	if entity.Id > 0 {

	} else {
		entity.CreateUser = "1"                  //添加人
		entity.CreateTime = utils.GetTimestamp() //添加时间
		_, err := entity.Add()
		if err != nil {
			logs.Error("反馈添加失败")
			return -2, errors.New("反馈添加失败")
		}
	}
	return stat, nil
}
