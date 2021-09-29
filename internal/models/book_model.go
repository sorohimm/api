package models

type Book struct {
	Uuid         string `json:"uuid"`
	Name         string `json:"name"         validate:"required"`
	Year         string `json:"year"         validate:"required"`
	Author       string `json:"author"       validate:"required"`
	Category     string `json:"category"     validate:"required"`
	Price        string `json:"price"        validate:"required"`
	Descriptions string `json:"descriptions" validate:"required"`
}
