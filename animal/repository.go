package animal

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Animals, error)
	FindByID(ID int) (Animals, error)
	FindByName(name string) (Animals, error)
	Create(animal Animals) (Animals, error)
	Update(animal Animals) (Animals, error)
	Delete(animal Animals) (Animals, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Animals, error) {
	var animals []Animals

	err := r.db.Find(&animals).Error

	return animals, err

}

func (r *repository) FindByID(ID int) (Animals, error) {
	var animals Animals

	err := r.db.Find(&animals, ID).Error

	return animals, err

}

func (r *repository) FindByName(name string) (Animals, error) {
	var animals Animals
	err := r.db.Find(&animals, "name = ?", name).Error
	return animals, err
}

func (r *repository) Create(animal Animals) (Animals, error) {

	err := r.db.Create(&animal).Error

	return animal, err

}

func (r *repository) Update(animal Animals) (Animals, error) {
	err := r.db.Save(&animal).Error

	return animal, err
}

func (r *repository) Delete(animal Animals) (Animals, error) {
	err := r.db.Delete(&animal).Error

	return animal, err
}
