package dosen

import (
	"backend/domain/contract/handlecontract"
	"backend/domain/contract/usecasecontract"
	"backend/domain/request"
	"backend/domain/response"
	"backend/helper"

	echo "github.com/labstack/echo/v4"
)

type HandlerDosen struct {
	ud usecasecontract.UseCaseDosen
}

func NewHandleDosen(ud usecasecontract.UseCaseDosen) handlecontract.HandleDosen {
	return &HandlerDosen{
		ud: ud,
	}
}

// AddDosen implements handlecontract.HandleDosen.
func (hd *HandlerDosen) AddDosen(e echo.Context) error {
	reqdosen := request.DosenRequest{}

	binderr := e.Bind(&reqdosen)

	if binderr != nil {
		return e.JSON(400, helper.GetResponse("binding error", 400, true))
	}
	uc, err := hd.ud.AddDosen(reqdosen)

	if err != nil {
		return e.JSON(500, helper.GetResponse("internal server error", 500, true))
	}

	responsdosen := response.RequestoFormResponse(uc)
	return e.JSON(201, helper.GetResponse(responsdosen, 201, false))

}

// GetAllDosen implements handlecontract.HandleDosen.
func (hd *HandlerDosen) GetAllDosen(e echo.Context) error {
	datauc, err := hd.ud.GetAllDosen()

	if err != nil {
		return e.JSON(500, helper.GetResponse("internal server error", 500, true))
	}
	respon := response.ListReqToRespon(datauc)
	return e.JSON(201, helper.GetResponse(respon, 201, false))
}
