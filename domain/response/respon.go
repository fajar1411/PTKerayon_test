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
type KuliahRespon struct {
	MatkulKode string `json:"kd_matakuliah"`
	IndukDosen string `json:"nip"`
	IndukMhs   string `json:"nim"`
	NilaiUts   uint   `json:"uts"`
	NilaiUas   uint   `json:"uas"`
}
type ParamRes struct {
	NamaMataKuliah string `json:"matkul"`
	KotaAsal       string `json:"kota"`
	NamaDosen      string `json:"nama_dosen"`
	NamaMhs        string `json:"nama_mahasiswa"`
	NilaiUts       uint   `json:"uts"`
	NilaiUas       uint   `json:"uas"`
	Nilai          string `json:"total_nilai"`
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
func KuliahRequestoFormResponse(data request.KuliahRequest) KuliahRespon {
	return KuliahRespon{
		NilaiUts:   data.NilaiUts,
		NilaiUas:   data.NilaiUas,
		IndukDosen: data.IndukDosen,
		IndukMhs:   data.IndukMhs,
		MatkulKode: data.MatkulKode,
	}
}
func ParamRequestoFormResponse(data request.ParamReq) ParamRes {
	return ParamRes{
		NilaiUts:       data.NilaiUts,
		NilaiUas:       data.NilaiUas,
		NamaDosen:      data.NamaDosen,
		NamaMhs:        data.NamaMahasiswa,
		Nilai:          data.Nilai,
		NamaMataKuliah: data.NamaMatakuliah,
		KotaAsal:       data.KotaAsal,
	}
}
