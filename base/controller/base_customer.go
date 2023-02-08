package controller

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	lib "github.com/iikmaulana/mini-replacement/base/libs"
	"github.com/iikmaulana/mini-replacement/base/service"
	"github.com/iikmaulana/mini-replacement/base/service/repository/query"
	"github.com/jmoiron/sqlx"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
	"github.com/uzzeet/uzzeet-gateway/libs/utils/uttime"
	"reflect"
)

type customerUsecase struct {
	DB               *sqlx.DB
	userRepo         service.UserRepo
	organizationRepo service.OrganizationRepo
}

func NewCustomerUsecase(db *sqlx.DB, userRepo service.UserRepo, organizationRepo service.OrganizationRepo) service.CustomerUsecase {
	return customerUsecase{DB: db, userRepo: userRepo, organizationRepo: organizationRepo}
}

func (c customerUsecase) CreateMtMemberUsecase(form []byte) (result string, serr serror.SError) {
	val := map[string]interface{}{}
	_ = json.Unmarshal(form, &val)

	tmpUser, _ := c.userRepo.GetUserByUsernameRepo(val["super_username"].(string))
	if tmpUser.OrganizationId == nil {
		return "", nil
	}

	tmpOrganization, _ := c.organizationRepo.GetOrganizationByIdRepo(*tmpUser.OrganizationId)

	var dealerId *string
	if tmpOrganization.ParentID != nil {
		if err := c.DB.QueryRow(fmt.Sprintf(query.GetDealerIdByOrganizationId, *tmpOrganization.ParentID)).Scan(&dealerId); err != nil {
			fmt.Println("Dealer ID " + err.Error())
		}
	}

	if tmpOrganization.ID != "" {
		tmpDefaultValue := ""
		if tmpOrganization.Name == nil {
			tmpOrganization.Name = &tmpDefaultValue
		}
		if tmpOrganization.ContactName == nil {
			tmpOrganization.ContactName = &tmpDefaultValue
		}
		if tmpOrganization.Email == nil {
			tmpOrganization.Email = &tmpDefaultValue
		}
		if tmpOrganization.Telp == nil {
			tmpOrganization.Telp = &tmpDefaultValue
		}
		if dealerId == nil {
			dealerId = &tmpDefaultValue
		}
		if tmpOrganization.Code == nil {
			tmpOrganization.Code = &tmpDefaultValue
		}
		if tmpOrganization.BusinessType == nil {
			tmpOrganization.BusinessType = &tmpDefaultValue
		}

		tmpQuery := fmt.Sprintf(query.CreateMtMember,
			*tmpOrganization.Name,
			*tmpOrganization.ContactName,
			*tmpOrganization.Email,
			*tmpOrganization.Telp,
			*dealerId,
			*tmpOrganization.Code,
			*tmpOrganization.BusinessType,
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

func (c customerUsecase) UpdateMtMemberUsecase(form []byte) (result string, serr serror.SError) {
	val := map[string]interface{}{}
	_ = json.Unmarshal(form, &val)

	tmpUser, _ := c.userRepo.GetUserByUsernameRepo(val["super_username"].(string))
	if tmpUser.OrganizationId == nil {
		return "", nil
	}

	tmpOrganization, _ := c.organizationRepo.GetOrganizationByIdRepo(*tmpUser.OrganizationId)

	var memberId string
	if err := c.DB.QueryRow(fmt.Sprintf(query.GetMemberIdByOrganizationId, *tmpUser.OrganizationId)).Scan(&memberId); err != nil {
		fmt.Println(err.Error())
	}

	if tmpOrganization.ID != "" {
		tmpDefaultValue := ""
		if tmpOrganization.Name == nil {
			tmpOrganization.Name = &tmpDefaultValue
		}
		if tmpOrganization.ContactName == nil {
			tmpOrganization.ContactName = &tmpDefaultValue
		}
		if tmpOrganization.Email == nil {
			tmpOrganization.Email = &tmpDefaultValue
		}
		if tmpOrganization.Telp == nil {
			tmpOrganization.Telp = &tmpDefaultValue
		}
		if tmpOrganization.Code == nil {
			tmpOrganization.Code = &tmpDefaultValue
		}
		if tmpOrganization.BusinessType == nil {
			tmpOrganization.BusinessType = &tmpDefaultValue
		}

		if memberId != "" {
			tmpQuery := fmt.Sprintf(query.UpdateMtMember,
				*tmpOrganization.Name,
				*tmpOrganization.Name,
				*tmpOrganization.Email,
				*tmpOrganization.Telp,
				*tmpOrganization.Code,
				*tmpOrganization.BusinessType,
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
	}

	return "", nil
}

func (c customerUsecase) CreateAuthRunnerUsecase(form []byte) (result string, serr serror.SError) {
	val := map[string]interface{}{}
	_ = json.Unmarshal(form, &val)

	tmpUser, _ := c.userRepo.GetUserByUsernameRepo(val["super_username"].(string))
	if tmpUser.OrganizationId == nil {
		return "", nil
	}

	tmpOrganization, _ := c.organizationRepo.GetOrganizationByIdRepo(*tmpUser.OrganizationId)

	var dealerId *string
	if tmpOrganization.ParentID != nil {
		if err := c.DB.QueryRow(fmt.Sprintf(query.GetDealerIdByOrganizationId, *tmpOrganization.ParentID)).Scan(&dealerId); err != nil {
			fmt.Println(err.Error())
		}
	}

	var memberId *string
	if err := c.DB.QueryRow(fmt.Sprintf(query.GetMemberIdByOrganizationId, *tmpUser.OrganizationId)).Scan(&memberId); err != nil {
		fmt.Println(err.Error())
	}

	var tmpRole string
	if tmpUser.RealmId != nil {
		if *tmpUser.RealmId == lib.RealmIdDealer {
			if tmpUser.AccessType != nil {
				if *tmpUser.AccessType == 2 {
					tmpRole = "13"
				} else if *tmpUser.AccessType == 3 {
					tmpRole = "25"
				}
			}
		} else if *tmpUser.RealmId == lib.RealmIdCustomer {
			if tmpUser.AccessType != nil {
				if *tmpUser.AccessType == 2 {
					tmpRole = "14"
				} else if *tmpUser.AccessType == 3 {
					tmpRole = "27"
				}
			}
		}
	}

	if memberId != nil {
		if *memberId != "" {
			if tmpUser.ID != "" {
				tmpDefaultValue := ""
				if tmpUser.Username == nil {
					tmpUser.Username = &tmpDefaultValue
				}
				if tmpUser.Password == nil {
					tmpUser.Password = &tmpDefaultValue
				}
				if tmpUser.ProfileName == nil {
					tmpUser.ProfileName = &tmpDefaultValue
				}
				if tmpUser.ProfileEmail == nil {
					tmpUser.ProfileEmail = &tmpDefaultValue
				}
				if tmpUser.ProfilePhone == nil {
					tmpUser.ProfilePhone = &tmpDefaultValue
				}
				if dealerId == nil {
					dealerId = &tmpDefaultValue
				}
				tmpQuery := fmt.Sprintf(query.CreateAuthRunner,
					"id",
					*tmpUser.Username,
					*tmpUser.Password,
					*tmpUser.ProfileName,
					*tmpUser.ProfileEmail,
					tmpRole,
					*memberId,
					*dealerId,
					*tmpUser.ProfilePhone,
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
		}
	}
	return "", nil
}

func (c customerUsecase) UpdateAuthRunnerUsecase(form []byte) (result string, serr serror.SError) {
	val := map[string]interface{}{}
	_ = json.Unmarshal(form, &val)

	tmpUser, _ := c.userRepo.GetUserByUsernameRepo(val["super_username"].(string))
	if tmpUser.OrganizationId == nil {
		return "", nil
	}

	tmpOrganization, _ := c.organizationRepo.GetOrganizationByIdRepo(*tmpUser.OrganizationId)

	var dealerId *string
	if tmpOrganization.ParentID != nil {
		if err := c.DB.QueryRow(fmt.Sprintf(query.GetDealerIdByOrganizationId, *tmpOrganization.ParentID)).Scan(&dealerId); err != nil {
			fmt.Println(err.Error())
		}
	}

	var memberId *string
	if err := c.DB.QueryRow(fmt.Sprintf(query.GetMemberIdByOrganizationId, *tmpUser.OrganizationId)).Scan(&memberId); err != nil {
		fmt.Println(err.Error())
	}

	var tmpRole string
	if tmpUser.RealmId != nil {
		if *tmpUser.RealmId == lib.RealmIdDealer {
			if tmpUser.AccessType != nil {
				if *tmpUser.AccessType == 2 {
					tmpRole = "13"
				} else if *tmpUser.AccessType == 3 {
					tmpRole = "25"
				}
			}
		} else if *tmpUser.RealmId == lib.RealmIdCustomer {
			if tmpUser.AccessType != nil {
				if *tmpUser.AccessType == 2 {
					tmpRole = "14"
				} else if *tmpUser.AccessType == 3 {
					tmpRole = "27"
				}
			}
		}
	}

	if memberId != nil {
		if *memberId != "" {
			if tmpUser.ID != "" {
				tmpQuery := fmt.Sprintf(query.UpdateAuthRunner,
					*tmpUser.Username,
					*tmpUser.Password,
					*tmpUser.ProfileName,
					*tmpUser.ProfileEmail,
					tmpRole,
					*memberId,
					*dealerId,
					*tmpUser.ProfilePhone,
					*tmpUser.Username,
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
		}
	}
	return "", nil
}

func (c customerUsecase) UserActivationUsecase(form []byte) (result string, serr serror.SError) {
	val := reflect.ValueOf(form).Elem()

	var memberId *int64
	if err := c.DB.QueryRow(fmt.Sprintf(query.GetMemberIdByOrganizationId, val.FieldByName("id").Interface().(string))).Scan(&memberId); err != nil {
		fmt.Println(err.Error())
	}

	tmpUser, _ := c.userRepo.GetUserByUsernameRepo(val.FieldByName("username").Interface().(string))

	if memberId != nil {
		if *memberId != 0 {
			if tmpUser.ID != "" {
				tmpQuery := fmt.Sprintf(query.UpdateAuthRunner,
					*tmpUser.Username,
				)

				tmpForm, _ := json.Marshal(form)
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
		}
	}
	return "", nil
}

func (c customerUsecase) ChangePasswordUsecase(form []byte) (result string, serr serror.SError) {
	val := map[string]interface{}{}
	_ = json.Unmarshal(form, &val)

	tmpQuery := fmt.Sprintf(query.ChangePassword,
		val["password"].(string),
		val["username"].(string),
	)

	tmpForm, _ := json.Marshal(form)
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

	return "", nil
}
