package models

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  uint16 `json:"age"`
}