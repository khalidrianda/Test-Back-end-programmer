package entity

import "gorm.io/gorm"

type Admin struct {
	ID           uint   `json:"id" form:"id"`
	NamaDepan    string `json:"nama_depan" form:"nama_depan" validate:"min=4,max=15"`
	NamaBelakang string `json:"nama_belakang" form:"nama_belakang" validate:"min=4,max=15"`
	Email        string `json:"email" form:"email" validate:"email" gorm:"unique"`
	TanggalLahir string `json:"tanggal_lahir" form:"tanggal_lahir" validate:"min=4,max=16"`
	JenisKelamin string `json:"jenis_kelamin" form:"jenis_kelamin" validate:"min=4,max=10"`
	Password     string `json:"password" form:"password"`
}

type KategoriProduct struct {
	ID                uint   `json:"id" form:"id"`
	NamaKategori      string `json:"nama_kategori" form:"nama_kategori" validate:"min=4,max=30"`
	DeskripsiKategori string `json:"deskripsi_kategori" validate:"min=4,max=30"`
}

type Product struct {
	ID                 uint                `json:"id" form:"id"`
	NamaProduct        string              `json:"nama_product" validate:"min=4,max=30"`
	DeskripsiProduct   string              `json:"deskripsi_product" validate:"min=4,max=30"`
	GambarProduct      string              `json:"gambar_product" validate:"min=4,max=30"`
	KategoriProduct    string              `json:"kategori_product" validate:"min=4,max=30"`
	StokProduct        int                 `json:"stok_product" validate:"min=1,max=30"`
	TransactionDetails []TransactionDetail `json:"-" gorm:"foreignKey:ProductID"`
}

type Transaction struct {
	gorm.Model
	Keterangan         string              `json:"keterangan"`
	TransactionDetails []TransactionDetail `json:"transaction_detail" gorm:"foreignKey:TransactionID"`
}

type TransactionDetail struct {
	gorm.Model    `json:"-"`
	TransactionID uint `json:"-"`
	ProductID     uint `json:"product_id"`
	Jumlah        int  `json:"jumlah"`
}

type TransactionFormat struct {
	Keterangan         string              `json:"keterangan"`
	TransactionDetails []TransactionDetail `json:"transaction_detail" gorm:"foreignKey:TransactionID"`
}

type TransactionDetailFormat struct {
	TransactionID uint
	ProductID     uint `json:"product_id"`
	Jumlah        int  `json:"jumlah"`
}

func FromDomain(s TransactionFormat) Transaction {
	return Transaction{
		Keterangan: s.Keterangan,
	}
}

func FromDomainDetail(s TransactionFormat, id uint) []TransactionDetail {
	var res []TransactionDetail
	for _, val := range s.TransactionDetails {
		res = append(res, TransactionDetail{TransactionID: id, ProductID: val.ProductID, Jumlah: val.Jumlah})
	}
	return res
}
func ToDomain(s Transaction, t []TransactionDetail) Transaction {
	var res Transaction = Transaction{Model: gorm.Model{ID: s.ID, CreatedAt: s.CreatedAt, UpdatedAt: s.UpdatedAt}, Keterangan: s.Keterangan}
	for _, val := range t {
		res.TransactionDetails = append(res.TransactionDetails, TransactionDetail{ProductID: val.ProductID, Jumlah: val.Jumlah})
	}
	return res
}

// type TransactionFormat struct {
// 	Keterangan        string              `json:"keterangan"`
// 	TransactionDetail []TransactionDetail `json:"transaction_detail"`
// }
