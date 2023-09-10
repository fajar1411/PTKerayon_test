package matkul

import (
	"backend/domain/contract/repocontract"
	"backend/domain/model"
	"backend/domain/request"
	"backend/domain/request/datareqmodel"
	"errors"

	"gorm.io/gorm"
)

type MatkulRepo struct {
	db *gorm.DB
}

func NewRepoMatkul(db *gorm.DB) repocontract.RepoMatakuliah {
	return &MatkulRepo{
		db: db,
	}
}

// AddMatakuliah implements repocontract.RepoMatakuliah.
func (mr *MatkulRepo) AddMatakuliah(newmatkul request.MatkulRequest) (data request.MatkulRequest, err error) {
	tomodel := datareqmodel.MatkulRequestToModel(newmatkul)
	tx := mr.db.Create(&tomodel)

	if tx.Error != nil {
		return data, errors.New("activities query error")
	}

	data.KodeMatakuliah = newmatkul.KodeMatakuliah
	data.JumlahSks = newmatkul.JumlahSks
	data.Nip = newmatkul.Nip
	data.NamaMatakuliah = newmatkul.NamaMatakuliah
	return
}

// KodematkulExist implements repocontract.RepoMatakuliah.
func (mr *MatkulRepo) KodematkulExist(kode string) bool {
	var datamodel model.MataKuliah

	tx := mr.db.First(&datamodel, "kode_matakuliah = ?", kode)
	return tx.Error == nil
}

// // AddDosen implements repocontract.RepoDosen.
// func (dr *DosenRepo) AddDosen(newDosen request.DosenRequest) (data request.DosenRequest, err error) {
// 	tomodel := datareqmodel.DosenRequestToModel(newDosen)
// 	tx := dr.db.Create(&tomodel)

// 	if tx.Error != nil {
// 		return data, errors.New("activities query error")
// 	}

// 	data.NamaDosen = tomodel.NamaDosen
// 	data.Nip = tomodel.Nip
// 	data.KotaAsal = tomodel.KotaAsal
// 	return
// }
