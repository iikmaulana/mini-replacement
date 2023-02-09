package query

const (
	CreateMtVehcile = `INSERT INTO runner_app.mt_vehicle (
							vehicle_id,
							chassis_number,
							imei,
							vehicle_name,
							vehicle_number,
							member_id,
							activation_date
							) VALUES (%d,'%s','%s','%s','%s',%d,'%s')
						`

	UpdateMtVehcile = `UPDATE runner_app.mt_vehicle
							SET chassis_number = '%s',
							imei = '%s',
							vehicle_name = '%s',
							vehicle_number = '%s',
							member_id = %d,
							activation_date = '%s'
						WHERE chassis_number = '%s' AND member_id = %d
						`

	CreateVehiclGroup = `INSERT INTO runner_app.mt_vehicle_group (group_id, group_name,description,member_id,dealer_id) VALUES
	 					(%d,'%s','%s',%d,0)`

	UpdateVehiclGroup = `UPDATE runner_app.mt_vehicle_group 
						SET group_name,
						description,
						member_id,
						dealer_id WHERE group_id = %d`

	DeleteVehiclGroup = `DELETE FROM runner_app.mt_vehicle_group WHERE group_id = %d`

	CreateMSVehicleGroup = `INSERT INTO public.ms_vehicle_group (group_id,vehicle_id) VALUES
							(%d,%d);`

	UpdateMSVehicleGroup = `UPDATE public.ms_vehicle_group 
							SET group_id = %d,
							vehicle_id = %d,
							update_at = now() 
							WHERE group_id = %d and vehicle_id = %d`

	DeleteMSVehicleGroup = `DELETE FROM runner_app.ms_vehicle_group WHERE group_id = %d and vehicle_id = %d`

	CheckMtVehcile = `select exists ( select chassis_number from runner_app.mt_vehicle where chassis_number = '%s' limit 1) as value`
)
