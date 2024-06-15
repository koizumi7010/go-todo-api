package mysql

import (
	"github.com/koizumi7010/go-todo-api/domain/models"
	"github.com/koizumi7010/go-todo-api/domain/repository"
	"gorm.io/gorm"
)

type ToDo struct {
	db *gorm.DB
}

func NewTodo(db *gorm.DB) repository.ToDo {
	return &ToDo{
		db: db,
	}
}

func (td *ToDo) Create(t *models.ToDo) error {
	if err := td.db.Create(t).Error; err != nil {
		return err
	}
	return nil
}

func (td *ToDo) Delete(id int) error {
	if err := td.db.Where("id = ?", id).Delete(&models.ToDo{}).Error; err != nil {
		return err
	}
	return nil
}

func (td *ToDo) Update(t *models.ToDo) error {
	if err := td.db.Save(t).Error; err != nil {
		return err
	}
	return nil
}

func (td *ToDo) Get(id int) (*models.ToDo, error) {
	var t *models.ToDo
	if err := td.db.Where("id = ?", id).Take(&t).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return t, nil
}

func (td *ToDo) GetAll() ([]*models.ToDo, error) {
	var ts []*models.ToDo
	if err := td.db.Find(&ts).Error; err != nil {
		return nil, err
	}
	return ts, nil
}
