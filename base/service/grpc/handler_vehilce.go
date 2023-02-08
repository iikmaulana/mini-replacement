package grpc

import (
	"context"
	"encoding/json"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/iikmaulana/mini-replacement/base/packets"
	"github.com/iikmaulana/mini-replacement/base/service"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
)

type VehicleHandler struct {
	vehicleUsecase service.VehicleUsecase
}

func NewVehicleHandler(vehicleUsecase service.VehicleUsecase) VehicleHandler {
	return VehicleHandler{
		vehicleUsecase: vehicleUsecase}
}

func (handler VehicleHandler) CreateMtVehicle(ctx context.Context, form *packets.VehicleReplacementRequest) (output *packets.VehicleReplacementOutput, err error) {
	output = &packets.VehicleReplacementOutput{
		Status: 0,
	}

	tmpResult, serr := handler.vehicleUsecase.CreateMtVehicle(form.Data.Value)
	if serr != nil {
		return output, serr
	}

	b, err := json.Marshal(&tmpResult)
	if err != nil {
		return output, serror.NewFromError(err)
	}

	output.Status = 1
	output.Data = &any.Any{
		Value: b,
	}

	return output, nil
}

func (handler VehicleHandler) UpdateMtVehicle(ctx context.Context, form *packets.VehicleReplacementRequest) (output *packets.VehicleReplacementOutput, err error) {
	output = &packets.VehicleReplacementOutput{
		Status: 0,
	}

	tmpResult, serr := handler.vehicleUsecase.UpdateMtVehicle(form.Data.Value)
	if serr != nil {
		return output, serr
	}

	b, err := json.Marshal(&tmpResult)
	if err != nil {
		return output, serror.NewFromError(err)
	}

	output.Status = 1
	output.Data = &any.Any{
		Value: b,
	}

	return output, nil
}
