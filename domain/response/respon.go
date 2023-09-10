package response

import (
	"backend/domain/model"
	"backend/domain/request"
	"time"
)

type MahasiswaRespons struct {
	NamaMahasiswa string    `json:"nama_mahasiswa"`
	Nim           string    `json:"nim"`
	KotaAsal      string    `json:"kota"`
	TahunKuliah   time.Time `json:"tahun_masuk"`
}
type DosenRespons struct {
	NamaDosen string `json:"nama_Dosen"`
	Nip       string `json:"nip"`
	KotaAsal  string `json:"kota"`
}
type MatkulRespons struct {
	KodeMataKuliah string `json:"kd_matakuliah"`
	NamaMatakuliah string `json:"nama_matakuliah"`
	JumlahSks      uint   `json:"sks"`
	DosenNip       string `json:"Nip"`
}

func ToFormResponse(data model.Mahasiswa) MahasiswaRespons {
	return MahasiswaRespons{

		NamaMahasiswa: data.NamaMahasiswa,
		Nim:           data.Nim,
		KotaAsal:      data.KotaAsal,
		TahunKuliah:   data.TahunKuliah,
	}
}
func RequestoFormResponse(data request.DosenRequest) DosenRespons {
	return DosenRespons{
		NamaDosen: data.NamaDosen,
		Nip:       data.Nip,
		KotaAsal:  data.KotaAsal,
	}
}
func ListReqToRespon(data []request.DosenRequest) (datares []DosenRespons) {
	for _, val := range data {
		datares = append(datares, RequestoFormResponse(val))
	}
	return datares
}
func MatkulRequestoFormResponse(data request.MatkulRequest) MatkulRespons {
	return MatkulRespons{
		KodeMataKuliah: data.KodeMatakuliah,
		NamaMatakuliah: data.NamaMatakuliah,
		DosenNip:       data.Nip,
		JumlahSks:      data.JumlahSks,
	}
}
