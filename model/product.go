package model

import (
	"deptechdigital/entity"

	"gorm.io/gorm"
)

type ProductModel struct {
	DB *gorm.DB
}

func (um ProductModel) InsertProduct(newData entity.Product) (entity.Product, error) {
	err := um.DB.Create(&newData).Error
	if err != nil {
		return entity.Product{}, err
	}
	return newData, nil
}

func (um ProductModel) UpdateProduct(newData entity.Product) (entity.Product, error) {
	err := um.DB.Where("id=?", newData.ID).Updates(&newData).Error
	if err != nil {
		return entity.Product{}, err
	}

	return newData, nil
}

func (um ProductModel) DeleteProduct(id uint) error {
	err := um.DB.Where("id=?", id).Delete(&entity.Product{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (um ProductModel) GetAllProduct() ([]entity.Product, error) {
	var res []entity.Product
	err := um.DB.Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (um ProductModel) TransaksiProduct(newData entity.Product) (entity.Product, error) {
	err := um.DB.Where("id=?", newData.ID).Updates(&newData).Error
	if err != nil {
		return entity.Product{}, err
	}

	return newData, nil
}
