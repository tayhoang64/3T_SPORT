package dal

import (
	"3T_Shop/config"
	"3T_Shop/model"
)

type RoleDAL struct {
}

func (RoleDAL RoleDAL) Create(role *model.Role) error {
	db := config.GetConnection()
	err := db.Create(&role).Error
	if err != nil {
		return err
	}
	return nil
}

func (RoleDAL RoleDAL) Update(role *model.Role) error {
	db := config.GetConnection()
	var RoleUpdate model.Role
	errGet := db.Model(model.Role{}).Where(model.Role{RoleID: role.RoleID}).First(RoleUpdate).Error
	if errGet != nil {
		return errGet
	}
	RoleUpdate = *role
	errUpdate := db.Save(RoleUpdate).Error
	if errUpdate != nil {
		return errUpdate
	}
	return nil
}

func (RoleDAL RoleDAL) Delete(id string) error {
	db := config.GetConnection()
	var RoleDelete model.Role
	errGet := db.Model(model.Role{}).Where(model.Role{RoleID: id}).First(RoleDelete).Error
	if errGet != nil {
		return errGet
	}
	errDelete := db.Delete(RoleDelete).Error
	if errDelete != nil {
		return errDelete
	}
	return nil
}

func (RoleDAL RoleDAL) GetAll() ([]model.Role, error) {
	db := config.GetConnection()
	var roles []model.Role
	if err := db.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (RoleDAL RoleDAL) GetByID(id string) (*model.Role, error) {
	db := config.GetConnection()
	var role model.Role
	if err := db.Model(model.Role{}).Where(model.Role{RoleID: id}).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}
