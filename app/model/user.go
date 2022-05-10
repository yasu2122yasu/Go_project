package model

type User struct {
	Id   string `json:"id"	db:"id"`
	Name string `json:"name" db:"name"`
	Age  uint32 `json:"age"  db:"age"`
}
