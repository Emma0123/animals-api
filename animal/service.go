package animal

type Service interface {
	FindAll() ([]Animals, error)
	FindByID(ID int) (Animals, error)
	FindByName(name string) (Animals, error)
	Create(animalRequest AnimalRequest) (Animals, error)
	Update(ID int, updateAnimalRequest UpdateAnimalRequest) (Animals, error)
	Delete(ID int) (Animals, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Animals, error) {
	animals, err := s.repository.FindAll()
	return animals, err
}

func (s *service) FindByID(ID int) (Animals, error) {
	animal, err := s.repository.FindByID(ID)
	return animal, err
}

func (s *service) FindByName(name string) (Animals, error) {
	animal, err := s.repository.FindByName(name)
	return animal, err
}

func (s *service) Create(animalRequest AnimalRequest) (Animals, error) {

	animal := Animals{
		// Mapping disini
		Name:  animalRequest.Name,
		Class: animalRequest.Class,
		Legs:  animalRequest.Legs,
	}
	newAnimal, err := s.repository.Create(animal)
	return newAnimal, err
}

func (s *service) Update(ID int, updateAnimalRequest UpdateAnimalRequest) (Animals, error) {
	animal, err := s.repository.FindByID(ID)

	animal.Name = updateAnimalRequest.Name
	animal.Class = updateAnimalRequest.Class
	animal.Legs = updateAnimalRequest.Legs

	newAnimal, err := s.repository.Update(animal)
	return newAnimal, err
}

func (s *service) Delete(ID int) (Animals, error) {
	animal, err := s.repository.FindByID(ID)
	newAnimal, err := s.repository.Delete(animal)
	return newAnimal, err
}
