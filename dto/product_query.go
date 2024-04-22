package dto

import "synolux/utils"

type ProductQuery struct {
	utils.Pager
	Title  string `json:"title"`
	Catid  int    `json:"catid"`
	Status int    `json:"status"`
	Limit  int    `json:"limit"`
	Catids []int  `json:"catids"`
}
