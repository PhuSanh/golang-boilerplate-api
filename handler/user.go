package handler

import (
	"github.com/labstack/echo/v4"
	"golang-boilerplate/common"
	"golang-boilerplate/repo"
)

type UserHandler struct {}

func (u UserHandler) GetList(c echo.Context) (err error) {
	users, err := repo.GetListUsers()
	if err != nil {
		return RespondToClient(c, common.ERROR_GET_ROW_FROM_DB, common.MSG_GET_ROW_FROM_DB, nil)
	}
	return RespondToClient(c, common.ERROR_NO_ERORR, common.MSG_SUCEESS, users)
}