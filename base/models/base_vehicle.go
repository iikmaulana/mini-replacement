package models

type VehMessageResult struct {
	ChassisNumber string `json:"chassis_number" db:"chassis_number"`
	Error         bool   `json:"error"`
	Message       string `json:"message"`
}

type AddEditVehReq struct {
	ChassisNumber   string  `json:"chassis_number" db:"chassis_number" valid:"required"`
	Imei            string  `json:"imei" db:"imei" valid:"required"`
	EngineNumber    string  `json:"engine_number" db:"engine_number" valid:"required"`
	VehicleTypeId   string  `json:"vehicle_type_id" db:"vehicle_type_id" valid:"required"`
	DeviceStatus    string  `json:"device_status" db:"device_status"`
	ProductionDate  string  `json:"production_date" db:"production_date"`
	DeliverPosition string  `json:"deliver_position" db:"deliver_position"`
	EuroStd         *string `json:"euro_std" db:"euro_std"`
	Status          string  `json:"status" db:"status" valid:"required"`
}

type VehicleByChassisReq struct {
	ChassisNumber string `json:"chassis_number" db:"chassis_number" valid:"required"`
}

type GetVehicleReq struct {
	CountPage     string `json:"count_perpage" valid:"required"`
	Page          string `json:"page" valid:"required"`
	Search        string `json:"search"`
	StatusVehicle string `json:"status_vehicle"`
	Header        string `json:"header"`
}

type VehicleListMetaResult struct {
	TotalData        int64       `json:"total_data"`
	TotalDataPerpage int         `json:"total_data_perpage"`
	From             int         `json:"from"`
	To               int         `json:"to"`
	TotalPage        int         `json:"total_page"`
	Data             []VehResult `json:"data"`
}

type VehResult struct {
	ChassisNumber       string  `json:"chassis_number" db:"chassis_number"`
	EngineNumber        *string `json:"engine_number" db:"engine_number"`
	Imei                *string `json:"imei" db:"imei"`
	VehicleTypeId       string  `json:"vehicle_type_id" db:"vehicle_type_id"`
	VehicleTypeName     string  `json:"vehicle_type_name" db:"vehicle_type_name"`
	VehicleModelId      string  `json:"vehicle_model_id" db:"vehicle_model_id"`
	VehicleModelName    string  `json:"vehicle_model_name" db:"vehicle_model_name"`
	VehicleCategoryId   string  `json:"vehicle_category_id" db:"vehicle_category_id"`
	VehicleCategoryName string  `json:"vehicle_category_name" db:"vehicle_category_name"`
	DeviceStatus        *string `json:"device_status" db:"device_status"`
	ProductionDate      *string `json:"production_date" db:"production_date"`
	DeliverPosition     string  `json:"deliver_position" db:"deliver_position"`
	CreatedAt           string  `json:"created_at" db:"created_at"`
	UpdatedAt           *string `json:"updated_at" db:"updated_at"`
}

type VehicleV3 struct {
	ChassisNumber  string  `json:"chassis_number" db:"chassis_number"`
	Imei           *string `json:"imei" db:"imei"`
	VehicleName    *string `json:"vehicle_name" db:"vehicle_name"`
	VehicleNumber  *string `json:"vehicle_number" db:"vehicle_number"`
	MemberID       *string `json:"member_id" db:"member_id"`
	IsActive       *string `json:"is_active" db:"is_active"`
	ActivationDate *string `json:"activation_date" db:"activation_date"`
	DeviceStatus   *string `json:"device_status" db:"device_status"`
	PostDealerID   *string `json:"post_dealer_id" db:"post_dealer_id"`
	ActvDealerID   *string `json:"actv_dealer_id" db:"actv_dealer_id"`
	GsmNumber      *string `json:"gsm_number" db:"gsm_number"`
	EngineNumber   *string `json:"engine_number" db:"engine_number"`
	CreatedAt      *string `json:"created_at" db:"created_at"`
}

type VehicleReq struct {
	DealerID             string `json:"dealer_id"`
	RealmID              string `json:"realm_id"`
	OwnerID              string `json:"owner_id"`
	DeviceStatus         string `json:"device_status"`
	ChassisNumber        string `json:"chassis_number"`
	Imei                 string `json:"imei"`
	GsmNumber            string `json:"gsm_number"`
	VehicleName          string `json:"vehicle_name"`
	VehicleTypeID        string `json:"vehicle_type_id"`
	BusinessSectorID     string `json:"business_sector_id"`
	Area                 string `json:"area"`
	Odometer             string `json:"odometer"`
	RearBodyTypeID       string `json:"rear_body_type_id"`
	LeasingID            string `json:"leasing_id"`
	ProspectingLeasingID string `json:"prospecting_leasing_id"`
	PaymentMethod        string `json:"payment_method"`
}
