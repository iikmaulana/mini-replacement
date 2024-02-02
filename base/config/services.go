package config

import (
	"github.com/iikmaulana/mini-replacement/base/controller"
	"github.com/iikmaulana/mini-replacement/base/packets"
	"github.com/iikmaulana/mini-replacement/base/service/grpc"
	"github.com/iikmaulana/mini-replacement/base/service/handler"
	"github.com/iikmaulana/mini-replacement/base/service/repository/core"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
)

func (cfg Config) InitService() serror.SError {

	userRepo, serr := core.NewUserRepository(cfg.Registry)
	if serr != nil {
		return serr
	}

	organizationRepo, serr := core.NewOrganizationRepository(cfg.Registry)
	if serr != nil {
		return serr
	}

	vehicleRepo, serr := core.NewvehicleRepository(cfg.Registry)
	if serr != nil {
		return serr
	}

	vehicleHistoryRepo, serr := core.NewVehicleHistoryRepository(cfg.Registry)
	if serr != nil {
		return serr
	}

	customerUsecase := controller.NewCustomerUsecase(cfg.DB, userRepo, organizationRepo, cfg.NSQ)
	vehicleUsecase := controller.NewVehicleUsecase(cfg.DB, userRepo, organizationRepo, vehicleRepo, vehicleHistoryRepo, cfg.NSQ)

	packets.RegisterCustomerReplacementServer(cfg.Server.Instance(), grpc.NewCustomerHandler(customerUsecase))
	packets.RegisterVehicleReplacementServer(cfg.Server.Instance(), grpc.NewVehicleHandler(vehicleUsecase))

	handler.CreateHandler(cfg.Gateway, customerUsecase, vehicleUsecase)

	return nil
}
