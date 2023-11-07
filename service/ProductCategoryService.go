package service

import (
	"3T_Shop/dal"
	"3T_Shop/model"
	"errors"
	"github.com/gosimple/slug"
	"strconv"
	"strings"
	"time"
)

type ProductCategoryService struct {
}

var (
	ProductCategoryDAL dal.ProductCategoryDAL
)

func (ProductCategoryService ProductCategoryService) Create(ProCategory *model.ProductCategory) error {
	if len(strings.TrimSpace(ProCategory.CategoryName)) == 0 {
		return errors.New("Phải nhập tên loại sản phẩm")
	}
	sl := slug.Make(ProCategory.CategoryName)
	if productCategory, _ := ProductCategoryDAL.GetBySlug(sl); productCategory != nil {
		ProCategory.Slug = slug.Make(ProCategory.CategoryName + " " + strconv.FormatInt(time.Now().Unix(), 10))
	} else {
		ProCategory.Slug = sl
	}

	if err := ProductCategoryDAL.Create(ProCategory); err != nil {
		return errors.New("Lỗi khi tạo loại sản phẩm")
	}
	return nil
}

func (ProductCategoryService ProductCategoryService) Update(ProCategory *model.ProductCategory) error {
	if len(strings.TrimSpace(ProCategory.CategoryName)) == 0 {
		return errors.New("Phải nhập tên loại sản phẩm")
	}
	sl := slug.Make(ProCategory.CategoryName)
	if productCategory, _ := ProductCategoryDAL.GetBySlug(sl); productCategory != nil {
		ProCategory.Slug = slug.Make(ProCategory.CategoryName + " " + strconv.FormatInt(time.Now().Unix(), 10))
	} else {
		ProCategory.Slug = sl
	}

	if err := ProductCategoryDAL.Update(ProCategory); err != nil {
		return errors.New("Lỗi khi tạo loại sản phẩm")
	}
	return nil
}

func (ProductCategoryService ProductCategoryService) Delete(id int32) error {
	if category, err := ProductCategoryDAL.GetByID(id); category == nil || err != nil {
		return errors.New("không tìm thấy sản loai sản phẩm cần xóa")
	}
	if err := ProductCategoryDAL.Delete(id); err != nil {
		return errors.New("lỗi khi cập nhật dữ liệu")
	}
	return nil
}

func (ProductCategoryService ProductCategoryService) GetAll() ([]model.ProductCategory, error) {
	return ProductCategoryDAL.GetAll()
}

func (ProductCategoryService ProductCategoryService) GetByID(id int32) (*model.ProductCategory, error) {
	return ProductCategoryDAL.GetByID(id)
}
