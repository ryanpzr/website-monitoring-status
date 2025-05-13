package model

type Check struct {
	Id           int    `json:"id"`
	SiteId       int    `json:"site_id" validate:"required"`
	Status       string `json:"status" validate:"required,alpha"`
	TimeResponse int64  `json:"time_response" validate:"required"`
	HttpCode     int    `json:"http_code" validate:"required,numeric"`
	TimeCreated  string `json:"time_created"`
}
