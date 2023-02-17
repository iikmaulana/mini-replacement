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

	GetVehicle = `select 
					vh.chassis_number, 
					vh.imei, 
					vh.vehicle_name, 
					vh.vehicle_number, 
					mm.member_id as member_id, 
					case 
						when vh.activation_status = 1 then 0
						else 1
					end is_active,
					vh.activation_date, 
					vv.device_status, 
					md.dealer_id as post_dealer_id, 
					md.dealer_id as actv_dealer_id, 
					vh.gsm_number, 
					vv.engine_number,
					vh.created_at
				from db_vehicle.veh_vehicle_history vh
				left join db_vehicle.veh_vehicle vv on vh.chassis_number = vv.chassis_number
				left join db_vehicle.veh_vehicle_dealer vvd on vh.chassis_number = vvd.chassis_number
				left join um_runner.member_tmp mm on mm.organization_id = vh.owner_id
				left join um_runner.member_tmp md on md.organization_id = vvd.dealer_id
				where vh.chassis_number in = '%s'
				and vh.activation_status in (1,2)
				order by vh.created_at asc`

	CreateMtVehicleV2 = `INSERT INTO runner_app.mt_vehicle (
							vehicle_id,
							chassis_number,
							imei,
							vehicle_name,
							vehicle_number,
							member_id,
							is_active,
							activation_date,
							vehicle_status,
							post_dealer_id,
							actv_dealer_id,
							gps_gsm,
							engine_no) 
						VALUES
							 (%d,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s)`

	UpdateMtVehcileV2 = `UPDATE runner_app.mt_vehicle
							SET 
							chassis_number = %s,
							imei = %s,
							vehicle_name = %s,
							vehicle_number = %s,
							member_id = %s,
							is_active = %s,
							activation_date = %s,
							vehicle_status = %s,
							post_dealer_id = %s,
							actv_dealer_id = %s,
							gps_gsm = %s,
							engine_no = %s
						WHERE chassis_number = %s AND member_id = %s`
)
