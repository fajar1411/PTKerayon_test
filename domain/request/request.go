package request

type MahasiswaRequest struct {
	NamaMahasiswa string `json:"name" form:"name" validate:"required,min=5"`
	Nim           string `json:"nim" form:"nim" validate:"required,min=5"`
	KotaAsal      string `json:"kota" form:"kota" validate:"required,min=5"`
	TahunKuliah   string `json:"tahun" form:"tahun" validate:"required"`
}
type DosenRequest struct {
	NamaDosen string `json:"name" form:"name" validate:"required,min=5"`
	Nip       string `json:"nip" form:"nip" validate:"required,min=5"`
	KotaAsal  string `json:"kota" form:"kota" validate:"required,min=5"`
}
type MatkulRequest struct {
	NamaMatakuliah string `json:"name" form:"name" validate:"required,min=5"`
	Nip            string `json:"nip" form:"nip" validate:"required,min=5"`
	KodeMatakuliah string `json:"kode" form:"kode" validate:"required,min=5"`
	JumlahSks      uint   `json:"sks kode" form:"sks" validate:"required,min=1"`
}
type KuliahRequest struct {
	KodeKuliah string `json:"kode_kuliah" form:"kode_kuliah" validate:"required,min=5"`
	MatkulKode string `json:"kode" form:"kode" validate:"required,min=5"`
	IndukDosen string `json:"nip" form:"nip" validate:"required,min=5"`
	IndukMhs   string `json:"nim" form:"nim" validate:"required,min=5"`
	NilaiUts   uint   `json:"uts" form:"uts" validate:"required,min=2"`
	NilaiUas   uint   `json:"uas" form:"uas" validate:"required,min=2"`
}

type ParamReq struct {
	NilaiUts       uint   `json:"uts" form:"uts" validate:"required,min=2"`
	NilaiUas       uint   `json:"uas" form:"uas" validate:"required,min=2"`
	NamaMahasiswa  string `json:"namamhs" form:"namamhs" validate:"required,min=5"`
	KotaAsal       string `json:"kota" form:"kota" validate:"required,min=5"`
	NamaMatakuliah string `json:"nama_matkul" form:"nama_matkul" validate:"required,min=5"`
	NamaDosen      string `json:"namadsn" form:"namadsn" validate:"required,min=5"`
	Nilai          string `json:"nilai" form:"nilai" validate:"required,min=2"`
}
