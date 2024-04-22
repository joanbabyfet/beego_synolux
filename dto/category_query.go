package dto

import "synolux/utils"

type CategoryQuery struct {
	utils.Pager
	Type   int `json:"type"`
	Pid    int `json:"pid"`
	Status int `json:"status"`
}
