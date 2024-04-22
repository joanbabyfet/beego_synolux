package dto

import "synolux/utils"

type RoleQuery struct {
	utils.Pager
	Name string `json:"name"`
}
