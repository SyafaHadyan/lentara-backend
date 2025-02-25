package dto

type RequestCreateProduct struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Stock       int32  `json:"stock"`
	PhotoUrl    string `json:"photo_url"`
}

type ResponseCreateProduct struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Stock       int32  `json:"stock"`
	PhotoUrl    string `json:"photo_url"`
}

