package controller

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	lib "github.com/iikmaulana/mini-replacement/base/libs"
	"github.com/iikmaulana/mini-replacement/base/models"
	"github.com/iikmaulana/mini-replacement/base/service"
	"github.com/iikmaulana/mini-replacement/base/service/repository/query"
	"github.com/jmoiron/sqlx"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
	"github.com/uzzeet/uzzeet-gateway/libs/utils/uttime"
)

type vehicleUsecase struct {
	DB                 *sqlx.DB
	userRepo           service.UserRepo
	organizationRepo   service.OrganizationRepo
	vehicleRepo        service.VehicleRepo
	vehicleHistoryRepo service.VehicleHistoryRepo
}

func NewVehicleUsecase(db *sqlx.DB, userRepo service.UserRepo, organizationRepo service.OrganizationRepo, vehicleRepo service.VehicleRepo, vehicleHistoryRepo service.VehicleHistoryRepo) service.VehicleUsecase {
	return vehicleUsecase{DB: db, userRepo: userRepo, organizationRepo: organizationRepo, vehicleRepo: vehicleRepo, vehicleHistoryRepo: vehicleHistoryRepo}
}

func (v vehicleUsecase) CreateMtVehicle(form []byte) (result string, serr serror.SError) {
	val := map[string]interface{}{}
	_ = json.Unmarshal(form, &val)

	tmpVehicle, _ := v.vehicleHistoryRepo.VehicleHistory_GetByChassisRepo(models.GetVehHistoryByChassisReq{
		ChassisNumber: val["chassis_number"].(string),
	})

	var memberId string
	if err := v.DB.QueryRow(fmt.Sprintf(query.GetMemberIdByOrganizationId, *tmpVehicle.OwnerID)).Scan(&memberId); err != nil {
		fmt.Println(err.Error())
	}

	if tmpVehicle.ChassisNumber != "" && memberId != "" {
		tmpDefaultValue := ""
		if tmpVehicle.Imei == nil {
			tmpVehicle.Imei = &tmpDefaultValue
		}
		if tmpVehicle.VehicleNumber == nil {
			tmpVehicle.VehicleNumber = &tmpDefaultValue
		}
		if tmpVehicle.ActivationDate == nil {
			tmpVehicle.ActivationDate = &tmpDefaultValue
		}
		tmpQuery := fmt.Sprintf(query.CreateMtVehcile,
			tmpVehicle.ChassisNumber,
			*tmpVehicle.Imei,
			tmpVehicle.VehicleName,
			*tmpVehicle.VehicleNumber,
			memberId,
			*tmpVehicle.ActivationDate,
		)

		tmpForm, _ := json.Marshal(val)
		tmpOauthRunner := lib.PayloadNsq{
			RequestID:    uuid.New().String(),
			Time:         uttime.Now().String(),
			Service:      "vehicle",
			DatabaseName: "dev_runner_app",
			Payload:      string(tmpForm),
			Query:        tmpQuery,
		}
		tmpByteOauthRunner, _ := json.Marshal(tmpOauthRunner)
		_ = lib.SendNSQUsecase(tmpByteOauthRunner)

	}
	return "", nil
}

func (v vehicleUsecase) UpdateMtVehicle(form []byte) (result string, serr serror.SError) {
	val := map[string]interface{}{}
	_ = json.Unmarshal(form, &val)

	tmpVehicle, _ := v.vehicleHistoryRepo.VehicleHistory_GetByChassisRepo(models.GetVehHistoryByChassisReq{
		ChassisNumber: val["chassis_number"].(string),
	})

	var memberId string
	if err := v.DB.QueryRow(fmt.Sprintf(query.GetMemberIdByOrganizationId, *tmpVehicle.OwnerID)).Scan(&memberId); err != nil {
		fmt.Println(err.Error())
	}

	if tmpVehicle.ChassisNumber != "" && memberId != "" {
		tmpDefaultValue := ""
		if tmpVehicle.Imei == nil {
			tmpVehicle.Imei = &tmpDefaultValue
		}
		if tmpVehicle.VehicleNumber == nil {
			tmpVehicle.VehicleNumber = &tmpDefaultValue
		}
		if tmpVehicle.ActivationDate == nil {
			tmpVehicle.ActivationDate = &tmpDefaultValue
		}
		tmpQuery := fmt.Sprintf(query.UpdateMtVehcile,
			tmpVehicle.ChassisNumber,
			*tmpVehicle.Imei,
			tmpVehicle.VehicleName,
			*tmpVehicle.VehicleNumber,
			memberId,
			*tmpVehicle.ActivationDate,
			tmpVehicle.ChassisNumber,
			memberId,
		)

		tmpForm, _ := json.Marshal(val)
		tmpOauthRunner := lib.PayloadNsq{
			RequestID:    uuid.New().String(),
			Time:         uttime.Now().String(),
			Service:      "UM",
			DatabaseName: "dev_runner_app",
			Payload:      string(tmpForm),
			Query:        tmpQuery,
		}
		tmpByteOauthRunner, _ := json.Marshal(tmpOauthRunner)
		_ = lib.SendNSQUsecase(tmpByteOauthRunner)
	}

	return "", nil
}
