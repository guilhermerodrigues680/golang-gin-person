package person

import "errors"

type Person struct {
	Id   int
	Name string
	Age  int
}

func NewPerson(id int, name string, age int) (*Person, error) {
	if name == "" {
		return &Person{}, errors.New("A pessoa precisa possuir um nome.")
	}

	return &Person{
		Id:   id,
		Name: name,
		Age:  age,
	}, nil
}

func (p *Person) MadeBirthday() int {
	p.Age++
	return p.Age
}
