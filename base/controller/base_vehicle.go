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
	"github.com/uzzeet/uzzeet-gateway/libs/helper"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
	"github.com/uzzeet/uzzeet-gateway/libs/utils/uttime"
	"strings"
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

	var tmpMtVehicle bool
	if err := v.DB.QueryRow(fmt.Sprintf(query.CheckMtVehcile, val["chassis_number"].(string))).Scan(&tmpMtVehicle); err != nil {
		fmt.Println(err.Error())
	}

	if tmpMtVehicle {
		_, _ = v.UpdateMtVehicle(form)
	} else {
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
			tmpVehicleId := lib.RandomCharacter(14)
			tmpQuery := fmt.Sprintf(query.CreateMtVehcile,
				helper.StringToInt(tmpVehicleId, 0),
				tmpVehicle.ChassisNumber,
				*tmpVehicle.Imei,
				tmpVehicle.VehicleName,
				*tmpVehicle.VehicleNumber,
				helper.StringToInt(memberId, 0),
				*tmpVehicle.ActivationDate,
			)

			if *tmpVehicle.ActivationDate == "" {
				tmpQuery = strings.ReplaceAll(tmpQuery, "activation_date = ''", "activation_date = NULL")
			}

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
	if err := v.DB.QueryRow(fmt.Sprintf(query.GetMemberIdByOrganizationId, val["owner_id"].(string))).Scan(&memberId); err != nil {
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
			helper.StringToInt(memberId, 0),
			*tmpVehicle.ActivationDate,
			tmpVehicle.ChassisNumber,
			helper.StringToInt(memberId, 0),
		)

		if *tmpVehicle.ActivationDate == "" {
			tmpQuery = strings.ReplaceAll(tmpQuery, "activation_date = ''", "activation_date = NULL")
		}

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

func (v vehicleUsecase) CreateVehicleGroup(form []byte) (result string, serr serror.SError) {
	val := map[string]interface{}{}
	_ = json.Unmarshal(form, &val)

	var memberId string
	if err := v.DB.QueryRow(fmt.Sprintf(query.GetMemberIdByOrganizationId, val["owner_id"].(string))).Scan(&memberId); err != nil {
		fmt.Println(err.Error())
	}

	tmpGroupId := lib.RandomCharacter(14)
	tmpQuery := fmt.Sprintf(query.CreateVehiclGroup,
		helper.StringToInt(tmpGroupId, 0),
		val["vehicle_group_name"].(string),
		val["description"].(string),
		helper.StringToInt(memberId, 0),
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

	tmpChassisNumber := val["chassis_number"].(string)
	tmpChassisNumberArray := strings.Split(tmpChassisNumber, ",")
	for _, v := range tmpChassisNumberArray {
		tmpQuery = fmt.Sprintf(query.CreateMSVehicleGroup,
			helper.StringToInt(tmpGroupId, 0),
			helper.StringToInt(v, 0),
		)

		tmpForm, _ = json.Marshal(val)
		tmpOauthRunner = lib.PayloadNsq{
			RequestID:    uuid.New().String(),
			Time:         uttime.Now().String(),
			Service:      "vehicle",
			DatabaseName: "dev_runner_app",
			Payload:      string(tmpForm),
			Query:        tmpQuery,
		}

		tmpByteOauthRunner, _ = json.Marshal(tmpOauthRunner)
		_ = lib.SendNSQUsecase(tmpByteOauthRunner)
	}

	return "", nil
}

func (v vehicleUsecase) UpdateVehicleGroup(form []byte) (result string, serr serror.SError) {
	return "", nil
}
