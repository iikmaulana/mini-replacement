package grpc

import (
	"context"
	"encoding/json"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/iikmaulana/mini-replacement/base/packets"
	"github.com/iikmaulana/mini-replacement/base/service"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
)

type CustomerHandler struct {
	customerUsecase service.CustomerUsecase
}

func NewCustomerHandler(customerUsecase service.CustomerUsecase) CustomerHandler {
	return CustomerHandler{
		customerUsecase: customerUsecase}
}

func (handler CustomerHandler) UpdateMtMember(ctx context.Context, form *packets.CustomerReplacementRequest) (output *packets.CustomerReplacementOutput, err error) {
	output = &packets.CustomerReplacementOutput{
		Status: 0,
	}

	tmpResult, serr := handler.customerUsecase.UpdateMtMemberUsecase(form.Data.Value)
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

func (handler CustomerHandler) CreateMtMember(ctx context.Context, form *packets.CustomerReplacementRequest) (output *packets.CustomerReplacementOutput, err error) {
	output = &packets.CustomerReplacementOutput{
		Status: 0,
	}

	tmpResult, serr := handler.customerUsecase.CreateAuthRunnerUsecase(form.Data.Value)
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

func (handler CustomerHandler) UpdateAuthRunner(ctx context.Context, form *packets.CustomerReplacementRequest) (output *packets.CustomerReplacementOutput, err error) {
	output = &packets.CustomerReplacementOutput{
		Status: 0,
	}

	tmpResult, serr := handler.customerUsecase.UpdateAuthRunnerUsecase(form.Data.Value)
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

func (handler CustomerHandler) CreateAuthRunner(ctx context.Context, form *packets.CustomerReplacementRequest) (output *packets.CustomerReplacementOutput, err error) {
	output = &packets.CustomerReplacementOutput{
		Status: 0,
	}

	tmpResult, serr := handler.customerUsecase.CreateAuthRunnerUsecase(form.Data.Value)
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

func (handler CustomerHandler) UserActivation(ctx context.Context, form *packets.CustomerReplacementRequest) (output *packets.CustomerReplacementOutput, err error) {
	output = &packets.CustomerReplacementOutput{
		Status: 0,
	}

	tmpResult, serr := handler.customerUsecase.CreateAuthRunnerUsecase(form.Data.Value)
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

func (handler CustomerHandler) ChangePassword(ctx context.Context, form *packets.CustomerReplacementRequest) (output *packets.CustomerReplacementOutput, err error) {
	output = &packets.CustomerReplacementOutput{
		Status: 0,
	}

	tmpResult, serr := handler.customerUsecase.ChangePasswordUsecase(form.Data.Value)
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
