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

type organizationRepository struct {
	client packets.UmOrganizationClient
}

func NewOrganizationRepository(reg *controller.Registry) (res service.OrganizationRepo, serr serror.SError) {
	obj := organizationRepository{}
	i := 3

	for {
		if i <= 0 {
			serr = serror.Newc("Cannot connect to core Organization",
				"[repository][core][organization] while dialing core Organization")
			break
		}

		conn, err := reg.GetConnection("base_user")
		if err != nil {
			logger.Warn("[repository][core][organization] while dial core Organization")
			time.Sleep(1 * time.Second)

			i--
			continue
		}
		obj.client = packets.NewUmOrganizationClient(conn)
		break
	}

	return obj, serr
}

func (a organizationRepository) GetOrganizationByIdRepo(id string) (result models.RpcListOrganizationCustomerResult, serr serror.SError) {
	output, err := a.client.GetOrganizationByIdUsecase(context.Background(), &packets.OrganizationRequestByID{
		OrganizationID: id,
	})

	if err != nil {
		serrFmt := fmt.Sprintf("[service][repository][core][Organization] while grpc GetOrganizationByIdUsecase: %+v", err)
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
