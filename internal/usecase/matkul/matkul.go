package matkul

import (
	"backend/domain/contract/repocontract"
	"backend/domain/contract/usecasecontract"
	"backend/domain/request"
	"backend/helper/validasi"
	"errors"
	"log"
)

type MatkulCase struct {
	repo      repocontract.RepoMatakuliah
	repodosen repocontract.RepoDosen
}

func NewUsecaseMatkul(rd repocontract.RepoDosen, repo repocontract.RepoMatakuliah) usecasecontract.UseCaseMatkul {
	return &MatkulCase{
		repodosen: rd,
		repo:      repo,
	}
}

func (mc *MatkulCase) AddMatkul(newRequest request.MatkulRequest) (data request.MatkulRequest, err error) {
	existdosen := make(chan bool)
	test := make(chan *request.MatkulRequest)
	existmatkul := make(chan bool)
	go func() {
		existdosen <- mc.repodosen.NipExist(newRequest.Nip)

	}()
	nipExists := <-existdosen
	if !(nipExists) {
		return data, errors.New("data dosen tidak ada")
	}

	go func() {
		existmatkul <- mc.repo.KodematkulExist(newRequest.KodeMatakuliah)

	}()
	matkulExists := <-existmatkul
	if matkulExists {
		return data, errors.New("data matkul sudah ada")
	}

	// if !(nipExists && matkulExists) {
	// 	return data, errors.New("data tidak  ada")
	// }
	errvalid := validasi.ValidationErrorHandle(newRequest)

	if errvalid != nil {
		return data, errors.New(errvalid.Error())
	}

	datas, _ := mc.repo.AddMatakuliah(newRequest)

	go func() {

		for {
			test <- &datas
			log.Print("ini test uc", test)
		}
	}()
	test2 := *<-test

	return test2, nil
}
