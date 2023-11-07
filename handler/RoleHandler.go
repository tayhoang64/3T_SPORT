package handler

import (
	"3T_Shop/model"
	"3T_Shop/service"
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
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "đối trượng không đúng định dạng",
		})
	}
	if err := RoleService.Create(&role); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "tạo đối tươg thành công",
	})
}

func (RoleHandler RoleHandler) GetAll(c echo.Context) error {
	roles, err := RoleService.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "lỗi khi lấy roles",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Lấy thành công",
		"data":    roles,
	})
}
