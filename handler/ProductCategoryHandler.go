package handler

import (
	"3T_Shop/model"
	"3T_Shop/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ProductCategoryHandler struct {
}

var (
	ProductCategoryService service.ProductCategoryService
)

func (ProductCategoryHandler ProductCategoryHandler) Create(c echo.Context) error {
	var ProCategory model.ProductCategory
	if err := c.Bind(&ProCategory); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Tạo loại không thành công",
		})
	}
	if err := ProductCategoryService.Create(&ProCategory); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Tạo loại sản phẩm thành công",
	})
}

func (ProductCategoryHandler ProductCategoryHandler) GetAll(c echo.Context) error {
	ProCategories, err := ProductCategoryService.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "lỗi khi lấy dữ liệu",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Lấy thành công",
		"data":    ProCategories,
	})
}
