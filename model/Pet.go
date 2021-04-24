package model

type Pet struct {
	Id       uint   `gorm:"primaryKey"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Species  string `json:"species"`
	PeopleId int    `json:"-"`
}
