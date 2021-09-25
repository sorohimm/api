package models

type Book struct {
	Uuid         string `json:"name"         validate:"required"`
	Year         string `json:"year"         validate:"required"`
	Author       string `json:"author"       validate:"required"`
	Category     string `json:"category"     validate:"required"`
	Price        string `json:"price"        validate:"required"`
	Descriptions string `json:"descriptions" validate:"required"`
}

type BookResponse struct {
	Uuid         string `json:"uuid"`
	Name         string `json:"name"`
	Year         string `json:"year"`
	Author       string `json:"author"`
	Category     string `json:"category"`
	Price        string `json:"price"`
	Descriptions string `json:"descriptions"`
}
