package handlecontract

import (
	"github.com/labstack/echo/v4"
)

type HandleDosen interface {
	AddDosen(e echo.Context) error
	GetAllDosen(e echo.Context) error
	// DeleteMahasiswa (e ech)
}
type HandleMahasiswa interface {
	AddMahasiswa(e echo.Context) error
	GetAllMahasiswa(e echo.Context) error
	// DeleteMahasiswa (e ech)
}
type HandleMatakuliah interface {
	AddMatakuliah(e echo.Context) error
}
