package service

import (
	"3T_Shop/dal"
	"3T_Shop/model"
	"errors"
	"strings"
)

type RoleService struct {
}

var (
	RoleDAL dal.RoleDAL
)

func (RoleService RoleService) Create(role *model.Role) error {
	if len(strings.TrimSpace(role.RoleID)) == 0 {
		return errors.New("Phải nhập role ID")
	}
	if len(strings.TrimSpace(role.RoleName)) == 0 {
		return errors.New("Phải nhập role name")
	}
	if role, _ := RoleDAL.GetByID(role.RoleID); role != nil {
		return errors.New("RoleID đã tồn tại")
	}
	if err := RoleDAL.Create(role); err != nil {
		return errors.New("lỗi khi thêm role")
	}
	return nil
}

func (RoleService RoleService) GetAll() ([]model.Role, error) {
	return RoleDAL.GetAll()
}
