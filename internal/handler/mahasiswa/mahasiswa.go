package mahasiswa

import (
	"backend/domain/contract/handlecontract"
	"backend/domain/contract/usecasecontract"
	"backend/domain/request"
	"backend/domain/response"
	"backend/helper"
	"log"

	"github.com/labstack/echo/v4"
)

type HandlerMahasiswa struct {
	um usecasecontract.UseCaseMahasiswa
}

func NewHandleMahasiswa(um usecasecontract.UseCaseMahasiswa) handlecontract.HandleMahasiswa {
	return &HandlerMahasiswa{
		um: um,
	}
}

func (hm *HandlerMahasiswa) AddMahasiswa(e echo.Context) error {
	reqMahasiswa := request.MahasiswaRequest{}

	binderr := e.Bind(&reqMahasiswa)

	if binderr != nil {
		return e.JSON(400, helper.GetResponse("binding error", 400, true))
	}
	uc, err := hm.um.AddMahasiswa(reqMahasiswa)

	if err != nil {
		return e.JSON(500, helper.GetResponse("internal server error", 500, true))
	}

	responsmahasiswa := response.ToFormResponse(uc)
	log.Print(responsmahasiswa)
	return e.JSON(201, helper.GetResponse(responsmahasiswa, 201, false))

}

// GetAllMahasiswa implements handlecontract.HandleMahasiswa.
func (hm *HandlerMahasiswa) GetAllMahasiswa(e echo.Context) error {
	panic("unimplemented")
}
