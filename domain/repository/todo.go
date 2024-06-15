package repository

import "github.com/koizumi7010/go-todo-api/domain/models"

type ToDo interface {
	Create(t *models.ToDo) error
	Delete(id int) error
	Update(t *models.ToDo) error
	Get(id int) (*models.ToDo, error)
	GetAll() ([]*models.ToDo, error)
}
