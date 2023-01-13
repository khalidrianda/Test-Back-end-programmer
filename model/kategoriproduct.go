package model

import (
	"deptechdigital/entity"

	"gorm.io/gorm"
)

type KategoriModel struct {
	DB *gorm.DB
}

func (um KategoriModel) InsertKategori(newData entity.KategoriProduct) (entity.KategoriProduct, error) {
	err := um.DB.Create(&newData).Error
	if err != nil {
		return entity.KategoriProduct{}, err
	}
	return newData, nil
}

func (um KategoriModel) UpdateKategori(newData entity.KategoriProduct) (entity.KategoriProduct, error) {
	err := um.DB.Where("id=?", newData.ID).Updates(&newData).Error
	if err != nil {
		return entity.KategoriProduct{}, err
	}

	return newData, nil
}

func (um KategoriModel) DeleteKategori(id uint) error {
	err := um.DB.Where("id=?", id).Delete(&entity.KategoriProduct{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (um KategoriModel) GetAllKategori() ([]entity.KategoriProduct, error) {
	var res []entity.KategoriProduct
	err := um.DB.Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}
