package model

type People struct {
	Id   uint  `gorm:"primaryKey" json:"id"`
	Age  int   `json:"age"`
	Pets []Pet `gorm:"foreignKey:PeopleId" json:"pets"`
}
