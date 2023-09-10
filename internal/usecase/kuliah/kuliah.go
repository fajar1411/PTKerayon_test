package kuliah

import (
	"backend/domain/contract/repocontract"
	"backend/domain/contract/usecasecontract"
	"backend/domain/request"
	"backend/helper/validasi"
	"errors"
	"log"
	"strconv"
)

type KuliahCase struct {
	repo       repocontract.Repokuliah
	repodosen  repocontract.RepoDosen
	repomhs    repocontract.RepoMahasiswa
	repomatkul repocontract.RepoMatakuliah
}

func NewUsecaseKuliah(repodosen repocontract.RepoDosen, repomhs repocontract.RepoMahasiswa, repomatkul repocontract.RepoMatakuliah, repo repocontract.Repokuliah) usecasecontract.UseCasekuliah {
	return &KuliahCase{
		repodosen:  repodosen,
		repo:       repo,
		repomhs:    repomhs,
		repomatkul: repomatkul,
	}
}

// Addkuliah implements usecasecontract.UseCasekuliah.
func (kc *KuliahCase) Addkuliah(newkuliah request.KuliahRequest) (data request.KuliahRequest, err error) {
	existdosen := make(chan bool)
	test := make(chan *request.KuliahRequest)
	existmatkul := make(chan bool)
	existmhs := make(chan bool)
	go func() {
		existdosen <- kc.repodosen.NipExist(newkuliah.IndukDosen)

	}()
	nipExists := <-existdosen
	if !(nipExists) {
		return data, errors.New("data dosen tidak ada")
	}

	go func() {
		existmatkul <- kc.repomatkul.KodematkulExist(newkuliah.MatkulKode)

	}()
	matkulExists := <-existmatkul
	if !(matkulExists) {
		return data, errors.New("data matkul tidak ada")
	}

	go func() {
		existmhs <- kc.repomhs.NimExist(newkuliah.IndukMhs)

	}()
	kuliahExists := <-existmhs
	if !(kuliahExists) {
		return data, errors.New("data kuliah tidak ada")
	}

	errvalid := validasi.ValidationErrorHandle(newkuliah)

	if errvalid != nil {
		return data, errors.New(errvalid.Error())
	}

	datas, _ := kc.repo.Addkuliah(newkuliah)
	if datas.IndukDosen == "" {
		return data, errors.New("data kuliah tidak ada")
	}

	if datas.IndukMhs == "" {
		return data, errors.New("data kuliah tidak ada")
	}
	if datas.NilaiUts == 0 {
		return data, errors.New("data kuliah tidak ada")
	}
	if datas.NilaiUas == 0 {
		return data, errors.New("data kuliah tidak ada")
	}

	go func() {

		for {
			test <- &datas
			log.Print("ini test uc", test)
		}
	}()
	test2 := *<-test

	return test2, nil
}

// ParamReq implements usecasecontract.UseCasekuliah.
func (kc *KuliahCase) ParamReq(Idmhs string) (data request.ParamReq, err error) {
	data, err = kc.repo.ParamReq(Idmhs)
	// log.Print("uc", data.NamaMhs)
	total := (data.NilaiUas + data.NilaiUts) / 2

	convrt := strconv.Itoa(int(total))
	data.Nilai = convrt

	if err != nil {
		return data, errors.New("data Mahasiswa tidak ada")
	}
	return
}
