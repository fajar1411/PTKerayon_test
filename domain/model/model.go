package model

import (
	"time"
)

type Mahasiswa struct {
	Nim           string `gorm:"primaryKey"`
	NamaMahasiswa string
	KotaAsal      string
	TahunKuliah   time.Time
}
type Dosen struct {
	Nip        string `gorm:"primaryKey"`
	NamaDosen  string
	KotaAsal   string
	MataKuliah []MataKuliah `gorm:"foreignKey:DosenNip;references:Nip"`
}

type MataKuliah struct {
	KodeMatakuliah string `gorm:"primaryKey"`
	NamaMatakuliah string
	JumlahSks      uint
	DosenNip       string `gorm:"size:191"`
}
