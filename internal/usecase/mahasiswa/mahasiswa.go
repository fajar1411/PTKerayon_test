package mahasiswa

import (
	"backend/domain/contract/repocontract"
	"backend/domain/contract/usecasecontract"
	"backend/domain/model"
	"backend/domain/request"
	"backend/domain/request/datareqmodel"
	"backend/helper/validasi"
	"errors"
	"time"
)

type mahasiswaCase struct {
	repo repocontract.RepoMahasiswa
}

func NewUsecaseMahasiswa(rm repocontract.RepoMahasiswa) usecasecontract.UseCaseMahasiswa {
	return &mahasiswaCase{
		repo: rm,
	}
}

// AddMahasiswa implements usecasecontract.UseCaseMahasiswa.
func (mc *mahasiswaCase) AddMahasiswa(newRequest request.MahasiswaRequest) (data model.Mahasiswa, err error) {
	parsetime, errConvtime := time.Parse("2006-01-02 15:04:05", newRequest.TahunKuliah)

	if nimExists := mc.repo.NimExist(newRequest.Nim); nimExists {
		return data, errors.New("data sudah ada")
	}
	if errConvtime != nil {

		return data, errors.New(errConvtime.Error())
	}
	errvalid := validasi.ValidationErrorHandle(newRequest)

	if errvalid != nil {
		return data, errors.New(errvalid.Error())
	}
	tomodel := datareqmodel.RequestToModel(newRequest, parsetime)

	data, err = mc.repo.AddMahasiswa(tomodel)

	if err != nil {
		return data, errors.New(err.Error())
	}
	return data, nil

}
