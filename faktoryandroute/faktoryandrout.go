package faktoryandroute

import (
	"backend/internal/handler/dosen"
	"backend/internal/handler/kuliah"
	"backend/internal/handler/mahasiswa"
	"backend/internal/handler/matkul"
	rd "backend/internal/repo/dosen"
	rk "backend/internal/repo/kuliah"
	rm "backend/internal/repo/mahasiswa"
	rmk "backend/internal/repo/matkul"
	ud "backend/internal/usecase/dosen"
	uk "backend/internal/usecase/kuliah"
	um "backend/internal/usecase/mahasiswa"
	umk "backend/internal/usecase/matkul"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func FaktoryAndRoute(e *echo.Echo, db *gorm.DB) {
	rpm := rm.NewRepoMahasiswa(db)
	ucmhsw := um.NewUsecaseMahasiswa(rpm)
	hndlmhs := mahasiswa.NewHandleMahasiswa(ucmhsw)
	mahasiswagroup := e.Group("/mahasiswa")
	mahasiswagroup.POST("/addmahasiswa", hndlmhs.AddMahasiswa)

	rpd := rd.NewRepoDosen(db)
	ucd := ud.NewUsecaseDosen(rpd)
	handledosen := dosen.NewHandleDosen(ucd)
	dosengrp := e.Group("/dosen")
	dosengrp.POST("/adddosen", handledosen.AddDosen)
	dosengrp.GET("", handledosen.GetAllDosen)

	rpmk := rmk.NewRepoMatkul(db)
	umks := umk.NewUsecaseMatkul(rpd, rpmk)
	hndlmk := matkul.NewHandleMatkul(umks)
	matkulgrp := e.Group("/matkul")
	matkulgrp.POST("/addmatkul", hndlmk.AddMatakuliah)

	rp := rk.NewRepoKuliah(db)
	upk := uk.NewUsecaseKuliah(rpd, rpm, rpmk, rp)
	hndlpk := kuliah.NewHandleKuliah(upk)
	kuliahgrp := e.Group("/kuliah")
	kuliahgrp.POST("/addkuliah", hndlpk.AddPerkuliahaan)
	kuliahgrp.GET("", hndlpk.GetParam)
}
