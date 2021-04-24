package services

import (
	"gorm-test/errors"
	"gorm-test/model"
	"gorm-test/repository"
)

type PeopleService struct {
	peopleRepository *repository.PeopleRepository
}

func (s *PeopleService) GetPeoples() *[]model.People {
	return s.peopleRepository.GetPeoples()
}

func (s *PeopleService) AddPeople(people *model.People) (*model.People, error) {
	if people.Age < 0 {
		return nil, errors.New("Invalid")
	}
	return s.peopleRepository.AddPeople(people), nil
}

func New(peopleRepository *repository.PeopleRepository) *PeopleService {
	return &PeopleService{
		peopleRepository: peopleRepository,
	}
}
