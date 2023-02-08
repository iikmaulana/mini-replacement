package handler

import (
	"github.com/iikmaulana/mini-replacement/base/service"
	bv "github.com/uzzeet/uzzeet-gateway/service"
)

type routerHandler struct {
	customerUsecase service.CustomerUsecase
	vehicleUsecase  service.VehicleUsecase
}

func CreateHandler(svc *bv.Service, customerUsecase service.CustomerUsecase, vehicleUsecase service.VehicleUsecase) {
	obj := routerHandler{
		customerUsecase: customerUsecase,
		vehicleUsecase:  vehicleUsecase,
	}

	svc.POST("/check_create_mt_member", obj.CreateMtMember)
	svc.POST("/check_update_mt_member", obj.UpdateMtMember)
	svc.POST("/check_create_auth_runner", obj.CreateAuthRunner)
	svc.POST("/check_update_auth_runner", obj.UpdateAuthRunner)
	svc.POST("/check_create_vehicle_runner", obj.CreateMtVehicle)
	svc.POST("/check_update_vehicle_runner", obj.UpdateMtVehicle)

}
