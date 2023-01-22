package todo

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type ToDoService struct {
	Repo     ITodoRepository
	Validate *validator.Validate
}

type IToDoService interface {
	Create(data ToDoRequest) error
	GetAll() ([]ToDoModel, error)
	GetById(id uint32) (ToDoModel, error)
	Update(id uint32, data ToDoRequest) error
	Delete(id uint32) error
}

func NewToDoSevice(Repo ITodoRepository, validate *validator.Validate) IToDoService {
	return &ToDoService{
		Repo:     Repo,
		Validate: validate,
	}
}

func (t *ToDoService) Create(data ToDoRequest) error {

	err := t.Validate.Struct(data)
	if err != nil {
		return errors.New("data not valid")
	}

	var tempData = ToDoModel{
		Task: data.Task,
	}

	err = t.Repo.Create(tempData)
	if err != nil {
		return errors.New(fmt.Sprintf("Can't create data to db in repository: %s", err.Error()))
	}

	return nil
}

func (t *ToDoService) GetAll() ([]ToDoModel, error) {

	all, err := t.Repo.GetAll()
	if err != nil {
		return nil, err
	}

	return all, nil
}

func (t *ToDoService) GetById(id uint32) (ToDoModel, error) {
	byId, err := t.Repo.GetById(id)
	if err != nil {
		return ToDoModel{}, err
	}

	return byId, nil
}

func (t *ToDoService) Update(id uint32, data ToDoRequest) error {

	err := t.Validate.Struct(data)
	if err != nil {
		return errors.New("data not valid")
	}

	var tempData = ToDoModel{
		Task: data.Task,
	}

	err = t.Repo.Update(id, tempData)
	if err != nil {
		return err
	}

	return nil
}

func (t *ToDoService) Delete(id uint32) error {

	err := t.Repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
