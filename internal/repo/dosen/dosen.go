package dosen

import (
	"backend/domain/contract/repocontract"
	"backend/domain/model"
	"backend/domain/request"
	"backend/domain/request/datareqmodel"
	"errors"

	"gorm.io/gorm"
)

type DosenRepo struct {
	db *gorm.DB
}

func NewRepoDosen(db *gorm.DB) repocontract.RepoDosen {
	return &DosenRepo{
		db: db,
	}
}

// AddDosen implements repocontract.RepoDosen.
func (dr *DosenRepo) AddDosen(newDosen request.DosenRequest) (data request.DosenRequest, err error) {
	tomodel := datareqmodel.DosenRequestToModel(newDosen)
	tx := dr.db.Create(&tomodel)

	if tx.Error != nil {
		return data, errors.New("activities query error")
	}

	data.NamaDosen = tomodel.NamaDosen
	data.Nip = tomodel.Nip
	data.KotaAsal = tomodel.KotaAsal
	return
}

// NipExist implements repocontract.RepoDosen.
func (dr *DosenRepo) NipExist(nip string) bool {
	var datamodel model.Dosen

	tx := dr.db.First(&datamodel, "nip = ?", nip)
	return tx.Error == nil
}

// AllDosen implements repocontract.RepoDosen.
func (dr *DosenRepo) AllDosen() (data []request.DosenRequest, err error) {
	var activ []model.Dosen
	tx := dr.db.Raw("Select dosens.nama_dosen,dosens.nip, dosens.kota_asal from dosens").Find(&activ)
	if tx.Error != nil {
		return data, errors.New("activities query error")
	}
	dtmdlttoreq := datareqmodel.ListModelToReq(activ)
	return dtmdlttoreq, nil
}
