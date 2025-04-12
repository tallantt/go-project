package models

type Car struct {
	Id    int    `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Year  int    `json:"year"`
	Color string `json:"color"`
}
