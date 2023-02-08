package query

const (
	CreateMtVehcile = `INSERT INTO public.mt_vehicle (
							chassis_number,
							imei,
							vehicle_name,
							vehicle_number,
							member_id,
							activation_date
							) VALUES ('%s','%s','%s,'%s','%s','%s')
						`

	UpdateMtVehcile = `UPDATE public.mt_vehicle (
							SET chassis_number = '%s',
							imei = '%s',
							vehicle_name = '%s',
							vehicle_number = '%s',
							member_id = '%s',
							activation_date = '%s'
						WHERE chassis_number = '%s' AND member_id = '%s'
						`
)
