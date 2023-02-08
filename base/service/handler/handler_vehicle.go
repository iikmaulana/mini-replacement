package handler

import (
	"github.com/uzzeet/uzzeet-gateway/models"
	"github.com/uzzeet/uzzeet-gateway/service"
	"net/http"
)

func (ox routerHandler) CreateMtVehicle(ctx *service.Context) service.Result {
	body := ctx.RawBody()
	_, serr := ox.vehicleUsecase.CreateMtVehicle(body)
	if serr != nil {
		return ctx.JSONResponse(http.StatusBadRequest, models.ResponseBody{
			Result: map[string]string{
				"en": "Failed resend email",
				"id": "Gagal mengirim email",
			},
		})
	}
	return ctx.JSONResponse(http.StatusOK, models.ResponseBody{
		Result: map[string]string{
			"en": "Success",
			"id": "Berhasil",
		},
	})
}

func (ox routerHandler) UpdateMtVehicle(ctx *service.Context) service.Result {
	body := ctx.RawBody()
	_, serr := ox.vehicleUsecase.UpdateMtVehicle(body)
	if serr != nil {
		return ctx.JSONResponse(http.StatusBadRequest, models.ResponseBody{
			Result: map[string]string{
				"en": "Failed resend email",
				"id": "Gagal mengirim email",
			},
		})
	}
	return ctx.JSONResponse(http.StatusOK, models.ResponseBody{
		Result: map[string]string{
			"en": "Success",
			"id": "Berhasil",
		},
	})
}
