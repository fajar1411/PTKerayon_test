package kuliah

import (
	"backend/domain/contract/repocontract"
	"backend/domain/model"
	"backend/domain/request"
	"backend/domain/request/datareqmodel"
	"errors"

	"gorm.io/gorm"
)

type KuliahRepo struct {
	db *gorm.DB
}

func NewRepoKuliah(db *gorm.DB) repocontract.Repokuliah {
	return &KuliahRepo{
		db: db,
	}
}

// Addkuliah implements repocontract.Repokuliah.
func (kr *KuliahRepo) Addkuliah(newkuliah request.KuliahRequest) (data request.KuliahRequest, err error) {
	tomodel := datareqmodel.KuliahRequestToModel(newkuliah)
	tx := kr.db.Create(&tomodel)

	if tx.Error != nil {
		return data, errors.New("activities query error")
	}

	data.IndukDosen = tomodel.IndukDosen
	data.IndukMhs = tomodel.IndukMhs
	data.MatkulKode = tomodel.MatkulKode
	data.NilaiUas = tomodel.NilaiUas
	data.NilaiUts = tomodel.NilaiUts
	data.KodeKuliah = tomodel.KodeKuliah
	return
}

func (kr *KuliahRepo) KuliahExist(kode string) bool {
	var datamodel model.MataKuliah

	tx := kr.db.First(&datamodel, "id = ?", kode)
	return tx.Error == nil
}

// ParamReq implements repocontract.Repokuliah.
func (kr *KuliahRepo) ParamReq(Idmhs string) (data request.ParamReq, err error) {

	tx := kr.db.Raw("SELECT k.nilai_uts, k.nilai_uas ,m.nama_mahasiswa,m.kota_asal,mk.nama_matakuliah, d.nama_dosen FROM kuliahs k, mahasiswas m,dosens d , mata_kuliahs mk where k.matkul_kode = mk.kode_matakuliah and k.induk_mhs = m.nim and k.induk_dosen = d.nip and k.induk_mhs = ?", Idmhs).Find(&data)

	if tx.Error != nil {
		return data, errors.New("activities query error")
	}
	return
}
