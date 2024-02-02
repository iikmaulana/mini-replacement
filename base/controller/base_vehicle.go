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
	"github.com/nsqio/go-nsq"
	"github.com/uzzeet/uzzeet-gateway/libs/helper"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
	"github.com/uzzeet/uzzeet-gateway/libs/utils/uttime"
	"os"
	"strings"
)

type vehicleUsecase struct {
	DB                 *sqlx.DB
	userRepo           service.UserRepo
	organizationRepo   service.OrganizationRepo
	vehicleRepo        service.VehicleRepo
	vehicleHistoryRepo service.VehicleHistoryRepo
	NSQ                *nsq.Producer
}

func NewVehicleUsecase(db *sqlx.DB, userRepo service.UserRepo, organizationRepo service.OrganizationRepo, vehicleRepo service.VehicleRepo, vehicleHistoryRepo service.VehicleHistoryRepo, nsq *nsq.Producer) service.VehicleUsecase {
	return vehicleUsecase{DB: db, userRepo: userRepo, organizationRepo: organizationRepo, vehicleRepo: vehicleRepo, vehicleHistoryRepo: vehicleHistoryRepo, NSQ: nsq}
}

func (v vehicleUsecase) CreateMtVehicle(form []byte) (result string, serr serror.SError) {
	val := map[string]interface{}{}
	err := json.Unmarshal(form, &val)
	if err != nil {
		val2 := []models.VehicleReq{}
		_ = json.Unmarshal(form, &val2)
		for _, vr := range val2 {
			var tmpMtVehicle bool
			tmpQuery := fmt.Sprintf(query.CheckMtVehcile, vr.ChassisNumber)
			if err := v.DB.QueryRow(tmpQuery).Scan(&tmpMtVehicle); err != nil {
				fmt.Println(err.Error())
			}

			if tmpMtVehicle {
				_, _ = v.UpdateMtVehicle(form)
				fmt.Println("============= update vehicle =============")
			} else {
				fmt.Println("============= create vehicle =============")
				tmpQueryVehicle := fmt.Sprintf(query.GetVehicle, vr.ChassisNumber)
				rows, err := v.DB.Queryx(tmpQueryVehicle)
				if err != nil {
					fmt.Println(err.Error())
					return result, serror.New(err.Error())
				}

				var tmpVehicleHistoryArray []models.VehicleV3
				defer rows.Close()
				for rows.Next() {
					tmpVehicleHistory := models.VehicleV3{}
					if err := rows.StructScan(&tmpVehicleHistory); err != nil {
						fmt.Println(err.Error())
					}
					tmpVehicleHistoryArray = append(tmpVehicleHistoryArray, tmpVehicleHistory)
				}

				if len(tmpVehicleHistoryArray) == 0 {
					fmt.Println("============= create vehicle gagal =============")
				}

				for _, vv := range tmpVehicleHistoryArray {
					if vv.ChassisNumber != "" {
						var tmpMtVehicle bool
						if err := v.DB.QueryRow(fmt.Sprintf(query.CheckMtVehcile, vr.ChassisNumber)).Scan(&tmpMtVehicle); err != nil {
							fmt.Println(err.Error())
						}

						if !tmpMtVehicle {
							tmpVehicleId := lib.RandomCharacter(14)
							tmpDefaultValue := "NULL"

							if vv.ChassisNumber == "" {
								vv.ChassisNumber = tmpDefaultValue
							} else {
								tmpStr := fmt.Sprintf("'%s'", vv.ChassisNumber)
								vv.ChassisNumber = tmpStr
							}

							if vv.Imei == nil {
								vv.Imei = &tmpDefaultValue
							} else {
								tmpStr := fmt.Sprintf("'%s'", *vv.Imei)
								vv.Imei = &tmpStr
							}
							if vv.VehicleName == nil {
								vv.VehicleName = &tmpDefaultValue
							} else {
								tmpStr := fmt.Sprintf("'%s'", *vv.VehicleName)
								vv.VehicleName = &tmpStr
							}
							if vv.VehicleNumber == nil {
								vv.VehicleNumber = &tmpDefaultValue
							} else {
								tmpStr := fmt.Sprintf("'%s'", *vv.VehicleNumber)
								vv.VehicleNumber = &tmpStr
							}
							if vv.MemberID == nil {
								vv.MemberID = &tmpDefaultValue
							} else {
								tmpStr := fmt.Sprintf("'%s'", *vv.MemberID)
								vv.MemberID = &tmpStr
							}
							if vv.IsActive == nil {
								vv.IsActive = &tmpDefaultValue
							} else {
								tmpStr := fmt.Sprintf("'%s'", *vv.IsActive)
								vv.IsActive = &tmpStr
							}
							if vv.ActivationDate == nil {
								vv.ActivationDate = &tmpDefaultValue
							} else {
								tmpStr := fmt.Sprintf("'%s'", *vv.ActivationDate)
								vv.ActivationDate = &tmpStr
							}
							if vv.DeviceStatus == nil {
								vv.DeviceStatus = &tmpDefaultValue
							} else {
								tmpStr := fmt.Sprintf("'%s'", *vv.DeviceStatus)
								vv.DeviceStatus = &tmpStr
							}
							if vv.PostDealerID == nil {
								vv.PostDealerID = &tmpDefaultValue
							} else {
								tmpStr := fmt.Sprintf("'%s'", *vv.PostDealerID)
								vv.PostDealerID = &tmpStr
							}
							if vv.ActvDealerID == nil {
								vv.ActvDealerID = &tmpDefaultValue
							} else {
								tmpStr := fmt.Sprintf("'%s'", *vv.ActvDealerID)
								vv.ActvDealerID = &tmpStr
							}
							if vv.GsmNumber == nil {
								vv.GsmNumber = &tmpDefaultValue
							} else {
								tmpStr := fmt.Sprintf("'%s'", *vv.GsmNumber)
								vv.GsmNumber = &tmpStr
							}
							if vv.EngineNumber == nil {
								vv.EngineNumber = &tmpDefaultValue
							} else {
								tmpStr := fmt.Sprintf("'%s'", *vv.EngineNumber)
								vv.EngineNumber = &tmpStr
							}
							if vv.TypeId == nil {
								vv.TypeId = &tmpDefaultValue
							} else {
								tmpStr := fmt.Sprintf("'%s'", *vv.TypeId)
								vv.TypeId = &tmpStr
							}

							tmpQuery := fmt.Sprintf(query.CreateMtVehicleV2,
								helper.StringToInt(tmpVehicleId, 0),
								vv.ChassisNumber,
								*vv.Imei,
								*vv.VehicleName,
								*vv.VehicleNumber,
								*vv.MemberID,
								*vv.IsActive,
								*vv.ActivationDate,
								*vv.DeviceStatus,
								*vv.PostDealerID,
								*vv.ActvDealerID,
								*vv.GsmNumber,
								*vv.EngineNumber,
								*vv.TypeId)

							tmpForm, _ := json.Marshal(val)
							tmpOauthRunner := lib.PayloadNsq{
								RequestID:    uuid.New().String(),
								Time:         uttime.Now().String(),
								Service:      "vehicle",
								DatabaseName: os.Getenv("DBV2_RUNNER"),
								Payload:      string(tmpForm),
								Query:        tmpQuery,
							}
							tmpByteOauthRunner, _ := json.Marshal(tmpOauthRunner)
							_ = lib.SendNSQUsecase(v.NSQ, tmpByteOauthRunner)
						}
					}
				}
			}
		}
	} else {
		var tmpMtVehicle bool
		if err := v.DB.QueryRow(fmt.Sprintf(query.CheckMtVehcile, val["chassis_number"].(string))).Scan(&tmpMtVehicle); err != nil {
			fmt.Println(err.Error())
		}

		if tmpMtVehicle {
			_, _ = v.UpdateMtVehicle(form)
			fmt.Println("============= update vehicle =============")
		} else {
			fmt.Println("============= create vehicle =============")
			tmpQueryVehicle := fmt.Sprintf(query.GetVehicle, val["chassis_number"].(string))
			rows, err := v.DB.Queryx(tmpQueryVehicle)
			if err != nil {
				return result, serror.New(err.Error())
			}

			var tmpVehicleHistoryArray []models.VehicleV3
			defer rows.Close()
			for rows.Next() {
				tmpVehicleHistory := models.VehicleV3{}
				if err := rows.StructScan(&tmpVehicleHistory); err != nil {
					fmt.Println(err.Error())
				}
				tmpVehicleHistoryArray = append(tmpVehicleHistoryArray, tmpVehicleHistory)
			}

			if len(tmpVehicleHistoryArray) == 0 {
				fmt.Println("============= create vehicle gagal =============")
			}

			for _, vv := range tmpVehicleHistoryArray {
				if vv.ChassisNumber != "" {

					var tmpMtVehicle bool
					if err := v.DB.QueryRow(fmt.Sprintf(query.CheckMtVehcile, val["chassis_number"].(string))).Scan(&tmpMtVehicle); err != nil {
						fmt.Println(err.Error())
					}

					if !tmpMtVehicle {
						tmpVehicleId := lib.RandomCharacter(14)
						tmpDefaultValue := "NULL"

						if vv.ChassisNumber == "" {
							vv.ChassisNumber = tmpDefaultValue
						} else {
							tmpStr := fmt.Sprintf("'%s'", vv.ChassisNumber)
							vv.ChassisNumber = tmpStr
						}

						if vv.Imei == nil {
							vv.Imei = &tmpDefaultValue
						} else {
							tmpStr := fmt.Sprintf("'%s'", *vv.Imei)
							vv.Imei = &tmpStr
						}
						if vv.VehicleName == nil {
							vv.VehicleName = &tmpDefaultValue
						} else {
							tmpStr := fmt.Sprintf("'%s'", *vv.VehicleName)
							vv.VehicleName = &tmpStr
						}
						if vv.VehicleNumber == nil {
							vv.VehicleNumber = &tmpDefaultValue
						} else {
							tmpStr := fmt.Sprintf("'%s'", *vv.VehicleNumber)
							vv.VehicleNumber = &tmpStr
						}
						if vv.MemberID == nil {
							vv.MemberID = &tmpDefaultValue
						} else {
							tmpStr := fmt.Sprintf("'%s'", *vv.MemberID)
							vv.MemberID = &tmpStr
						}
						if vv.IsActive == nil {
							vv.IsActive = &tmpDefaultValue
						} else {
							tmpStr := fmt.Sprintf("'%s'", *vv.IsActive)
							vv.IsActive = &tmpStr
						}
						if vv.ActivationDate == nil {
							vv.ActivationDate = &tmpDefaultValue
						} else {
							tmpStr := fmt.Sprintf("'%s'", *vv.ActivationDate)
							vv.ActivationDate = &tmpStr
						}
						if vv.DeviceStatus == nil {
							vv.DeviceStatus = &tmpDefaultValue
						} else {
							tmpStr := fmt.Sprintf("'%s'", *vv.DeviceStatus)
							vv.DeviceStatus = &tmpStr
						}
						if vv.PostDealerID == nil {
							vv.PostDealerID = &tmpDefaultValue
						} else {
							tmpStr := fmt.Sprintf("'%s'", *vv.PostDealerID)
							vv.PostDealerID = &tmpStr
						}
						if vv.ActvDealerID == nil {
							vv.ActvDealerID = &tmpDefaultValue
						} else {
							tmpStr := fmt.Sprintf("'%s'", *vv.ActvDealerID)
							vv.ActvDealerID = &tmpStr
						}
						if vv.GsmNumber == nil {
							vv.GsmNumber = &tmpDefaultValue
						} else {
							tmpStr := fmt.Sprintf("'%s'", *vv.GsmNumber)
							vv.GsmNumber = &tmpStr
						}
						if vv.EngineNumber == nil {
							vv.EngineNumber = &tmpDefaultValue
						} else {
							tmpStr := fmt.Sprintf("'%s'", *vv.EngineNumber)
							vv.EngineNumber = &tmpStr
						}
						if vv.TypeId == nil {
							vv.TypeId = &tmpDefaultValue
						} else {
							tmpStr := fmt.Sprintf("'%s'", *vv.TypeId)
							vv.TypeId = &tmpStr
						}

						tmpQuery := fmt.Sprintf(query.CreateMtVehicleV2, helper.StringToInt(tmpVehicleId, 0),
							vv.ChassisNumber, *vv.Imei, *vv.VehicleName, *vv.VehicleNumber, *vv.MemberID,
							*vv.IsActive, *vv.ActivationDate, *vv.DeviceStatus, *vv.PostDealerID, *vv.ActvDealerID, *vv.GsmNumber, *vv.EngineNumber, *vv.TypeId)

						tmpForm, _ := json.Marshal(val)
						tmpOauthRunner := lib.PayloadNsq{
							RequestID:    uuid.New().String(),
							Time:         uttime.Now().String(),
							Service:      "vehicle",
							DatabaseName: os.Getenv("DBV2_RUNNER"),
							Payload:      string(tmpForm),
							Query:        tmpQuery,
						}
						tmpByteOauthRunner, _ := json.Marshal(tmpOauthRunner)
						_ = lib.SendNSQUsecase(v.NSQ, tmpByteOauthRunner)
					}
				}
			}
		}
	}
	return "", nil
}

func (v vehicleUsecase) UpdateMtVehicle(form []byte) (result string, serr serror.SError) {
	val := map[string]interface{}{}
	err := json.Unmarshal(form, &val)
	if err != nil {
		val2 := []models.VehicleReq{}
		_ = json.Unmarshal(form, &val2)
		for _, vr := range val2 {
			tmpQueryVehicle := fmt.Sprintf(query.GetVehicle, vr.ChassisNumber)
			rows, err := v.DB.Queryx(tmpQueryVehicle)
			if err != nil {
				return result, serror.New(err.Error())
			}

			var tmpVehicleHistoryArray []models.VehicleV3
			defer rows.Close()
			for rows.Next() {
				tmpVehicleHistory := models.VehicleV3{}
				if err := rows.StructScan(&tmpVehicleHistory); err != nil {
					fmt.Println(err.Error())
				}
				tmpVehicleHistoryArray = append(tmpVehicleHistoryArray, tmpVehicleHistory)
			}

			for _, vv := range tmpVehicleHistoryArray {
				if vv.ChassisNumber != "" {
					tmpDefaultValue := "NULL"
					if vv.ChassisNumber == "" {
						vv.ChassisNumber = tmpDefaultValue
					} else {
						tmpStr := fmt.Sprintf("'%s'", vv.ChassisNumber)
						vv.ChassisNumber = tmpStr
					}
					if vv.Imei == nil {
						vv.Imei = &tmpDefaultValue
					} else {
						tmpStr := fmt.Sprintf("'%s'", *vv.Imei)
						vv.Imei = &tmpStr
					}
					if vv.VehicleName == nil {
						vv.VehicleName = &tmpDefaultValue
					} else {
						tmpStr := fmt.Sprintf("'%s'", *vv.VehicleName)
						vv.VehicleName = &tmpStr
					}
					if vv.VehicleNumber == nil {
						vv.VehicleNumber = &tmpDefaultValue
					} else {
						tmpStr := fmt.Sprintf("'%s'", *vv.VehicleNumber)
						vv.VehicleNumber = &tmpStr
					}
					if vv.MemberID == nil {
						vv.MemberID = &tmpDefaultValue
					} else {
						tmpStr := fmt.Sprintf("'%s'", *vv.MemberID)
						vv.MemberID = &tmpStr
					}
					if vv.IsActive == nil {
						vv.IsActive = &tmpDefaultValue
					} else {
						tmpStr := fmt.Sprintf("'%s'", *vv.IsActive)
						vv.IsActive = &tmpStr
					}
					if vv.ActivationDate == nil {
						vv.ActivationDate = &tmpDefaultValue
					} else {
						tmpStr := fmt.Sprintf("'%s'", *vv.ActivationDate)
						vv.ActivationDate = &tmpStr
					}
					if vv.DeviceStatus == nil {
						vv.DeviceStatus = &tmpDefaultValue
					} else {
						tmpStr := fmt.Sprintf("'%s'", *vv.DeviceStatus)
						vv.DeviceStatus = &tmpStr
					}
					if vv.PostDealerID == nil {
						vv.PostDealerID = &tmpDefaultValue
					} else {
						tmpStr := fmt.Sprintf("'%s'", *vv.PostDealerID)
						vv.PostDealerID = &tmpStr
					}
					if vv.ActvDealerID == nil {
						vv.ActvDealerID = &tmpDefaultValue
					} else {
						tmpStr := fmt.Sprintf("'%s'", *vv.ActvDealerID)
						vv.ActvDealerID = &tmpStr
					}
					if vv.GsmNumber == nil {
						vv.GsmNumber = &tmpDefaultValue
					} else {
						tmpStr := fmt.Sprintf("'%s'", *vv.GsmNumber)
						vv.GsmNumber = &tmpStr
					}
					if vv.EngineNumber == nil {
						vv.EngineNumber = &tmpDefaultValue
					} else {
						tmpStr := fmt.Sprintf("'%s'", *vv.EngineNumber)
						vv.EngineNumber = &tmpStr
					}
					if vv.TypeId == nil {
						vv.TypeId = &tmpDefaultValue
					} else {
						tmpStr := fmt.Sprintf("'%s'", *vv.TypeId)
						vv.TypeId = &tmpStr
					}

					tmpQuery := fmt.Sprintf(query.UpdateMtVehcileV2,
						vv.ChassisNumber,
						*vv.Imei,
						*vv.VehicleName,
						*vv.VehicleNumber,
						*vv.MemberID,
						*vv.IsActive,
						*vv.ActivationDate,
						*vv.DeviceStatus,
						*vv.PostDealerID,
						*vv.ActvDealerID,
						*vv.GsmNumber,
						*vv.EngineNumber,
						*vv.TypeId,
						vv.ChassisNumber,
						*vv.MemberID)

					tmpForm, _ := json.Marshal(val)
					tmpOauthRunner := lib.PayloadNsq{
						RequestID:    uuid.New().String(),
						Time:         uttime.Now().String(),
						Service:      "vehicle",
						DatabaseName: os.Getenv("DBV2_RUNNER"),
						Payload:      string(tmpForm),
						Query:        tmpQuery,
					}
					tmpByteOauthRunner, _ := json.Marshal(tmpOauthRunner)
					_ = lib.SendNSQUsecase(v.NSQ, tmpByteOauthRunner)
				}
			}
		}
	} else {
		tmpQueryVehicle := fmt.Sprintf(query.GetVehicle, val["chassis_number"].(string))
		rows, err := v.DB.Queryx(tmpQueryVehicle)
		if err != nil {
			return result, serror.New(err.Error())
		}

		var tmpVehicleHistoryArray []models.VehicleV3
		defer rows.Close()
		for rows.Next() {
			tmpVehicleHistory := models.VehicleV3{}
			if err := rows.StructScan(&tmpVehicleHistory); err != nil {
				fmt.Println(err.Error())
			}
			tmpVehicleHistoryArray = append(tmpVehicleHistoryArray, tmpVehicleHistory)
		}

		for _, vv := range tmpVehicleHistoryArray {
			if vv.ChassisNumber != "" {
				tmpDefaultValue := "NULL"
				if vv.ChassisNumber == "" {
					vv.ChassisNumber = tmpDefaultValue
				} else {
					tmpStr := fmt.Sprintf("'%s'", vv.ChassisNumber)
					vv.ChassisNumber = tmpStr
				}
				if vv.Imei == nil {
					vv.Imei = &tmpDefaultValue
				} else {
					tmpStr := fmt.Sprintf("'%s'", *vv.Imei)
					vv.Imei = &tmpStr
				}
				if vv.VehicleName == nil {
					vv.VehicleName = &tmpDefaultValue
				} else {
					tmpStr := fmt.Sprintf("'%s'", *vv.VehicleName)
					vv.VehicleName = &tmpStr
				}
				if vv.VehicleNumber == nil {
					vv.VehicleNumber = &tmpDefaultValue
				} else {
					tmpStr := fmt.Sprintf("'%s'", *vv.VehicleNumber)
					vv.VehicleNumber = &tmpStr
				}
				if vv.MemberID == nil {
					vv.MemberID = &tmpDefaultValue
				} else {
					tmpStr := fmt.Sprintf("'%s'", *vv.MemberID)
					vv.MemberID = &tmpStr
				}
				if vv.IsActive == nil {
					vv.IsActive = &tmpDefaultValue
				} else {
					tmpStr := fmt.Sprintf("'%s'", *vv.IsActive)
					vv.IsActive = &tmpStr
				}
				if vv.ActivationDate == nil {
					vv.ActivationDate = &tmpDefaultValue
				} else {
					tmpStr := fmt.Sprintf("'%s'", *vv.ActivationDate)
					vv.ActivationDate = &tmpStr
				}
				if vv.DeviceStatus == nil {
					vv.DeviceStatus = &tmpDefaultValue
				} else {
					tmpStr := fmt.Sprintf("'%s'", *vv.DeviceStatus)
					vv.DeviceStatus = &tmpStr
				}
				if vv.PostDealerID == nil {
					vv.PostDealerID = &tmpDefaultValue
				} else {
					tmpStr := fmt.Sprintf("'%s'", *vv.PostDealerID)
					vv.PostDealerID = &tmpStr
				}
				if vv.ActvDealerID == nil {
					vv.ActvDealerID = &tmpDefaultValue
				} else {
					tmpStr := fmt.Sprintf("'%s'", *vv.ActvDealerID)
					vv.ActvDealerID = &tmpStr
				}
				if vv.GsmNumber == nil {
					vv.GsmNumber = &tmpDefaultValue
				} else {
					tmpStr := fmt.Sprintf("'%s'", *vv.GsmNumber)
					vv.GsmNumber = &tmpStr
				}
				if vv.EngineNumber == nil {
					vv.EngineNumber = &tmpDefaultValue
				} else {
					tmpStr := fmt.Sprintf("'%s'", *vv.EngineNumber)
					vv.EngineNumber = &tmpStr
				}
				if vv.TypeId == nil {
					vv.TypeId = &tmpDefaultValue
				} else {
					tmpStr := fmt.Sprintf("'%s'", *vv.TypeId)
					vv.TypeId = &tmpStr
				}

				tmpQuery := fmt.Sprintf(query.UpdateMtVehcileV2,
					vv.ChassisNumber,
					*vv.Imei,
					*vv.VehicleName,
					*vv.VehicleNumber,
					*vv.MemberID,
					*vv.IsActive,
					*vv.ActivationDate,
					*vv.DeviceStatus,
					*vv.PostDealerID,
					*vv.ActvDealerID,
					*vv.GsmNumber,
					*vv.EngineNumber,
					*vv.TypeId,
					vv.ChassisNumber,
					*vv.MemberID)

				tmpForm, _ := json.Marshal(val)
				tmpOauthRunner := lib.PayloadNsq{
					RequestID:    uuid.New().String(),
					Time:         uttime.Now().String(),
					Service:      "vehicle",
					DatabaseName: os.Getenv("DBV2_RUNNER"),
					Payload:      string(tmpForm),
					Query:        tmpQuery,
				}
				tmpByteOauthRunner, _ := json.Marshal(tmpOauthRunner)
				_ = lib.SendNSQUsecase(v.NSQ, tmpByteOauthRunner)
			}
		}
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
		DatabaseName: os.Getenv("DBV2_RUNNER"),
		Payload:      string(tmpForm),
		Query:        tmpQuery,
	}

	tmpByteOauthRunner, _ := json.Marshal(tmpOauthRunner)
	_ = lib.SendNSQUsecase(v.NSQ, tmpByteOauthRunner)

	tmpChassisNumber := val["chassis_number"].(string)
	tmpChassisNumberArray := strings.Split(tmpChassisNumber, ",")
	for _, vCN := range tmpChassisNumberArray {
		var tmpVehicleId int
		if err := v.DB.QueryRow(fmt.Sprintf(query.GetVehicleIdByChassis, vCN)).Scan(&tmpVehicleId); err != nil {
			fmt.Println(err.Error())
		}

		tmpQuery = fmt.Sprintf(query.CreateMSVehicleGroup,
			helper.StringToInt(tmpGroupId, 0),
			tmpVehicleId,
		)

		tmpForm, _ = json.Marshal(val)
		tmpOauthRunner = lib.PayloadNsq{
			RequestID:    uuid.New().String(),
			Time:         uttime.Now().String(),
			Service:      "vehicle",
			DatabaseName: os.Getenv("DBV2_RUNNER"),
			Payload:      string(tmpForm),
			Query:        tmpQuery,
		}

		tmpByteOauthRunner, _ = json.Marshal(tmpOauthRunner)
		_ = lib.SendNSQUsecase(v.NSQ, tmpByteOauthRunner)
	}

	return "", nil
}

func (v vehicleUsecase) UpdateVehicleGroup(form []byte) (result string, serr serror.SError) {
	val := map[string]interface{}{}
	_ = json.Unmarshal(form, &val)

	var memberId string
	if err := v.DB.QueryRow(fmt.Sprintf(query.GetMemberIdByOrganizationId, val["owner_id"].(string))).Scan(&memberId); err != nil {
		fmt.Println(err.Error())
	}

	var vehicleGroup string
	if err := v.DB.QueryRow(fmt.Sprintf(query.SelectVehicleGroup, memberId, val["vehicle_group_name"].(string))).Scan(&vehicleGroup); err != nil {
		fmt.Println(err.Error())
	}

	if _, err := v.DB.Queryx(fmt.Sprintf(query.DeleteVehicleGroup, vehicleGroup)); err != nil {
		fmt.Println(err.Error())
	}

	if _, err := v.DB.Queryx(fmt.Sprintf(query.DeleteMsVehicleGroup, vehicleGroup)); err != nil {
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
		DatabaseName: os.Getenv("DBV2_RUNNER"),
		Payload:      string(tmpForm),
		Query:        tmpQuery,
	}

	tmpByteOauthRunner, _ := json.Marshal(tmpOauthRunner)
	_ = lib.SendNSQUsecase(v.NSQ, tmpByteOauthRunner)

	tmpChassisNumber := val["chassis_number"].(string)
	tmpChassisNumberArray := strings.Split(tmpChassisNumber, ",")
	for _, vCN := range tmpChassisNumberArray {

		var tmpVehicleId int
		if err := v.DB.QueryRow(fmt.Sprintf(query.GetVehicleIdByChassis, vCN)).Scan(&tmpVehicleId); err != nil {
			fmt.Println(err.Error())
		}

		tmpQuery = fmt.Sprintf(query.CreateMSVehicleGroup,
			helper.StringToInt(tmpGroupId, 0),
			tmpVehicleId,
		)

		tmpForm, _ = json.Marshal(val)
		tmpOauthRunner = lib.PayloadNsq{
			RequestID:    uuid.New().String(),
			Time:         uttime.Now().String(),
			Service:      "vehicle",
			DatabaseName: os.Getenv("DBV2_RUNNER"),
			Payload:      string(tmpForm),
			Query:        tmpQuery,
		}

		tmpByteOauthRunner, _ = json.Marshal(tmpOauthRunner)
		_ = lib.SendNSQUsecase(v.NSQ, tmpByteOauthRunner)
	}

	return "", nil
}
