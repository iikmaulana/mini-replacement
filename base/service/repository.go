package service

import (
	"github.com/iikmaulana/mini-replacement/base/models"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
)

type UserRepo interface {
	GetUserByUsernameRepo(username string) (result models.UserResult, serr serror.SError)
	GetUserByIDRepo(id string) (result models.UserResult, serr serror.SError)
	GetUserByIDOrganizationRepo(id string) (result models.UserResult, serr serror.SError)
}

type OrganizationRepo interface {
	GetOrganizationByIdRepo(id string) (result models.RpcListOrganizationCustomerResult, serr serror.SError)
}

type VehicleRepo interface {
	Vehicle_GetByChassisRepo(chassisNumber string) (result models.VehResult, serr serror.SError)
}

type VehicleHistoryRepo interface {
	VehicleHistory_GetByChassisRepo(form models.GetVehHistoryByChassisReq) (result models.VehHistoryByChassisResult, serr serror.SError)
}
