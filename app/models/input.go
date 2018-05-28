package models

type Input struct {
	Filename string `json:"filename,omitempty" form:"filename"`
	Body     string `json:"html,omitempty" form:"input"`
}
