package service

import (
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
)

type CustomerUsecase interface {
	CreateMtMemberUsecase(form []byte) (result string, serr serror.SError)
	UpdateMtMemberUsecase(form []byte) (result string, serr serror.SError)
	CreateAuthRunnerUsecase(form []byte) (result string, serr serror.SError)
	UpdateAuthRunnerUsecase(form []byte) (result string, serr serror.SError)
	UserActivationUsecase(form []byte) (result string, serr serror.SError)
	ChangePasswordUsecase(form []byte) (result string, serr serror.SError)
	PrivacyPolicyUsecase(form []byte, tmpUserId string) (result string, serr serror.SError)
	ApprovePrivacyPolicyUsecase(form []byte) (result string, serr serror.SError)
	DeleteAuthRunnerUsecase(form []byte) (result string, serr serror.SError)
}

type VehicleUsecase interface {
	CreateMtVehicle(form []byte) (result string, serr serror.SError)
	UpdateMtVehicle(form []byte) (result string, serr serror.SError)
	CreateVehicleGroup(form []byte) (result string, serr serror.SError)
	UpdateVehicleGroup(form []byte) (result string, serr serror.SError)
}
