package handler

import (
	"github.com/uzzeet/uzzeet-gateway/models"
	"github.com/uzzeet/uzzeet-gateway/service"
	"net/http"
)

func (ox routerHandler) CreateMtMember(ctx *service.Context) service.Result {
	body := ctx.RawBody()
	_, serr := ox.customerUsecase.CreateMtMemberUsecase(body)
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

func (ox routerHandler) UpdateMtMember(ctx *service.Context) service.Result {
	body := ctx.RawBody()
	_, serr := ox.customerUsecase.UpdateMtMemberUsecase(body)
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

func (ox routerHandler) CreateAuthRunner(ctx *service.Context) service.Result {
	body := ctx.RawBody()
	_, serr := ox.customerUsecase.CreateAuthRunnerUsecase(body)
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

func (ox routerHandler) UpdateAuthRunner(ctx *service.Context) service.Result {
	body := ctx.RawBody()
	_, serr := ox.customerUsecase.UpdateAuthRunnerUsecase(body)
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
