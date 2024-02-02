package core

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/iikmaulana/mini-replacement/base/models"
	"github.com/iikmaulana/mini-replacement/base/packets"
	"github.com/iikmaulana/mini-replacement/base/service"
	"github.com/uzzeet/uzzeet-gateway/controller"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/logger"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
	"time"
)

type vehicleHistoryRepository struct {
	client packets.VehicleHistoryClient
}

func NewVehicleHistoryRepository(reg *controller.Registry) (res service.VehicleHistoryRepo, serr serror.SError) {
	obj := vehicleHistoryRepository{}
	i := 3

	for {
		if i <= 0 {
			serr = serror.Newc("Cannot connect to core Base Vehicle",
				"[repository][core][vehicle] while dialing core Base Vehicle")
			break
		}

		conn, err := reg.GetConnection("runner-base-vehicle")
		if err != nil {
			logger.Warn("[repository][core][vehicle] while dial core Base Vehicle")
			time.Sleep(1 * time.Second)

			i--
			continue
		}
		obj.client = packets.NewVehicleHistoryClient(conn)
		break
	}

	return obj, serr
}

func (v vehicleHistoryRepository) VehicleHistory_GetByChassisRepo(form models.GetVehHistoryByChassisReq) (result models.VehHistoryByChassisResult, serr serror.SError) {
	byte, err := json.Marshal(form)
	if err != nil {
		return result, serror.NewFromError(err)
	}

	data := any.Any{
		Value: byte,
	}

	output, err := v.client.VehicleHistory_GetByChassis(context.Background(), &packets.VehHistoryRequest{
		Data: &data,
	})

	if err != nil {
		serrFmt := fmt.Sprintf("[service][repository][core][vehicles][ChassisNumber:%s] while grpc GetVehicleHistoryByChassisRepo: %v", form.ChassisNumber, err)
		logger.Err(serrFmt)
		return result, serror.NewFromErrorc(err, serrFmt)
	}
	if output.GetStatus() == 1 {
		err := json.Unmarshal(output.GetData().Value, &result)
		if err != nil {
			return result, serror.NewFromErrorc(err, "Gagal mapping")
		}
	}

	return result, nil
}

func (v vehicleHistoryRepository) VehicleHistory_GetRepo(form models.GetVehHistoryReq) (result models.VehHistoryListMetaResult, serr serror.SError) {
	byte, err := json.Marshal(form)
	if err != nil {
		return result, serror.NewFromError(err)
	}

	data := any.Any{
		Value: byte,
	}

	output, err := v.client.VehicleHistory_Get(context.Background(), &packets.VehHistoryRequest{
		Data: &data,
	})

	if err != nil {
		serrFmt := fmt.Sprintf("[service][repository][core][vehicles][OwnerID:%s] while grpc GetVehicleHistoryRepo: %v", form.OwnerID, err)
		logger.Err(serrFmt)
		return result, serror.NewFromErrorc(err, serrFmt)
	}
	if output.GetStatus() == 1 {
		err := json.Unmarshal(output.GetData().Value, &result)
		if err != nil {
			return result, serror.NewFromErrorc(err, "Gagal mapping")
		}
	}

	return result, nil
}
