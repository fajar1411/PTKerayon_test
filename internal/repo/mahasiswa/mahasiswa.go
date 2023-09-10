package mahasiswa

import (
	"backend/domain/contract/repocontract"
	"backend/domain/model"
	"errors"

	"gorm.io/gorm"
)

type MahasiswaRepo struct {
	db *gorm.DB
}

func NewRepoMahasiswa(db *gorm.DB) repocontract.RepoMahasiswa {
	return &MahasiswaRepo{
		db: db,
	}
}

// AddMahasiswa implements repocontract.RepoMahasiswa.
func (mr *MahasiswaRepo) AddMahasiswa(newMahasiswa model.Mahasiswa) (data model.Mahasiswa, err error) {
	tx := mr.db.Create(&newMahasiswa)

	if tx.Error != nil {
		return data, errors.New("activities query error")
	}
	data = newMahasiswa
	return
}

// NimExist implements repocontract.RepoMahasiswa.
func (mr *MahasiswaRepo) NimExist(nim string) bool {
	var datamodel model.Mahasiswa

	tx := mr.db.First(&datamodel, "nim = ?", nim)
	return tx.Error == nil
}
