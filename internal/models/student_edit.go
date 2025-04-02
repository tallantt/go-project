package models

type StudentEdit struct {
	FullName  string `json:"fullName"`
	Birthdate string `json:"birthdate"`
	Age       int    `json:"age"`
}
