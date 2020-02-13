package database

import (
	"github.com/Muha113/6-th-term-labs/itirod/lab0/pkg/model"
)

type Database struct {
}

func configureDatabase() {
}

func (d *Database) CreateUser(model.User) error {
	return nil
}

func (d *Database) CreateDialogue(model.Dialogue) error {
	return nil
}

func (d *Database) CreateGroup(model.Group) error {
	return nil
}

func (d *Database) CreateMessage(model.Message) error {
	return nil
}

func (d *Database) GetAllUsers() ([]model.User, error) {
	return []model.User{}, nil
}

func (d *Database) GetUserByID(id uint) (model.User, error) {
	return model.User{}, nil
}

func (d *Database) GetUserByName(name string) (model.User, error) {
	return model.User{}, nil
}
