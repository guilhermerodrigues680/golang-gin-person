package person

import (
	"fmt"
)

type personRepository struct {
	persons []*Person
	serial  int
}

func NewPersonRepository() *personRepository {
	return &personRepository{
		persons: make([]*Person, 0),
		serial:  1,
	}
}

func (r *personRepository) Create(name string, age int) (*Person, error) {
	p, err := NewPerson(r.serial, name, age)
	r.serial++
	if err != nil {
		return nil, err
	}

	r.persons = append(r.persons, p)
	return p, err
}

func (r *personRepository) ReadAll() ([]*Person, error) {
	return r.persons, nil
}

func (r *personRepository) Read(id int) (*Person, error) {
	for _, v := range r.persons {
		if v.Id == id {
			return v, nil
		}
	}

	return nil, fmt.Errorf("Não há uma pessoa associada ao ID %d", id)
}

func (r *personRepository) Update(id int, name string, age int) (*Person, error) {
	for _, v := range r.persons {
		if v.Id == id {
			v.Name = name
			v.Age = age
			return v, nil
		}
	}

	return nil, fmt.Errorf("Não há uma pessoa associada ao ID %d", id)
}
