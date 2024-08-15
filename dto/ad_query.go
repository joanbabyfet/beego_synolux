package dto

import "synolux/utils"

type AdQuery struct {
	utils.Pager
	Catid  int `json:"catid"`
	Status int `json:"status"`
	Limit  int `json:"limit"`
}
