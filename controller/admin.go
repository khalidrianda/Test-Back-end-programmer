package controller

import (
	"deptechdigital/entity"
	"deptechdigital/model"
)

type AdminControll struct {
	Model model.AdminModel
}

func (uc AdminControll) LogIn(data entity.Admin) (entity.Admin, error) {
	res, err := uc.Model.LogIn(data)
	if err != nil {
		return entity.Admin{}, err
	}
	return res, nil
}

func (uc AdminControll) Add(data entity.Admin) (entity.Admin, error) {
	res, err := uc.Model.Insert(data)
	if err != nil {
		return entity.Admin{}, err
	}
	return res, nil
}

func (uc AdminControll) Edit(data entity.Admin) (entity.Admin, error) {
	res, err := uc.Model.Update(data)
	if err != nil {
		return entity.Admin{}, err
	}
	return res, nil
}

func (uc AdminControll) Remove(id uint) error {
	err := uc.Model.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (uc AdminControll) ShowAll() ([]entity.Admin, error) {
	res, err := uc.Model.GetAll()
	if err != nil {
		return nil, err
	}
	return res, nil
}
