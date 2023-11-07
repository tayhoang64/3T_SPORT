package handler

import (
	"3T_Shop/model"
	"3T_Shop/service"
	"3T_Shop/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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

func (ProductCategoryHandler ProductCategoryHandler) Update(c echo.Context) error {
	var ProCategory model.ProductCategory
	if err := c.Bind(&ProCategory); err != nil {
		return c.JSON(http.StatusBadRequest, utils.CreateResponse(http.StatusBadRequest, "dữ liệu không hợp lệ", utils.EmptyObject{}))
	}
	if err := ProductCategoryService.Update(&ProCategory); err != nil {
		return c.JSON(http.StatusBadRequest, utils.CreateResponse(http.StatusBadRequest, err.Error(), utils.EmptyObject{}))
	}
	return c.JSON(http.StatusOK, utils.CreateResponse(http.StatusOK, "cập nhật dữ liệu thành công", utils.EmptyObject{}))
}

func (ProductCategoryHandler ProductCategoryHandler) Delete(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.CreateResponse(http.StatusBadRequest, "id không hơp lệ", utils.EmptyObject{}))
	}
	if category, err := ProductCategoryService.GetByID(int32(id)); err != nil || category == nil {
		return c.JSON(http.StatusBadRequest, utils.CreateResponse(http.StatusBadRequest, err.Error(), utils.EmptyObject{}))
	}
	if err := ProductCategoryService.Delete(int32(id)); err != nil {
		return c.JSON(http.StatusBadRequest, utils.CreateResponse(http.StatusBadRequest, err.Error(), utils.EmptyObject{}))
	}
	return c.JSON(http.StatusOK, utils.CreateResponse(http.StatusOK, "xóa thành công", utils.EmptyObject{}))
}

func (ProductCategoryHandler ProductCategoryHandler) GetByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.CreateResponse(http.StatusBadRequest, "id không hơp lệ", utils.EmptyObject{}))
	}
	if category, err := ProductCategoryService.GetByID(int32(id)); err != nil || category == nil {
		return c.JSON(http.StatusBadRequest, utils.CreateResponse(http.StatusBadRequest, "không tồn tại loại sản pẩm có id này", utils.EmptyObject{}))
	} else {
		return c.JSON(http.StatusOK, utils.CreateResponse(http.StatusOK, "lấy dữ liệu thành công", category))
	}
}
