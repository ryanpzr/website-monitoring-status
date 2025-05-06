package model

type Check struct {
	Id           int    `json:"id"`
	SiteId       int    `json:"site_id"`
	Status       string `json:"status"`
	TimeResponse int64  `json:"time_response"`
	HttpCode     int    `json:"http_code"`
	TimeCreated  string `json:"time_created"`
}
