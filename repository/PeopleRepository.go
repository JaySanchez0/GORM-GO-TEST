package repository

import (
	"gorm-test/model"

	"gorm.io/gorm"
)

type PeopleRepository struct {
	db *gorm.DB
}

func (r *PeopleRepository) GetPeoples() *[]model.People {
	var peoples []model.People
	r.db.Preload("Pets").Find(&peoples)
	return &peoples
}

func (r *PeopleRepository) AddPeople(people *model.People) *model.People {
	r.db.Create(&people)
	r.db.Create(&model.Pet{Name: "Dan", Age: 20, PeopleId: int(people.Id)})
	return people
}

func New(db *gorm.DB) *PeopleRepository {
	return &PeopleRepository{
		db: db,
	}
}
