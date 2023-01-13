package model

import (
	"deptechdigital/entity"
	"errors"

	"gorm.io/gorm"
)

type TransactionModel struct {
	DB *gorm.DB
}

func (um TransactionModel) InsertTransaction(newData entity.TransactionFormat) (entity.Transaction, error) {
	var temp entity.Product
	cnvData := entity.FromDomain(newData)

	for _, val := range newData.TransactionDetails {
		um.DB.Where("id=?", val.ProductID).Find(&temp)
		if temp.StokProduct < val.Jumlah {
			return cnvData, errors.New("barang tidak cukup")
		}
	}

	err := um.DB.Create(&cnvData).Error
	if err != nil {
		return entity.Transaction{}, err
	}

	cnvDetail := entity.FromDomainDetail(newData, cnvData.ID)
	err = um.DB.Create(&cnvDetail).Error
	if err != nil {
		return entity.Transaction{}, err
	}

	for _, val := range cnvDetail {
		um.DB.Where("id=?", val.ProductID).Find(&temp)
		if cnvData.Keterangan == "Stock In" {
			temp.StokProduct += val.Jumlah
		} else if cnvData.Keterangan == "Stock Out" {
			temp.StokProduct -= val.Jumlah
		}
		um.DB.Where("id=?", val.ProductID).Updates(&temp)
	}

	cnvData = entity.ToDomain(cnvData, cnvDetail)
	// log.Print(cnvDetail)
	return cnvData, nil
}

func (um TransactionModel) GetAllTransaction() ([]entity.Transaction, error) {
	var res []entity.Transaction
	err := um.DB.Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (um TransactionModel) GetTransactionDetail(id uint) (entity.Transaction, error) {
	var res entity.Transaction
	var resDetail []entity.TransactionDetail
	err := um.DB.Where("id=?", id).Find(&res).Error
	if err != nil {
		return entity.Transaction{}, err
	}

	err = um.DB.Where("transaction_id=?", id).Find(&resDetail).Error
	if err != nil {
		return entity.Transaction{}, err
	}

	res = entity.ToDomain(res, resDetail)

	return res, nil
}
