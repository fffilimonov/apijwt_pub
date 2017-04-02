package models

type Result struct {
	Id   string `json:"id" form:"id"`
	Data string `json:"data" form:"data"`
}
