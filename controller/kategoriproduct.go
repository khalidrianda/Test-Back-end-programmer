package controller

import (
	"deptechdigital/entity"
	"deptechdigital/model"
)

type KategoriControll struct {
	Model model.KategoriModel
}

func (kc KategoriControll) AddKategori(data entity.KategoriProduct) (entity.KategoriProduct, error) {
	res, err := kc.Model.InsertKategori(data)
	if err != nil {
		return entity.KategoriProduct{}, err
	}
	return res, nil
}

func (uc KategoriControll) EditKategori(data entity.KategoriProduct) (entity.KategoriProduct, error) {
	res, err := uc.Model.UpdateKategori(data)
	if err != nil {
		return entity.KategoriProduct{}, err
	}
	return res, nil
}

func (uc KategoriControll) RemoveKategori(id uint) error {
	err := uc.Model.DeleteKategori(id)
	if err != nil {
		return err
	}
	return nil
}

func (uc KategoriControll) ShowAllKategori() ([]entity.KategoriProduct, error) {
	res, err := uc.Model.GetAllKategori()
	if err != nil {
		return nil, err
	}
	return res, nil
}
