package models

type RpcListOrganizationResult struct {
	TotalData        int64                               `json:"total_data"`
	TotalDataPerpage int                                 `json:"total_data_perpage"`
	From             int                                 `json:"from"`
	To               int                                 `json:"to"`
	TotalPage        int                                 `json:"total_page"`
	Data             []RpcListOrganizationCustomerResult `json:"data"`
}

type RpcListOrganizationCustomerResult struct {
	ID                  string   `db:"id" json:"id"`
	RealmID             string   `db:"realm_id" json:"realm_id"`
	RealmName           string   `db:"realm_name" json:"realm_name"`
	BusinessType        *string  `db:"business_type" json:"business_type"`
	Code                *string  `db:"code" json:"code"`
	Name                *string  `db:"name" json:"name"`
	Address             *string  `db:"address" json:"address"`
	CityID              *int     `db:"city_id" json:"city_id"`
	ProvinceID          *int     `db:"province_id" json:"province_id"`
	Email               *string  `db:"email" json:"email"`
	Telp                *string  `db:"telp" json:"phone"`
	Fax                 *string  `db:"fax" json:"fax"`
	AddressPosition     *string  `db:"address_position" json:"address_position"`
	PreferensiDealer    *string  `db:"preferensi_dealer" json:"preferensi_dealer"`
	TypeDocument        *string  `db:"type_document" json:"type_document"`
	NomorDocument       *string  `db:"nomor_document" json:"nomor_document"`
	BusinessSector      *string  `db:"business_sector" json:"business_sector"`
	CreatedAt           *string  `db:"created_at" json:"created_at"`
	UpdatedAt           *string  `db:"updated_at" json:"updated_at"`
	DeletedAt           *string  `db:"deleted_at" json:"deleted_at"`
	ProfileID           *string  `db:"profile_id" json:"profile_id"`
	ContactName         *string  `db:"contact_name" json:"contact_name"`
	ContactPhone        *string  `db:"contact_phone" json:"contact_phone"`
	ContactEmail        *string  `db:"contact_email" json:"contact_email"`
	UserId              *string  `db:"user_id" json:"user_id"`
	Username            *string  `db:"username" json:"username"`
	UserStatus          *int     `db:"user_status" json:"user_status"`
	ParentID            *string  `db:"parent_id" json:"parent_id"`
	DistrictId          *int     `db:"district_id" json:"district_id"`
	CustomerId          *string  `db:"customer_id" json:"customer_id"`
	OrgContactPerson    *string  `db:"org_contact_person" json:"org_contact_person"`
	OrgContactPersonJob *string  `db:"org_contact_person_job" json:"org_contact_person_job"`
	SubDistrictId       *int     `db:"sub_district_id" json:"sub_district_id"`
	PostCode            *int     `db:"post_code" json:"post_code"`
	AreaId              *string  `db:"area_id" json:"area_id"`
	AreaName            *string  `db:"area_name" json:"area_name"`
	Longitude           *float64 `json:"longitude"`
	Latitude            *float64 `json:"latitude"`
}

type CreateOrganizationRequest struct {
	UserId              string  `json:"user_id"`
	BusinessType        string  `json:"business_type"`
	RealmID             string  `json:"realm_id"`
	Code                string  `json:"code"`
	Name                string  `json:"name"`
	Address             string  `json:"address"`
	CityID              int     `json:"city_id"`
	ProvinceID          int     `json:"province_id"`
	Email               string  `json:"email"`
	Telp                string  `json:"telp"`
	Fax                 string  `json:"fax"`
	AddressPosition     string  `json:"address_position"`
	PreferensiDealer    string  `json:"preferensi_dealer"`
	TypeDocument        string  `json:"type_document"`
	NomorDocument       string  `json:"nomor_document"`
	BusinessSector      string  `json:"business_sector"`
	ParentID            *string `json:"parent_id"`
	ContactName         string  `json:"contact_name"`
	SuperUsername       string  `json:"super_username"`
	SuperAdminEmail     string  `json:"super_admin_email"`
	SuperAdminPhone     string  `json:"super_admin_phone"`
	AccessType          *int    `json:"access_type"`
	Longitude           float64 `json:"longitude"`
	Latitude            float64 `json:"latitude"`
	DistrictId          int     `json:"district_id"`
	AreaId              *string `db:"area_id" json:"area_id"`
	AreaName            *string `db:"area_name" json:"area_name"`
	CustomerId          string  `db:"customer_id" json:"customer_id"`
	OrgContactPerson    string  `db:"org_contact_person" json:"org_contact_person"`
	OrgContactPersonJob string  `db:"org_contact_person_job" json:"org_contact_person_job"`
	SubDistrictId       int     `db:"sub_district_id" json:"sub_district_id"`
	PostCode            int     `db:"post_code" json:"post_code"`
}

type UpdateOrganizationRequest struct {
	ID                  string  `db:"id" json:"id"`
	UserId              string  `db:"user_id" json:"user_id"`
	BusinessType        string  `db:"business_type" json:"business_type"`
	RealmID             string  `db:"realm_id" json:"realm_id"`
	Code                string  `db:"code" json:"code"`
	Name                string  `db:"name" json:"name"`
	Address             string  `db:"address" json:"address"`
	CityID              int     `db:"city_id" json:"city_id"`
	ProvinceID          int     `db:"province_id" json:"province_id"`
	Email               string  `db:"email" json:"email"`
	Telp                string  `db:"telp" json:"phone"`
	Fax                 string  `db:"fax" json:"fax"`
	AddressPosition     string  `db:"address_position" json:"address_position"`
	PreferensiDealer    string  `db:"preferensi_dealer" json:"preferensi_dealer"`
	TypeDocument        string  `db:"type_document" json:"type_document"`
	NomorDocument       string  `db:"nomor_document" json:"nomor_document"`
	BusinessSector      string  `db:"business_sector" json:"business_sector"`
	ContactName         string  `json:"contact_name"`
	SuperUsername       string  `json:"super_username"`
	SuperAdminEmail     string  `json:"super_admin_email"`
	SuperAdminPhone     string  `json:"super_admin_phone"`
	Longitude           float64 `json:"longitude"`
	Latitude            float64 `json:"latitude"`
	DistrictId          int     `db:"district_id" json:"district_id"`
	AreaId              *string `db:"area_id" json:"area_id"`
	AreaName            *string `db:"area_name" json:"area_name"`
	CustomerId          string  `db:"customer_id" json:"customer_id"`
	OrgContactPerson    string  `db:"org_contact_person" json:"org_contact_person"`
	OrgContactPersonJob string  `db:"org_contact_person_job" json:"org_contact_person_job"`
	SubDistrictId       int     `db:"sub_district_id" json:"sub_district_id"`
	PostCode            int     `db:"post_code" json:"post_code"`
}

type OrganizationResult struct {
	ID               string   `db:"id" json:"id"`
	UserId           string   `db:"user_id" json:"user_id"`
	BusinessType     string   `db:"business_type" json:"business_type"`
	RealmID          string   `db:"realm_id" json:"realm_id"`
	RealmName        string   `db:"realm_name" json:"realm_name"`
	Code             *string  `db:"code" json:"code"`
	Name             *string  `db:"name" json:"name"`
	Address          *string  `db:"address" json:"address"`
	CityID           *int     `db:"city_id" json:"city_id"`
	ProvinceID       *int     `db:"province_id" json:"province_id"`
	Email            *string  `db:"email" json:"email"`
	Telp             *string  `db:"telp" json:"phone"`
	Fax              *string  `db:"fax" json:"fax"`
	AddressPosition  *string  `db:"address_position" json:"address_position"`
	PreferensiDealer *string  `db:"preferensi_dealer" json:"preferensi_dealer"`
	TypeDocument     *string  `db:"type_document" json:"type_document"`
	NomorDocument    *string  `db:"nomor_document" json:"nomor_document"`
	BusinessSector   *string  `db:"business_sector" json:"business_sector"`
	CreatedAt        *string  `db:"created_at" json:"created_at"`
	UpdatedAt        *string  `db:"updated_at" json:"updated_at"`
	DeletedAt        *string  `db:"deleted_at" json:"deleted_at"`
	ProfileID        *string  `db:"profile_id" json:"profile_id"`
	ContactName      *string  `db:"contact_name" json:"contact_name"`
	SuperUsername    *string  `db:"super_username" json:"super_username"`
	ParentID         *string  `db:"parent_id" json:"parent_id"`
	Longitude        *float64 `json:"longitude"`
	Latitude         *float64 `json:"latitude"`
	DistrictId       *int     `db:"district_id" json:"district_id"`
	AreaId           *string  `db:"area_id" json:"area_id"`
	AreaName         *string  `db:"area_name" json:"area_name"`
}

type ListOrganizationResult struct {
	TotalData        int64                `json:"total_data"`
	TotalDataPerpage int                  `json:"total_data_perpage"`
	From             int                  `json:"from"`
	To               int                  `json:"to"`
	TotalPage        int                  `json:"total_page"`
	Data             []OrganizationResult `json:"data"`
}

type ListOrganizationDealerResult struct {
	TotalData        int64                      `json:"total_data"`
	TotalDataPerpage int                        `json:"total_data_perpage"`
	From             int                        `json:"from"`
	To               int                        `json:"to"`
	TotalPage        int                        `json:"total_page"`
	Data             []OrganizationDealerResult `json:"data"`
}

type OrganizationDealerResult struct {
	Id         string `json:"id"`
	DealerName string `json:"name"`
	Telp       string `json:"telp"`
	Address    string `json:"address"`
	Latitude   string `json:"latitude"`
	Longitude  string `json:"longitude"`
}

type ListAdvanceSearchOrganizationCustomerResult struct {
	TotalData        int64                                   `json:"total_data"`
	TotalDataPerpage int                                     `json:"total_data_perpage"`
	From             int                                     `json:"from"`
	To               int                                     `json:"to"`
	TotalPage        int                                     `json:"total_page"`
	Data             []AdvanceSearchListOrganizationCustomer `json:"data"`
}

type AdvanceSearchListOrganizationCustomer struct {
	ID         string  `json:"org_id"`
	Username   *string `json:"username"`
	Name       *string `json:"name"`
	Address    *string `json:"address"`
	CreatedAt  *string `json:"created_at"`
	RegisterAt *string `json:"register_at"`
	ParentId   *string `json:"parent_id"`
	UserStatus *int    `json:"user_status"`
}

type ListOrganizationCustomerResult struct {
	TotalData        int64                      `json:"total_data"`
	TotalDataPerpage int                        `json:"total_data_perpage"`
	From             int                        `json:"from"`
	To               int                        `json:"to"`
	TotalPage        int                        `json:"total_page"`
	Data             []ListOrganizationCustomer `json:"data"`
}

type ListOrganizationCustomer struct {
	ID           string  `json:"org_id"`
	Username     *string `json:"username"`
	OrgName      *string `json:"name"`
	OrgEmail     *string `json:"email"`
	OrgPhone     *string `json:"phone"`
	UserStatus   *int    `json:"user_status"`
	VehicleCount *int64  `json:"vehicle_count"`
	OrgDealer    *string `json:"dealer"`
	LastContact  *string `json:"last_contact"`
}

type ViewOrganizationCustomerResult struct {
	ID                   string                 `db:"id" json:"id"`
	RegisteredDealer     string                 `db:"registered_dealer" json:"registered_dealer"`
	CustomerID           *string                `db:"customer_id" json:"customer_id"`
	CompanyName          *string                `db:"company_name" json:"company_name"`
	CompanyType          *string                `db:"company_type" json:"company_type"`
	BusinessSector       *string                `db:"business_sector" json:"business_sector"`
	NomorDocument        *string                `db:"nomor_document" json:"nomor_document"`
	TypeDocument         *string                `db:"type_document" json:"type_document"`
	CreatedAt            *string                `db:"created_at" json:"created_at"`
	UpdatedAt            *string                `db:"updated_at" json:"updated_at"`
	UserStatus           *int                   `json:"user_status"`
	CustomerAccount      *CustomerAccountResult `json:"customer_account"`
	CustomerDetailResult *CustomerDetailResult  `json:"customer_detail"`
}

type CustomerDetailResult struct {
	ProvinceID          *string `db:"province" json:"province"`
	CityID              *string `db:"city" json:"city"`
	DistrictId          *string `db:"district" json:"district"`
	SubDistrictId       *int    `db:"sub_district_id" json:"sub_district_id"`
	PostCode            *int    `db:"post_code" json:"post_code"`
	Address             *string `db:"address" json:"address"`
	Email               *string `db:"email" json:"email"`
	Telp                *string `db:"telp" json:"phone"`
	OrgContactPerson    *string `db:"org_contact_person" json:"org_contact_person"`
	OrgContactPersonJob *string `db:"org_contact_person_job" json:"org_contact_person_job"`
}

type CustomerAccountResult struct {
	Username *string `db:"username" json:"username"`
	Email    *string `db:"email" json:"email"`
	Phone    *string `db:"phone" json:"phone"`
}

type ListOrganizationCustomerInActiveResult struct {
	TotalData        int64                              `json:"total_data"`
	TotalDataPerpage int                                `json:"total_data_perpage"`
	From             int                                `json:"from"`
	To               int                                `json:"to"`
	TotalPage        int                                `json:"total_page"`
	Data             []ListOrganizationCustomerInActive `json:"data"`
}

type ListOrganizationCustomerInActive struct {
	ID          string  `json:"org_id"`
	OrgName     *string `json:"name"`
	OrgPhone    *string `json:"phone"`
	OrgEmail    *string `json:"email"`
	OrgDealer   string  `json:"dealer"`
	LastContact *string `json:"last_contact"`
	UserStatus  *int    `json:"user_status"`
}

type AddEditOrganizationActivationLogReq struct {
	OrgId       string `json:"org_id" db:"org_id" valid:"required"`
	ContactByID string `json:"contact_by" db:"contact_by" valid:"required"`
	Message     string `json:"message" db:"message" valid:"required"`
	Via         string `json:"via" db:"via" valid:"required"`
	CreatedAt   string `json:"created_at" db:"created_at" valid:"required"`
	Status      int    `json:"status" db:"status"`
}

type ListOrganizationLeasingCustomerResult struct {
	TotalData        int64                             `json:"total_data"`
	TotalDataPerpage int                               `json:"total_data_perpage"`
	From             int                               `json:"from"`
	To               int                               `json:"to"`
	TotalPage        int                               `json:"total_page"`
	Data             []ListOrganizationLeasingCustomer `json:"data"`
}

type ListOrganizationLeasingCustomer struct {
	ID      string  `json:"org_id"`
	OrgName *string `json:"name"`
}

type ListOrganizationKemoodiResult struct {
	TotalData        int64                       `json:"total_data"`
	TotalDataPerpage int                         `json:"total_data_perpage"`
	From             int                         `json:"from"`
	To               int                         `json:"to"`
	TotalPage        int                         `json:"total_page"`
	Data             []OrganizationKemoodiResult `json:"data"`
}
type OrganizationKemoodiResult struct {
	ID   string  `db:"id" json:"id"`
	Name *string `db:"name" json:"name"`
}
