package models

type CreateUserRequest struct {
	Username       string  `json:"username"`
	Password       string  `json:"password"`
	AccessType     int64   `json:"access_type"`
	UserStatus     int64   `json:"user_status"`
	UserType       *string `json:"user_type"`
	RealmId        *string `json:"realm_id"`
	ProfileId      *string `json:"profile_id"`
	OrganizationId *string `json:"organization_id"`
}

type UpdateUserRequest struct {
	ID             string `db:"id" json:"id"`
	Username       string `db:"username" json:"username,omitempty"`
	Password       string `db:"password" json:"password,omitempty"`
	AccessType     int64  `db:"access_type" json:"access_type,omitempty"`
	UserStatus     int64  `db:"user_status" json:"user_status,omitempty"`
	UserType       string `db:"user_type" json:"user_type,omitempty"`
	RealmId        string `db:"realm_id" json:"realm_id,omitempty"`
	ProfileId      string `db:"profile_id" json:"profile_id,omitempty"`
	OrganizationId string `db:"organization_id" json:"organization_id,omitempty"`
}

type UserResult struct {
	ID                string  `db:"id" json:"id"`
	Username          *string `db:"username" json:"username"`
	Password          *string `db:"password" json:"password"`
	AccessType        *int64  `db:"access_type" json:"access_type"`
	UserStatus        *int64  `db:"user_status" json:"user_status"`
	UserType          *string `db:"user_type" json:"user_type"`
	RealmId           *string `db:"realm_id" json:"realm_id"`
	ProfileId         *string `db:"profile_id" json:"profile_id"`
	ProfileName       *string `db:"profile_name" json:"profile_name"`
	ProfileEmail      *string `db:"profile_email" json:"profile_email"`
	ProfilePhone      *string `db:"profile_phone" json:"profile_phone"`
	OrganizationId    *string `db:"organization_id" json:"organization_id"`
	OrganizationName  *string `db:"organization_name" json:"organization_name"`
	OrganizationEmail *string `db:"organization_email" json:"organization_email"`
	OrganizationPhone *string `db:"organization_phone" json:"organization_phone"`
	CreatedAt         *string `db:"created_at" json:"created_at"`
	UpdatedAt         *string `db:"updated_at" json:"updated_at"`
	DeletedAt         *string `db:"deleted_at" json:"deleted_at"`
}

type ListUserResult struct {
	TotalData        int64        `json:"total_data"`
	TotalDataPerpage int          `json:"total_data_perpage"`
	From             int          `json:"from"`
	To               int          `json:"to"`
	TotalPage        int          `json:"total_page"`
	Data             []UserResult `json:"data"`
}

type ChangePasswordRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
