package query

const (
	CreateMtMember = `INSERT INTO runner_app.mt_member (
							member_id,
							member_name,
							company,
							email,
							phone,
							dealer_id,
							member_code,
							company_type,
							date_registration,
							status_engage
							) VALUES (%d,'%s','%s','%s','%s','%s','%s','%s',now(),'not_active')
						`

	UpdateMtMember = `UPDATE runner_app.mt_member 
							SET member_name = '%s',
							company = '%s',
							email = '%s',
							phone = '%s',
							member_code = '%s',
							company_type = '%s' 
						WHERE
					  		member_id = %d and email = '%s'
						`

	CreateAuthRunner = `INSERT INTO runner_app.auth_runner (
							user_id,
							username,
							password,
							user_fullname,
							email,
							role_id,
							member_id,
							dealer_id,
							phone,
							creation_date
							) VALUES (%d,'%s','%s','%s','%s','%s',%s,%s,'%s',now())
						`

	UpdateAuthRunner = `UPDATE runner_app.auth_runner
							SET username = '%s',
							password = '%s',
							user_fullname = '%s',
							email = '%s',
							role_id = '%s',
							member_id = %s,
							dealer_id = %s,
							phone = '%s'
						WHERE
					  		username = '%s'
						`
	CreatePrivacyPolicy = `INSERT INTO runner_app.log_privacy_policy (id,policy_id,user_id,"status",accepted_date,created_at) VALUES ('%s',7,'%s',0,NULL,now())`
	CheckPrivacyPolicy  = `select exists ( select id from runner_app.log_privacy_policy lp where user_id = '%s' limit 1) as value`
	GetUserIdByUsername = `select user_id from runner_app.auth_runner where username = '%s' limit 1`
	DeletePrivacyPolicy = `delete from runner_app.log_privacy_policy where user_id = '%s'`

	GetMemberIdByEmailAndMemberName = `select member_id from runner_app.mt_member where email = '%s' and member_name = '%s' limit 1`

	UpdateAuthRunnerActivationCustomer = `UPDATE public.auth_runner (
												SET username = '%s',
												password = '%s',
												user_fullname = '%s',
												email = '%s',
												role_id = '%s',
												member_id = '%s',
												dealer_id = '%s',
												phone = '%s'
											WHERE
												username = '%s'
											`

	GetMemberIdByOrganizationId = `select member_id from um_runner.public.member_tmp where organization_id = '%s'`
	GetDealerIdByOrganizationId = `select dealer_id from um_runner.public.member_tmp where organization_id = '%s'`

	ChangePassword = `UPDATE runner_app.auth_runner
						SET password = '%s'
					WHERE
						username = '%s'
					`

	CreateMemberTmp = `insert into um_runner.member_tmp (organization_id,member_id,dealer_id) values ($1,$2,$3)`

	GetUserByUsername = `SELECT u.id, u.username, u.password, u.access_type, u.user_status, u.user_type, u.realm_id, u.profile_id, up.name as profile_name, up.phone as profile_phone, up.email as profile_email, u.organization_id, org.name as organization_name, org.telp as organization_phone, org.email as organization_email, coalesce(u.created_at::TIMESTAMP(0)::STRING, '') as created_at, coalesce(u.updated_at::TIMESTAMP(0)::STRING, '') as updated_at, coalesce(u.deleted_at::TIMESTAMP(0)::STRING, '') as deleted_at 
					 FROM um_runner.public.user u
					 INNER JOIN um_runner.public.user_profile up on up.id = u.profile_id 
					 INNER JOIN um_runner.public.organization org on org.id = u.organization_id
					where u.username = $1 and u.deleted_at IS NULL`

	GetUserByUserId = `SELECT u.id, u.username, u.password, u.access_type, u.user_status, u.user_type, u.realm_id, u.profile_id, up.name as profile_name, up.phone as profile_phone, up.email as profile_email, u.organization_id, org.name as organization_name, org.telp as organization_phone, org.email as organization_email, coalesce(u.created_at::TIMESTAMP(0)::STRING, '') as created_at, coalesce(u.updated_at::TIMESTAMP(0)::STRING, '') as updated_at, coalesce(u.deleted_at::TIMESTAMP(0)::STRING, '') as deleted_at 
					 FROM um_runner.public.user u
					 INNER JOIN um_runner.public.user_profile up on up.id = u.profile_id 
					 INNER JOIN um_runner.public.organization org on org.id = u.organization_id
					where u.id = $1 and u.deleted_at IS NULL`

	GetUserByUserIdDeleted = `SELECT u.id, u.username, u.password, u.access_type, u.user_status, u.user_type, u.realm_id, u.profile_id, up.name as profile_name, up.phone as profile_phone, up.email as profile_email, u.organization_id, org.name as organization_name, org.telp as organization_phone, org.email as organization_email, coalesce(u.created_at::TIMESTAMP(0)::STRING, '') as created_at, coalesce(u.updated_at::TIMESTAMP(0)::STRING, '') as updated_at, coalesce(u.deleted_at::TIMESTAMP(0)::STRING, '') as deleted_at 
					 FROM um_runner.public.user u
					 INNER JOIN um_runner.public.user_profile up on up.id = u.profile_id 
					 INNER JOIN um_runner.public.organization org on org.id = u.organization_id
					where u.id = $1`

	DeleteAuthRunner = `delete from runner_app.auth_runner where user_id = '%s'`
)
