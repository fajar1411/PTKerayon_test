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
