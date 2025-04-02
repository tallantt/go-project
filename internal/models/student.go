package models

type Student struct {
	Id        int    `json:"id"`
	FullName  string `json:"fullName"`
	Birthdate string `json:"birthdate"`
	Age       int    `json:"age"`
}
