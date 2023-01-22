package todo

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type ITodoRepository interface {
	Create(data ToDoModel) error
	GetAll() ([]ToDoModel, error)
	GetById(id uint32) (ToDoModel, error)
	Update(id uint32, uData ToDoModel) error
	Delete(id uint32) error
}

func NewToDoRepository(db *gorm.DB) ITodoRepository {
	return &ToDoRepository{db: db}
}

func (t *ToDoRepository) Create(data ToDoModel) error {
	result := t.db.Create(&data)

	if result.Error != nil {
		return errors.New(fmt.Sprintf("can't create data: %s", result.Error.Error()))
	}

	return nil
}

func (t *ToDoRepository) GetAll() ([]ToDoModel, error) {
	var data []ToDoModel
	result := t.db.Find(&data)

	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("can't get data: %s", result.Error.Error()))
	}

	return data, nil
}

func (t *ToDoRepository) GetById(id uint32) (ToDoModel, error) {
	var data = ToDoModel{
		Id: id,
	}

	err := t.db.First(&data).Error
	if err != nil {
		return ToDoModel{}, errors.New("Data not found")
	}

	return data, nil
}

func (t *ToDoRepository) Update(id uint32, uData ToDoModel) error {
	var data = ToDoModel{
		Id: id,
	}

	err := t.db.First(&data).Error
	if err != nil {
		return errors.New(fmt.Sprintf("Data with id %d not found: %s", id, err.Error()))
	}

	data.Task = uData.Task

	err = t.db.Save(&data).Error
	if err != nil {
		return errors.New(fmt.Sprintf("Can't update data %d: %s", id, err.Error()))
	}

	return nil
}

func (t *ToDoRepository) Delete(id uint32) error {
	err := t.db.Delete(&ToDoModel{}, id).Error
	if err != nil {
		return errors.New(fmt.Sprintf("Can't delete data %d: %s", id, err.Error()))
	}

	return nil
}

type ToDoRepository struct {
	db *gorm.DB
}
