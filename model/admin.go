package model

import (
	"deptechdigital/entity"

	"gorm.io/gorm"
)

type AdminModel struct {
	DB *gorm.DB
}

func (um AdminModel) LogIn(data entity.Admin) (entity.Admin, error) {
	var res entity.Admin
	err := um.DB.Where("email = ?", data.Email).Find(&res).Error
	if err != nil {
		return entity.Admin{}, err
	}
	return res, nil
}

func (um AdminModel) Insert(newData entity.Admin) (entity.Admin, error) {
	err := um.DB.Create(&newData).Error
	if err != nil {
		return entity.Admin{}, err
	}
	return newData, nil
}

func (um AdminModel) Update(newData entity.Admin) (entity.Admin, error) {
	err := um.DB.Where("id=?", newData.ID).Updates(&newData).Error
	if err != nil {
		return entity.Admin{}, err
	}

	return newData, nil
}

func (um AdminModel) Delete(id uint) error {
	err := um.DB.Where("id=?", id).Delete(&entity.Admin{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (um AdminModel) GetAll() ([]entity.Admin, error) {
	var res []entity.Admin
	err := um.DB.Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}
