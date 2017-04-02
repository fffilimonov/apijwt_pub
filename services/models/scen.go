package models

type Scen struct {
	Name string `json:"name" form:"name"`
	Text string `json:"text" form:"text"`
}

type Dashboard struct {
	Scens   []Scen `json:"scens" from:"scens"`
	Seconds string `json:"seconds" form:"seconds"`
}
