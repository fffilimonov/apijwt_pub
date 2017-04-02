package models

type Run struct {
	Scen string `json:"scen" form:"scen"`
	Browser string `json:"browser" form:"browser"`
}
