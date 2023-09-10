package dosen

import (
	"backend/domain/contract/repocontract"
	"backend/domain/contract/usecasecontract"
	"backend/domain/request"
	"backend/helper/validasi"
	"errors"
	"fmt"
	"log"
)

type DosenCase struct {
	repo repocontract.RepoDosen
}

func NewUsecaseDosen(rd repocontract.RepoDosen) usecasecontract.UseCaseDosen {
	return &DosenCase{
		repo: rd,
	}
}

// AddDosen implements usecasecontract.UseCaseDosen.
func (dc *DosenCase) AddDosen(newRequest request.DosenRequest) (data request.DosenRequest, err error) {
	existDosen := make(chan bool)
	test := make(chan *request.DosenRequest)

	go func() {
		existDosen <- dc.repo.NipExist(newRequest.Nip)
	}()
	if nipExists := <-existDosen; nipExists {
		return data, errors.New("data sudah ada")
	}

	errvalid := validasi.ValidationErrorHandle(newRequest)

	if errvalid != nil {
		return data, errors.New(errvalid.Error())
	}

	datas, _ := dc.repo.AddDosen(newRequest)

	go func() {

		for {
			test <- &datas
			log.Print("ini test uc", test)
		}
	}()
	test2 := *<-test

	return test2, nil
}
func (dc *DosenCase) GetAllDosen() (request []request.DosenRequest, err error) {
	var pesan = make(chan string, len(request))

	data, err := dc.repo.AllDosen()
	go func() {

		for i := range pesan {
			fmt.Println("terima data", i)
		}
	}()
	for _, s := range data {
		fmt.Println("kirim data", s)
		pesan <- s.Nip
	}
	close(pesan)
	if err != nil {
		return request, errors.New(err.Error())
	}
	return data, nil
}
