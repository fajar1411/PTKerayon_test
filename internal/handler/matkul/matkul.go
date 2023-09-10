package matkul

import (
	"backend/domain/contract/handlecontract"
	"backend/domain/contract/usecasecontract"
	"backend/domain/request"
	"backend/domain/response"
	"backend/helper"

	echo "github.com/labstack/echo/v4"
)

type HandlerMatkul struct {
	ud usecasecontract.UseCaseMatkul
}

func NewHandleMatkul(ud usecasecontract.UseCaseMatkul) handlecontract.HandleMatakuliah {
	return &HandlerMatkul{
		ud: ud,
	}
}

// AddMatakuliah implements handlecontract.HandleMatakuliah.
func (hm *HandlerMatkul) AddMatakuliah(e echo.Context) error {
	reqmatkul := request.MatkulRequest{}

	binderr := e.Bind(&reqmatkul)

	if binderr != nil {
		return e.JSON(400, helper.GetResponse("binding error", 400, true))
	}
	uc, err := hm.ud.AddMatkul(reqmatkul)

	if err != nil {
		return e.JSON(500, helper.GetResponse("internal server error", 500, true))
	}

	responsmatkul := response.MatkulRequestoFormResponse(uc)
	return e.JSON(201, helper.GetResponse(responsmatkul, 201, false))
}
