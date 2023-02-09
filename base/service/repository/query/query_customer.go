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
					  		member_id = '%s'
						`

	CreateAuthRunner = `INSERT INTO runner_app.auth_runner (
							username,
							password,
							user_fullname,
							email,
							role_id,
							member_id,
							dealer_id,
							phone,
							creation_date
							) VALUES ('%s','%s','%s','%s','%s','%s','%s','%s',now())
						`

	UpdateAuthRunner = `UPDATE runner_app.auth_runner
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
	CreatePrivacyPolicy = `INSERT INTO runner_app.log_privacy_policy (id,policy_id,user_id,"status",accepted_date,created_at) VALUES ('%s',7,'%s',0,NULL,now())`
	CheckPrivacyPolicy  = `select exists ( select id from runner_app.log_privacy_policy lp where user_id = '%s' limit 1) as value`
	GetUserIdByUsername = `select user_id from runner_app.auth_runner where username = '%s' limit 1`

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
)
