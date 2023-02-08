package core

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/iikmaulana/mini-replacement/base/models"
	"github.com/iikmaulana/mini-replacement/base/packets"
	"github.com/iikmaulana/mini-replacement/base/service"
	"github.com/uzzeet/uzzeet-gateway/controller"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/logger"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
	"time"
)

type vehicleRepository struct {
	client packets.VehicleClient
}

func NewvehicleRepository(reg *controller.Registry) (res service.VehicleRepo, serr serror.SError) {
	obj := vehicleRepository{}
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
		obj.client = packets.NewVehicleClient(conn)
		break
	}

	return obj, serr
}

func (v vehicleRepository) Vehicle_GetByChassisRepo(chassisNumber string) (result models.VehResult, serr serror.SError) {
	output, err := v.client.Vehicle_GetByChassis(context.Background(), &packets.VehRequestByChassis{
		ChassisNumber: chassisNumber,
	})

	if err != nil {
		serrFmt := fmt.Sprintf("[service][repository][core][vehicle] while grpc Vehicle_GetByChassisRepo: %+v", err)
		logger.Err(serrFmt)
		return result, serror.NewFromErrorc(err, serrFmt)
	}

	if output.GetStatus() == 1 {
		err := json.Unmarshal(output.GetData().Value, &result)
		if err != nil {
			return result, serror.NewFromError(err)
		}
	}

	return result, nil
}
