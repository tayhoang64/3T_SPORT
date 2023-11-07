package dal

import (
	"3T_Shop/config"
	"3T_Shop/model"
)

type ProductCategoryDAL struct {
}

func (ProductCategoryDAL ProductCategoryDAL) Create(productCategory *model.ProductCategory) error {
	db := config.GetConnection()

	if err := db.Create(&productCategory).Error; err != nil {
		return err
	}
	return nil
}

func (ProductCategoryDAL ProductCategoryDAL) Update(productCategory *model.ProductCategory) error {
	db := config.GetConnection()
	var CategoryUpdate model.ProductCategory
	if errGet := db.Model(model.ProductCategory{}).Where(model.ProductCategory{CategoryID: productCategory.CategoryID}).First(&CategoryUpdate).Error; errGet != nil {
		return errGet
	}
	CategoryUpdate = *productCategory
	if errUpdate := db.Save(&CategoryUpdate).Error; errUpdate != nil {
		return errUpdate
	}
	return nil
}

func (ProductCategoryDAL ProductCategoryDAL) Delete(id int32) error {
	db := config.GetConnection()
	var CategoryDelete model.ProductCategory
	if errGet := db.Model(model.ProductCategory{}).Where(model.ProductCategory{CategoryID: id}).First(&CategoryDelete).Error; errGet != nil {
		return errGet
	}
	if errDelete := db.Delete(CategoryDelete).Error; errDelete != nil {
		return errDelete
	}
	return nil
}

func (ProductCategoryDAL ProductCategoryDAL) GetAll() ([]model.ProductCategory, error) {
	db := config.GetConnection()
	var proCategories []model.ProductCategory
	if err := db.Find(&proCategories).Error; err != nil {
		return nil, err
	}
	return proCategories, nil
}
func (ProductCategoryDAL ProductCategoryDAL) GetByID(id int32) (*model.ProductCategory, error) {
	db := config.GetConnection()
	var proCategory model.ProductCategory
	if errGet := db.Model(model.ProductCategory{}).Where(model.ProductCategory{CategoryID: id}).First(&proCategory).Error; errGet != nil {
		return nil, errGet
	}
	return &proCategory, nil
}
func (ProductCategoryDAL ProductCategoryDAL) GetBySlug(slug string) (*model.ProductCategory, error) {
	db := config.GetConnection()
	var proCategory model.ProductCategory
	if errGet := db.Model(model.ProductCategory{}).Where(model.ProductCategory{Slug: slug}).First(&proCategory).Error; errGet != nil {
		return nil, errGet
	}
	return &proCategory, nil
}
