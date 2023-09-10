package kuliah

import (
	"backend/domain/contract/handlecontract"
	"backend/domain/contract/usecasecontract"
	"backend/domain/request"
	"backend/domain/response"
	"backend/helper"

	echo "github.com/labstack/echo/v4"
)

type HandlerKuliah struct {
	ud usecasecontract.UseCasekuliah
}

// AddPerkuliahaan implements handlecontract.HandlePerkuliahaan.

func NewHandleKuliah(ud usecasecontract.UseCasekuliah) handlecontract.HandlePerkuliahaan {
	return &HandlerKuliah{
		ud: ud,
	}
}
func (hk *HandlerKuliah) AddPerkuliahaan(e echo.Context) error {
	reqKuliah := request.KuliahRequest{}
	binderr := e.Bind(&reqKuliah)

	if binderr != nil {
		return e.JSON(400, helper.GetResponse("binding error", 400, true))
	}
	uc, err := hk.ud.Addkuliah(reqKuliah)

	if err != nil {
		return e.JSON(500, helper.GetResponse("internal server error", 500, true))
	}

	responsKuliah := response.KuliahRequestoFormResponse(uc)
	return e.JSON(201, helper.GetResponse(responsKuliah, 201, false))
}

// GetParam implements handlecontract.HandlePerkuliahaan.
func (hk *HandlerKuliah) GetParam(e echo.Context) error {
	idmhs := e.QueryParam("nim")

	res, err := hk.ud.ParamReq(idmhs)

	if err != nil {
		return e.JSON(500, helper.GetResponse("internal server error", 500, true))
	}

	respon := response.ParamRequestoFormResponse(res)
	return e.JSON(201, helper.GetResponse(respon, 201, false))
}
