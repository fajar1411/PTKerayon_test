package repocontract

import (
	"backend/domain/model"
	"backend/domain/request"
)

type RepoMahasiswa interface {
	AddMahasiswa(newMahasiswa model.Mahasiswa) (data model.Mahasiswa, err error)
	NimExist(nim string) bool
}
type RepoDosen interface {
	AddDosen(newDosen request.DosenRequest) (data request.DosenRequest, err error)
	AllDosen() (data []request.DosenRequest, err error)
	NipExist(nip string) bool
}
type RepoMatakuliah interface {
	AddMatakuliah(newmatkul request.MatkulRequest) (data request.MatkulRequest, err error)
	KodematkulExist(kode string) bool
}
