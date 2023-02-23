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

		tmpMemberId := lib.RandomCharacter(18)
		if tmpUser.RealmId != nil {
			if *tmpUser.RealmId == lib.RealmIdCustomer {
				c.DB.QueryRow(query.CreateMemberTmp, tmpOrganization.ID, tmpMemberId, nil)
			}
		}

		fmt.Println(tmpOrganization.ID)
		fmt.Println(tmpMemberId)

		tmpQuery := fmt.Sprintf(query.CreateMtMember,
			helper.StringToInt(tmpMemberId, 0),
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

		_, _ = c.CreateAuthRunnerUsecase(form)
	}

	return "", nil
}

func (c customerUsecase) UpdateMtMemberUsecase(form []byte) (result string, serr serror.SError) {
	val := map[string]interface{}{}
	_ = json.Unmarshal(form, &val)

	var tmpUser models.UserResult
	_, ok := val["super_username"].(string)
	if ok {
		tmpUser, _ = c.userRepo.GetUserByUsernameRepo(val["super_username"].(string))
	}

	if tmpUser.ID == "" {
		tmpUser, _ = c.userRepo.GetUserByIDOrganizationRepo(val["id"].(string))
	}

	if tmpUser.OrganizationId == nil {
		fmt.Println("username or id not found")
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
		if tmpOrganization.OrgContactPerson == nil {
			tmpOrganization.OrgContactPerson = &tmpDefaultValue
		}

		if memberId != "" {
			tmpQuery := fmt.Sprintf(query.UpdateMtMember,
				*tmpOrganization.Name,
				*tmpOrganization.OrgContactPerson,
				*tmpOrganization.Email,
				*tmpOrganization.Telp,
				*tmpOrganization.Code,
				*tmpOrganization.BusinessType,
				helper.StringToInt(memberId, 0),
				*tmpOrganization.Email,
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

			_, _ = c.UpdateAuthRunnerUsecase(form)
		}
	}

	return "", nil
}

func (c customerUsecase) CreateAuthRunnerUsecase(form []byte) (result string, serr serror.SError) {
	val := map[string]interface{}{}
	_ = json.Unmarshal(form, &val)

	var tmpUser models.UserResult
	_, ok := val["username"].(string)
	if ok {
		rows, err := c.DB.Queryx(query.GetUserByUsername, val["username"].(string))
		if err != nil {
			return result, serror.New(err.Error())
		}

		defer rows.Close()
		for rows.Next() {
			if err := rows.StructScan(&tmpUser); err != nil {
				return result, serror.New("Failed scan for struct models")
			}
		}
	} else {
		_, ok = val["super_username"].(string)
		if ok {
			rows, err := c.DB.Queryx(query.GetUserByUsername, val["super_username"].(string))
			if err != nil {
				return result, serror.New(err.Error())
			}

			defer rows.Close()
			for rows.Next() {
				if err := rows.StructScan(&tmpUser); err != nil {
					return result, serror.New("Failed scan for struct models")
				}
			}
		}
	}

	tmpOrganization, _ := c.organizationRepo.GetOrganizationByIdRepo(*tmpUser.OrganizationId)

	var dealerId *string
	var memberId *string
	_, tmpRealm := val["realm_id"].(string)
	if tmpRealm {
		if val["realm_id"].(string) == lib.RealmIdDealer {
			if err := c.DB.QueryRow(fmt.Sprintf(query.GetDealerIdByOrganizationId, val["organization_id"].(string))).Scan(&dealerId); err != nil {
				fmt.Println(err.Error())
			}
		} else if val["realm_id"].(string) == lib.RealmIdCustomer {
			if tmpOrganization.ParentID != nil {
				if err := c.DB.QueryRow(fmt.Sprintf(query.GetDealerIdByOrganizationId, *tmpOrganization.ParentID)).Scan(&dealerId); err != nil {
					fmt.Println(err.Error())
				}
			}

			if err := c.DB.QueryRow(fmt.Sprintf(query.GetMemberIdByOrganizationId, *tmpUser.OrganizationId)).Scan(&memberId); err != nil {
				fmt.Println(err.Error())
			}

			if memberId == nil {
				if err := c.DB.QueryRow(fmt.Sprintf(query.GetMemberIdByEmailAndMemberName, *tmpUser.OrganizationEmail, *tmpUser.OrganizationName)).Scan(&memberId); err != nil {
					fmt.Println(err.Error())
				}
			}
		}
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

	if tmpUser.RealmId != nil {
		if *tmpUser.RealmId == lib.RealmIdDealer {
			if dealerId != nil {
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
						tmpDefaultValue = "NULL"
						dealerId = &tmpDefaultValue
					} else {
						tmpStr := fmt.Sprintf("%s", *dealerId)
						dealerId = &tmpStr
					}
					if memberId == nil {
						tmpDefaultValue = "NULL"
						memberId = &tmpDefaultValue
					} else {
						tmpStr := fmt.Sprintf("%s", *memberId)
						memberId = &tmpStr
					}
					tmpUserId := lib.RandomCharacter(18)
					tmpQuery := fmt.Sprintf(query.CreateAuthRunner,
						helper.StringToInt(tmpUserId, 0),
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

					if tmpRole == "13" || tmpRole == "14" {
						_, _ = c.PrivacyPolicyUsecase(form, tmpUserId)
					}
				}
			}
		} else if *tmpUser.RealmId == lib.RealmIdCustomer {
			if dealerId != nil {
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
							tmpDefaultValue = "NULL"
							dealerId = &tmpDefaultValue
						} else {
							tmpStr := fmt.Sprintf("%s", *dealerId)
							dealerId = &tmpStr
						}
						if memberId == nil {
							tmpDefaultValue = "NULL"
							memberId = &tmpDefaultValue
						} else {
							tmpStr := fmt.Sprintf("%s", *memberId)
							memberId = &tmpStr
						}
						tmpUserId := lib.RandomCharacter(18)
						tmpQuery := fmt.Sprintf(query.CreateAuthRunner,
							helper.StringToInt(tmpUserId, 0),
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

						if tmpRole == "13" || tmpRole == "14" {
							_, _ = c.PrivacyPolicyUsecase(form, tmpUserId)
						}
					}
				}
			}
		}
	}
	return "", nil
}

func (c customerUsecase) UpdateAuthRunnerUsecase(form []byte) (result string, serr serror.SError) {
	val := map[string]interface{}{}
	_ = json.Unmarshal(form, &val)

	var tmpOrganization models.RpcListOrganizationCustomerResult
	var tmpUser models.UserResult
	_, ok := val["organization_id"].(string)
	if ok {
		tmpOrganization, _ = c.organizationRepo.GetOrganizationByIdRepo(val["organization_id"].(string))
	} else {
		_, ok = val["user_id"].(string)
		if ok {
			rows, err := c.DB.Queryx(query.GetUserByUserId, val["user_id"].(string))
			if err != nil {
				fmt.Println(err.Error())
				return result, serror.New(err.Error())
			}

			defer rows.Close()
			for rows.Next() {
				if err := rows.StructScan(&tmpUser); err != nil {
					return result, serror.New("Failed scan for struct models")
				}
			}

			if tmpUser.OrganizationId != nil {
				tmpOrganization, _ = c.organizationRepo.GetOrganizationByIdRepo(*tmpUser.OrganizationId)
			}
		}
	}
	if tmpOrganization.ID == "" {
		return "", nil
	}

	var dealerId *string
	_, tmpRealm := val["realm_id"].(string)
	if tmpRealm {
		if val["realm_id"].(string) == lib.RealmIdDealer {
			if err := c.DB.QueryRow(fmt.Sprintf(query.GetDealerIdByOrganizationId, val["organization_id"].(string))).Scan(&dealerId); err != nil {
				fmt.Println(err.Error())
			}
		} else if val["realm_id"].(string) == lib.RealmIdCustomer {
			if tmpOrganization.ParentID != nil {
				if err := c.DB.QueryRow(fmt.Sprintf(query.GetDealerIdByOrganizationId, *tmpOrganization.ParentID)).Scan(&dealerId); err != nil {
					fmt.Println(err.Error())
				}
			}
		}
	}

	var memberId *string
	if err := c.DB.QueryRow(fmt.Sprintf(query.GetMemberIdByOrganizationId, tmpOrganization.ID)).Scan(&memberId); err != nil {
		fmt.Println(err.Error())
	}

	var tmpUserId string
	if tmpUser.ID == "" {
		_, ok = val["username"].(string)
		if ok {
			rows, err := c.DB.Queryx(query.GetUserByUsername, val["username"].(string))
			if err != nil {
				return result, serror.New(err.Error())
			}

			defer rows.Close()
			for rows.Next() {
				if err := rows.StructScan(&tmpUser); err != nil {
					return result, serror.New("Failed scan for struct models")
				}
			}

			_, userBefore := val["username_before"].(string)
			if userBefore {
				if val["username"].(string) != val["username_before"].(string) {
					if err := c.DB.QueryRow(fmt.Sprintf(query.GetUserIdByUsername, val["username_before"].(string))).Scan(&tmpUserId); err != nil {
						fmt.Println(err.Error())
					}
				}
			} else {
				if err := c.DB.QueryRow(fmt.Sprintf(query.GetUserIdByUsername, val["username"].(string))).Scan(&tmpUserId); err != nil {
					fmt.Println(err.Error())
				}
			}

		} else {
			_, ok = val["super_username"].(string)
			if ok {
				rows, err := c.DB.Queryx(query.GetUserByUsername, val["super_username"].(string))
				if err != nil {
					return result, serror.New(err.Error())
				}

				defer rows.Close()
				for rows.Next() {
					if err := rows.StructScan(&tmpUser); err != nil {
						return result, serror.New("Failed scan for struct models")
					}
				}

				_, userBefore := val["username_before"].(string)
				if userBefore {
					if val["super_username"].(string) != val["username_before"].(string) {
						if err := c.DB.QueryRow(fmt.Sprintf(query.GetUserIdByUsername, val["username_before"].(string))).Scan(&tmpUserId); err != nil {
							fmt.Println(err.Error())
						}
					}
				} else {
					if err := c.DB.QueryRow(fmt.Sprintf(query.GetUserIdByUsername, val["super_username"].(string))).Scan(&tmpUserId); err != nil {
						fmt.Println(err.Error())
					}
				}
			}
		}
	} else {
		tmpUserId = tmpUser.ID
	}

	if tmpUserId == "" {
		_, _ = c.CreateAuthRunnerUsecase(form)
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

	if tmpUser.RealmId != nil {
		if *tmpUser.RealmId == lib.RealmIdDealer {
			if dealerId != nil {
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
						tmpDefaultValue = "NULL"
						dealerId = &tmpDefaultValue
					} else {
						tmpStr := fmt.Sprintf("%s", *dealerId)
						dealerId = &tmpStr
					}
					tmpUserUpdate := ""
					_, userBefore := val["username_before"].(string)
					if userBefore {
						tmpUserUpdate = val["username_before"].(string)
						if tmpUserUpdate == "" {
							tmpUserUpdate = *tmpUser.Username
						}
					} else {
						tmpUserUpdate = *tmpUser.Username
					}
					if memberId == nil {
						tmpDefaultValue = "NULL"
						memberId = &tmpDefaultValue
					} else {
						tmpStr := fmt.Sprintf("%s", *memberId)
						memberId = &tmpStr
					}
					tmpQuery := fmt.Sprintf(query.UpdateAuthRunner,
						*tmpUser.Username,
						*tmpUser.Password,
						*tmpUser.ProfileName,
						*tmpUser.ProfileEmail,
						tmpRole,
						*memberId,
						*dealerId,
						*tmpUser.ProfilePhone,
						tmpUserUpdate,
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
		} else if *tmpUser.RealmId == lib.RealmIdCustomer {
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
							tmpDefaultValue = "NULL"
							dealerId = &tmpDefaultValue
						} else {
							tmpStr := fmt.Sprintf("%s", *dealerId)
							dealerId = &tmpStr
						}
						tmpUserUpdate := ""
						_, userBefore := val["username_before"].(string)
						if userBefore {
							tmpUserUpdate = val["username_before"].(string)
							if tmpUserUpdate == "" {
								tmpUserUpdate = *tmpUser.Username
							}
						} else {
							tmpUserUpdate = *tmpUser.Username
						}
						if memberId == nil {
							tmpDefaultValue = "NULL"
							memberId = &tmpDefaultValue
						} else {
							tmpStr := fmt.Sprintf("%s", *memberId)
							memberId = &tmpStr
						}
						tmpQuery := fmt.Sprintf(query.UpdateAuthRunner,
							*tmpUser.Username,
							*tmpUser.Password,
							*tmpUser.ProfileName,
							*tmpUser.ProfileEmail,
							tmpRole,
							*memberId,
							*dealerId,
							*tmpUser.ProfilePhone,
							tmpUserUpdate,
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

func (c customerUsecase) PrivacyPolicyUsecase(form []byte, tmpUserId string) (result string, serr serror.SError) {
	val := map[string]interface{}{}
	_ = json.Unmarshal(form, &val)

	var tmpPrivacyPolicy bool
	if err := c.DB.QueryRow(fmt.Sprintf(query.CheckPrivacyPolicy, tmpUserId)).Scan(&tmpPrivacyPolicy); err != nil {
		fmt.Println(err.Error())
	}

	if !tmpPrivacyPolicy {
		tmpIdPrivacyPolicy := lib.RandomCharacter(18)
		tmpQuery := fmt.Sprintf(query.CreatePrivacyPolicy,
			tmpIdPrivacyPolicy,
			tmpUserId,
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

	return "", nil
}

func (c customerUsecase) ApprovePrivacyPolicyUsecase(form []byte) (result string, serr serror.SError) {
	val := map[string]interface{}{}
	_ = json.Unmarshal(form, &val)

	var tmpUserId string
	if err := c.DB.QueryRow(fmt.Sprintf(query.GetUserIdByUsername, val["username"].(string))).Scan(&tmpUserId); err != nil {
		fmt.Println(err.Error())
	}

	if tmpUserId != "" {
		var tmpPrivacyPolicy bool
		if err := c.DB.QueryRow(fmt.Sprintf(query.CheckPrivacyPolicy, tmpUserId)).Scan(&tmpPrivacyPolicy); err != nil {
			fmt.Println(err.Error())
		}

		if tmpPrivacyPolicy {
			tmpQuery := fmt.Sprintf(query.DeletePrivacyPolicy,
				tmpUserId,
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

	return "", nil
}

func (c customerUsecase) DeleteAuthRunnerUsecase(form []byte) (result string, serr serror.SError) {
	val := map[string]interface{}{}
	_ = json.Unmarshal(form, &val)

	var tmpUser models.UserResult
	_, ok := val["id"].(string)
	if ok {
		_, ok = val["id"].(string)
		if ok {
			rows, err := c.DB.Queryx(query.GetUserByUserIdDeleted, val["id"].(string))
			if err != nil {
				fmt.Println(err.Error())
				return result, serror.New(err.Error())
			}

			defer rows.Close()
			for rows.Next() {
				if err := rows.StructScan(&tmpUser); err != nil {
					return result, serror.New("Failed scan for struct models")
				}
			}
		}
	}

	if tmpUser.ID == "" {
		return "", nil
	}

	if tmpUser.Username != nil {
		var tmpUserId string
		if err := c.DB.QueryRow(fmt.Sprintf(query.GetUserIdByUsername, *tmpUser.Username)).Scan(&tmpUserId); err != nil {
			fmt.Println(err.Error())
		}

		if tmpUserId != "" {
			tmpQuery := fmt.Sprintf(query.DeleteAuthRunner,
				tmpUserId,
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

	return "", nil
}
