package model

type Site struct {
	Name string `json:"name" validate:"required,min=4,max=100"`
	Url  string `json:"url" validate:"required,startswith=http"`
	Freq int    `json:"freq" validate:"required,numeric"`
	Id   int    `json:"id"`
}
