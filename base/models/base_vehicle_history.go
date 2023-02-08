package models

type GetVehHistoryReq struct {
	CountPage         string `json:"count_perpage" valid:"required"`
	Page              string `json:"page" valid:"required"`
	OwnerID           string `json:"owner_id"`
	DealerID          string `json:"dealer_id"`
	RealmID           string `json:"realm_id"`
	IsOrgAdmin        string `json:"isorgadmin"`
	UserID            string `json:"user_id"`
	Search            string `json:"search"`
	Sort              string `json:"sort"`
	VehicleCategoryId string `json:"vehicle_category_id"`
	VehicleModelId    string `json:"vehicle_model_id"`
	VehicleTypeId     string `json:"vehicle_type_id"`
	ActivationStatus  string `json:"activation_status"`
	Last_service      string `json:"last_service"`
	Next_service      string `json:"next_service"`
	Sales_year        string `json:"sales_year"`
	Header            string `json:"header"`
}

type VehHistoryListMetaResult struct {
	TotalData        int64              `json:"total_data"`
	TotalDataPerpage int                `json:"total_data_perpage"`
	From             int                `json:"from"`
	To               int                `json:"to"`
	TotalPage        int                `json:"total_page"`
	Header           []string           `json:"header"`
	Data             []VehHistoryResult `json:"data"`
}

type VehHistoryResult struct {
	ChassisNumber        string  `json:"chassis_number" db:"chassis_number"`
	EngineNumber         *string `json:"engine_number" db:"engine_number"`
	Imei                 *string `json:"imei" db:"imei"`
	VehicleName          string  `json:"vehicle_name" db:"vehicle_name"`
	VehicleNumber        *string `json:"vehicle_number" db:"vehicle_number"`
	OwnerID              *string `json:"owner_id" db:"owner_id"`
	DealerID             *string `json:"dealer_id" db:"dealer_id"`
	EuroStd              *string `json:"euro_std" db:"euro_std"`
	ActivationDate       *string `json:"activation_date" db:"activation_date"`
	ActivationStatus     string  `json:"activation_status" db:"activation_status"`
	SubscriptionPeriod   *string `json:"subscription_period" db:"subcription_period"`
	StnkDate             *string `json:"stnk_date" db:"stnk_date"`
	KirDate              *string `json:"kir_date" db:"kir_date"`
	Area                 *string `json:"area" db:"area"`
	Odometer             *string `json:"odometer" db:"odometer"`
	BusinessSectorID     *string `json:"business_sector_id" db:"business_sector_id"`
	VehicleTypeId        *string `json:"vehicle_type_id" db:"vehicle_type_id"`
	VehicleTypeName      *string `json:"vehicle_type_name" db:"vehicle_type_name"`
	VehicleModelId       *string `json:"vehicle_model_id" db:"vehicle_model_id"`
	VehicleModelName     *string `json:"vehicle_model_name" db:"vehicle_model_name"`
	VehicleCategoryId    *string `json:"vehicle_category_id" db:"vehicle_category_id"`
	VehicleCategoryName  *string `json:"vehicle_category_name" db:"vehicle_category_name"`
	StatusVehicleHistory string  `json:"status_vehicle_history" db:"status_vehicle_history"`
	ProofOfOwnership     *string `json:"proof_of_ownership" db:"proof_of_ownership"`
	FuelCapacity         *string `json:"fuel_capacity" db:"fuel_capacity"`
	Weight               *string `json:"weight" db:"weight"`
	Length               *string `json:"length" db:"length"`
	Width                *string `json:"width" db:"width"`
	Height               *string `json:"height" db:"height"`
	RearBodyTypeID       *string `json:"rear_body_type_id" db:"rear_body_type_id"`
	RearBodyTypeName     *string `json:"rear_body_type_name" db:"rear_body_type_name"`
	LoadTypeID           *string `json:"load_type_id" db:"load_type_id"`
	LoadTypeName         *string `json:"load_type_name" db:"load_type_name"`
	CreatedAt            string  `json:"created_at" db:"created_at"`
	EndDate              *string `json:"end_date" db:"end_date"`
}

type VehHistoryByChassisResult struct {
	ChassisNumber               string  `json:"chassis_number" db:"chassis_number"`
	EngineNumber                *string `json:"engine_number" db:"engine_number"`
	VehicleName                 string  `json:"vehicle_name" db:"vehicle_name"`
	VehicleNumber               *string `json:"vehicle_number" db:"vehicle_number"`
	OwnerID                     *string `json:"owner_id" db:"owner_id"`
	DealerID                    *string `json:"dealer_id" db:"dealer_id"`
	EuroStd                     *string `json:"euro_std" db:"euro_std"`
	ActivationDate              *string `json:"activation_date" db:"activation_date"`
	ActivationStatus            string  `json:"activation_status" db:"activation_status"`
	SubscriptionPeriod          *string `json:"subscription_period" db:"subcription_period"`
	StnkDate                    *string `json:"stnk_date" db:"stnk_date"`
	KirDate                     *string `json:"kir_date" db:"kir_date"`
	Area                        *string `json:"area" db:"area"`
	Odometer                    *string `json:"odometer" db:"odometer"`
	BusinessSectorID            *string `json:"business_sector_id" db:"business_sector_id"`
	VehicleTypeId               *string `json:"vehicle_type_id" db:"vehicle_type_id"`
	VehicleTypeName             *string `json:"vehicle_type_name" db:"vehicle_type_name"`
	VehicleModelId              *string `json:"vehicle_model_id" db:"vehicle_model_id"`
	VehicleModelName            *string `json:"vehicle_model_name" db:"vehicle_model_name"`
	VehicleCategoryId           *string `json:"vehicle_category_id" db:"vehicle_category_id"`
	VehicleCategoryName         *string `json:"vehicle_category_name" db:"vehicle_category_name"`
	StatusVehicleHistory        string  `json:"status_vehicle_history" db:"status_vehicle_history"`
	ProofOfOwnership            *string `json:"proof_of_ownership" db:"proof_of_ownership"`
	FuelCapacity                *string `json:"fuel_capacity" db:"fuel_capacity"`
	Weight                      *string `json:"weight" db:"weight"`
	Length                      *string `json:"length" db:"length"`
	Width                       *string `json:"width" db:"width"`
	Height                      *string `json:"height" db:"height"`
	RearBodyTypeID              *string `json:"rear_body_type_id" db:"rear_body_type_id"`
	RearBodyTypeName            *string `json:"rear_body_type_name" db:"rear_body_type_name"`
	LoadTypeID                  *string `json:"load_type_id" db:"load_type_id"`
	LoadTypeName                *string `json:"load_type_name" db:"load_type_name"`
	CreatedAt                   string  `json:"created_at" db:"created_at"`
	EndDate                     *string `json:"end_date" db:"end_date"`
	Imei                        *string `json:"imei" db:"imei"`
	GsmNumber                   *string `json:"gsm_number" db:"gsm_number"`
	Iccid                       *string `json:"iccid" db:"iccid"`
	StarterKillStatus           *string `json:"starter_kill_status" db:"starter_kill_status"`
	Reason                      *string `json:"reason" db:"reason"`
	LastModifiedStarterKill     *string `json:"last_modified_starter_kill" db:"last_modified_starter_kill"`
	StarterKillHistoryStatus    *string `json:"starter_kill_history_status" db:"starter_kill_history_status"`
	StarterKillHistoryCreatedBy *string `json:"starter_kill_history_created_by" db:"starter_kill_history_created_by"`
}

type GetVehHistoryCountReq struct {
	OwnerID          string `json:"owner_id"`
	DealerID         string `json:"dealer_id"`
	RealmID          string `json:"realm_id"`
	IsOrgAdmin       string `json:"isorgadmin"`
	UserID           string `json:"user_id"`
	Search           string `json:"search"`
	ActivationStatus string `json:"activation_status"`
}

type VehHistoryCountTmpResult struct {
	ActivationStatus string `json:"activation_status" db:"activation_status"`
	Count            int64  `json:"count" db:"count"`
}

type VehHistoryCountResult struct {
	VehicleActive   int64 `json:"vehicle_active" db:"vehicle_active"`
	VehicleInactive int64 `json:"vehicle_inactive" db:"vehicle_inactive"`
	Euro4           int64 `json:"euro4" db:"euro4"`
	Euro2           int64 `json:"euro2" db:"euro2"`
}

type GetVehHistoryCountByOwnerIDReq struct {
	CountPage        string `json:"count_perpage" valid:"required"`
	Page             string `json:"page" valid:"required"`
	OwnerID          string `json:"owner_id"`
	DealerID         string `json:"dealer_id"`
	ActivationStatus string `json:"activation_status"`
}

type VehHistoryCountByOwnerIDListMetaResult struct {
	TotalData        int64                            `json:"total_data"`
	TotalDataPerpage int                              `json:"total_data_perpage"`
	From             int                              `json:"from"`
	To               int                              `json:"to"`
	TotalPage        int                              `json:"total_page"`
	Data             []VehHistoryCountByOwnerIDResult `json:"data"`
}

type VehHistoryCountByOwnerIDResult struct {
	OwnerID string `json:"owner_id" db:"owner_id"`
	Count   int64  `json:"count" db:"count"`
}

type GetVehHistoryByDealerIdReq struct {
	DealerID string `json:"dealer_id"`
}

type VehHistoryOwnerIdResult struct {
	OwnerID string `json:"owner_id" db:"owner_id"`
}

type GetVehHistoryByChassisReq struct {
	ChassisNumber string `json:"chassis_number" db:"chassis_number" valid:"required"`
	DealerID      string `json:"dealer_id"`
	RealmID       string `json:"realm_id"`
}

type VehHistoryByChassisReq struct {
	ChassisNumber string `json:"chassis_number" db:"chassis_number" valid:"required"`
}

type VehHistoryByImeiReq struct {
	Imei string `json:"imei" db:"imei" valid:"required"`
}

type GetCountVehHistoryByMonthReq struct {
	OwnerID    string `json:"owner_id"`
	RealmID    string `json:"realm_id"`
	IsOrgAdmin string `json:"isorgadmin"`
	UserID     string `json:"user_id"`
	From       string `json:"from"`
	To         string `json:"to"`
}

type CountVehHistoryByMonthResult struct {
	Month        string `json:"month" db:"months"`
	CountVehicle string `json:"count_vehicle" db:"count_vehicle"`
}

type CountVehActiveInactiveResult struct {
	DealerID        string `json:"dealer_id" db:"dealer_id"`
	VehicleActive   int    `json:"vehicle_active" db:"vehicle_active"`
	VehicleInactive int    `json:"vehicle_inactive" db:"vehicle_inactive"`
}

type ChangeOwnershipVehHistoryReq struct {
	ChassisNumber    string `json:"chassis_number" db:"chassis_number" valid:"required,stringlength(1|17)"`
	OwnerID          string `json:"owner_id" db:"owner_id" valid:"required"`
	ProofOfOwnership string `json:"proof_of_ownership" db:"proof_of_ownership"`
}

type AddEditVehHistoryReq struct {
	ChassisNumber      string `json:"chassis_number" db:"chassis_number" valid:"required"`
	Imei               string `json:"imei" db:"imei"`
	VehicleName        string `json:"vehicle_name" db:"vehicle_name" valid:"required"`
	VehicleNumber      string `json:"vehicle_number" db:"vehicle_number" valid:"required"`
	OwnerID            string `json:"owner_id" db:"owner_id"`
	Activation_date    string `json:"activation_date" db:"activation_date"`
	ActivationStatus   string `json:"activation_status" db:"activation_status"`
	SubscriptionPeriod string `json:"subscription_period" db:"subcription_period"`
	StnkDate           string `json:"stnk_date" db:"stnk_date"`
	KirDate            string `json:"kir_date" db:"kir_date"`
	ProofOfOwnership   string `json:"proof_of_ownership" db:"proof_of_ownership"`
	Weight             string `json:"weight" db:"weight"`
	Length             string `json:"length" db:"length"`
	Width              string `json:"width" db:"width"`
	Height             string `json:"height" db:"height"`
	RearBodyTypeID     string `json:"rear_body_type_id" db:"rear_body_type_id"`
	LoadTypeID         string `json:"load_type_id" db:"load_type_id"`
	BusinessSectorID   string `json:"business_sector_id" db:"business_sector_id"`
	Area               string `json:"area" db:"area"`
	Odometer           string `json:"odometer" db:"odometer"`
	GsmNumber          string `json:"gsm_number" db:"gsm_number"`
}

type EditVehHistoryReq struct {
	ChassisNumber  string `json:"chassis_number" db:"chassis_number" valid:"required,stringlength(1|17)"`
	VehicleName    string `json:"vehicle_name" db:"vehicle_name" valid:"stringlength(1|40)"`
	VehicleNumber  string `json:"vehicle_number" db:"vehicle_number" valid:"required,stringlength(1|20)"`
	StnkDate       string `json:"stnk_date" db:"stnk_date"`
	KirDate        string `json:"kir_date" db:"kir_date"`
	Weight         string `json:"weight" db:"weight"`
	Length         string `json:"length" db:"length"`
	Width          string `json:"width" db:"width"`
	Height         string `json:"height" db:"height"`
	RearBodyTypeID string `json:"rear_body_type_id" db:"rear_body_type_id"`
	LoadTypeID     string `json:"load_type_id" db:"load_type_id"`
}

type VehHistoryChangeActStatusReq struct {
	ChassisNumber    string `json:"chassis_number" db:"chassis_number" valid:"required"`
	ActivationStatus string `json:"activation_status" db:"activation_status" valid:"required"`
}

type VehHistoryMessageResult struct {
	ChassisNumber string `json:"chassis_number" db:"chassis_number"`
	VehicleNumber string `json:"vehicle_number" db:"vehicle_number"`
	Error         bool   `json:"error"`
	Message       string `json:"message"`
}
