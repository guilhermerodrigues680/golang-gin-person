package person

type Repository interface {
	Create(name string, age int) (*Person, error)
	ReadAll() ([]*Person, error)
	Read(id int) (*Person, error)
	Update(id int, name string, age int) (*Person, error)
}

type personService struct {
	r Repository
}

func NewPersonService(r Repository) *personService {
	return &personService{r: r}
}

func (s *personService) Create(name string, age int) (*Person, error) {
	return s.r.Create(name, age)
}

func (s *personService) ReadAll() ([]*Person, error) {
	return s.r.ReadAll()
}

func (s *personService) Read(id int) (*Person, error) {
	return s.r.Read(id)
}

func (s *personService) MakeBirthday(id int) (*Person, error) {
	p, err := s.r.Read(id)
	if err != nil {
		return nil, err
	}

	p.MadeBirthday()
	return s.r.Update(p.Id, p.Name, p.Age)
}
