package usecase

import (
	"github.com/koizumi7010/go-todo-api/domain/models"
	"github.com/koizumi7010/go-todo-api/domain/repository"
)

type ToDo interface {
	Create(task string) error
	Delete(id int) error
	Update(id int, task string, status models.TaskStatus) error
	Get(id int) (*models.ToDo, error)
	GetAll() ([]*models.ToDo, error)
}

type todo struct {
	repository repository.ToDo
}

func NewTodo(r repository.ToDo) ToDo {
	return &todo{r}
}

func (t *todo) Create(task string) error {
	td := models.NewTodo(task)
	if err := t.repository.Create(td); err != nil {
		return err
	}
	return nil
}

func (t *todo) Delete(id int) error {
	if err := t.repository.Delete(id); err != nil {
		return err
	}
	return nil
}

func (t *todo) Update(id int, task string, status models.TaskStatus) error {
	td := models.NewUpdateTodo(id, task, status)
	if err := t.repository.Update(td); err != nil {
		return err
	}
	return nil
}

func (t *todo) Get(id int) (*models.ToDo, error) {
	td, err := t.repository.Get(id)
	if err != nil {
		return nil, err
	}
	return td, nil
}

func (t *todo) GetAll() ([]*models.ToDo, error) {
	tds, err := t.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return tds, nil
}
