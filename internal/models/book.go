package models

type Book struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Year         string `json:"year"`
	Author       string `json:"author"`
	Category     string `json:"category"`
	Price        string `json:"price"`
	Descriptions string `json:"descriptions"`
}

type BookResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Year         string `json:"year"`
	Author       string `json:"author"`
	Category     string `json:"category"`
	Price        string `json:"price"`
	Descriptions string `json:"descriptions"`
}