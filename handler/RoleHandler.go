package handler

import (
	"3T_Shop/model"
	"3T_Shop/service"
	"3T_Shop/utils"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type RoleHandler struct {
}

var (
	RoleService service.RoleService
)

func (RoleHandler RoleHandler) Create(c echo.Context) error {
	var role model.Role
	err := c.Bind(&role)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.CreateResponse(http.StatusBadRequest, errors.New("không thể bind data").Error(), utils.EmptyObject{}))
	}
	if err := RoleService.Create(&role); err != nil {
		return c.JSON(http.StatusBadRequest, utils.CreateResponse(http.StatusBadRequest, err.Error(), utils.EmptyObject{}))
	}
	return c.JSON(http.StatusOK, utils.CreateResponse(http.StatusOK, "tạo đối tượng thành công", utils.EmptyObject{}))
}

func (RoleHandler RoleHandler) GetAll(c echo.Context) error {
	roles, err := RoleService.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.CreateResponse(http.StatusBadRequest, err.Error(), utils.EmptyObject{}))
	}
	return c.JSON(http.StatusOK, utils.CreateResponse(http.StatusOK, "lấy dữ liêu thành công", roles))
}
