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

type userRepository struct {
	client packets.UmUserClient
}

func NewUserRepository(reg *controller.Registry) (res service.UserRepo, serr serror.SError) {
	obj := userRepository{}
	i := 3

	for {
		if i <= 0 {
			serr = serror.Newc("Cannot connect to core User",
				"[repository][core][user] while dialing core User")
			break
		}

		conn, err := reg.GetConnection("base_user")
		if err != nil {
			logger.Warn("[repository][core][user] while dial core User")
			time.Sleep(1 * time.Second)

			i--
			continue
		}
		obj.client = packets.NewUmUserClient(conn)
		break
	}

	return obj, serr
}

func (r userRepository) GetUserByUsernameRepo(username string) (result models.UserResult, serr serror.SError) {
	output, err := r.client.GetUserByUsernameUsecase(context.Background(), &packets.UserRequestByUsername{
		Username: username,
	})

	if err != nil {
		serrFmt := fmt.Sprintf("[service][repository][core][user] while grpc GetUserByUsernameRepo: %+v", err)
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

func (r userRepository) GetUserByIDRepo(id string) (result models.UserResult, serr serror.SError) {
	output, err := r.client.GetUserByIdUsecase(context.Background(), &packets.UserRequestByID{
		UserID: id,
	})

	if err != nil {
		serrFmt := fmt.Sprintf("[service][repository][core][user] while grpc GetUserByUsernameRepo: %+v", err)
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
