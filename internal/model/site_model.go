package model

type Site struct {
	Name string `json:"name"`
	Url  string `json:"url"`
	Freq int    `json:"freq"`
}
