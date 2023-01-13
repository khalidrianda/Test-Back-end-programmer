package controller

import (
	"deptechdigital/entity"
	"deptechdigital/model"
)

type TransactionControll struct {
	Model model.TransactionModel
}

func (kc TransactionControll) AddTransaction(data entity.TransactionFormat) (entity.Transaction, error) {
	res, err := kc.Model.InsertTransaction(data)
	if err != nil {
		return entity.Transaction{}, err
	}
	return res, nil
}

func (uc TransactionControll) ShowAllTransaction() ([]entity.Transaction, error) {
	res, err := uc.Model.GetAllTransaction()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (uc TransactionControll) ShowTransactionDetail(id uint) (entity.Transaction, error) {
	res, err := uc.Model.GetTransactionDetail(id)
	if err != nil {
		return entity.Transaction{}, err
	}
	return res, nil
}
