package usecasecontract

import (
	"backend/domain/model"
	"backend/domain/request"
)

type UseCaseMahasiswa interface {
	AddMahasiswa(newRequest request.MahasiswaRequest) (data model.Mahasiswa, err error)
}

type UseCaseDosen interface {
	AddDosen(newRequest request.DosenRequest) (data request.DosenRequest, err error)
	GetAllDosen() (request []request.DosenRequest, err error)
}
type UseCaseMatkul interface {
	AddMatkul(newRequest request.MatkulRequest) (data request.MatkulRequest, err error)
}
type UseCasekuliah interface {
	Addkuliah(newkuliah request.KuliahRequest) (data request.KuliahRequest, err error)
	ParamReq(Idmhs string) (data request.ParamReq, err error)
}
