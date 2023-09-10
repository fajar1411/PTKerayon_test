package model

import (
	"time"
)

type Mahasiswa struct {
	Nim           string `gorm:"primaryKey"`
	NamaMahasiswa string
	KotaAsal      string
	TahunKuliah   time.Time
	Kuliah        []Kuliah `gorm:"foreignKey:IndukMhs;references:Nim"`
}
type Dosen struct {
	Nip        string `gorm:"primaryKey"`
	NamaDosen  string
	KotaAsal   string
	MataKuliah []MataKuliah `gorm:"foreignKey:DosenNip;references:Nip"`
	Kuliah     []Kuliah     `gorm:"foreignKey:IndukDosen;references:Nip"`
}

type MataKuliah struct {
	KodeMatakuliah string `gorm:"primaryKey"`
	NamaMatakuliah string
	JumlahSks      uint
	DosenNip       string   `gorm:"size:191"`
	Kuliah         []Kuliah `gorm:"foreignKey:MatkulKode;references:KodeMatakuliah"`
}

type Kuliah struct {
	KodeKuliah string `gorm:"primaryKey"`
	MatkulKode string `gorm:"size:191"`
	IndukDosen string `gorm:"size:191"`
	IndukMhs   string `gorm:"size:191"`
	NilaiUts   uint
	NilaiUas   uint
}
