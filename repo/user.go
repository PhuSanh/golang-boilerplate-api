package repo

import (
	"golang-boilerplate/database"
	"golang-boilerplate/model"
)

type UserRepo interface {
	GetListUser() ([]model.User, error)
}

func GetListUsers() (listUser []model.User, err error) {
	err = database.MysqlConn.Find(&listUser).Error
	return
}