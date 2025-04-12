package models

type CarEdit struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
	Year  int    `json:"year"`
	Color string `json:"color"`
}
