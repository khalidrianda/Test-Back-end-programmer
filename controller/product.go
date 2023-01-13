package controller

import (
	"deptechdigital/entity"
	"deptechdigital/model"
)

type ProductControll struct {
	Model model.ProductModel
}

func (kc ProductControll) AddProduct(data entity.Product) (entity.Product, error) {
	res, err := kc.Model.InsertProduct(data)
	if err != nil {
		return entity.Product{}, err
	}
	return res, nil
}

func (uc ProductControll) EditProduct(data entity.Product) (entity.Product, error) {
	res, err := uc.Model.UpdateProduct(data)
	if err != nil {
		return entity.Product{}, err
	}
	return res, nil
}

func (uc ProductControll) RemoveProduct(id uint) error {
	err := uc.Model.DeleteProduct(id)
	if err != nil {
		return err
	}
	return nil
}

func (uc ProductControll) ShowAllProduct() ([]entity.Product, error) {
	res, err := uc.Model.GetAllProduct()
	if err != nil {
		return nil, err
	}
	return res, nil
}
