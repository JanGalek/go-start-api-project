package entities

import "jan-galek/api/src/types"

type Article struct {
	types.Abstract
	Title string `json:"title"`
	Link  string `json:"link"`
}
