package controllers

import "github.com/revel/revel"
import m "devices-server/app/models"
import "time"

type Devices struct {
	GormController
}

/*
	Deviceを作成
 	@param name:機種名
 	@param manufacturer:メーカー
 	@param carrier:キャリア
 	@param os:OS
 	@param size:サイズ
 	@param resolution:解像度
 	@param memory:メモリ
 	@param dateOfRelease:発売日
 	@param other:その他
 	return data{sucess, device}
*/
func (c Devices) Create(name string,
	manufacturer string,
	carrier string,
	os string,
	size string,
	resolution string,
	memory string,
	dateOfRelease time.Time,
	other string) revel.Result {
	data := struct {
		Success bool     `json:"success"`
		Device  m.Device `json:"device"`
	}{
		Success: false,
		Device:  m.Device{},
	}

	var devices []m.Device
	c.Txn.Find(&devices, "name = ?", name)
	if len(devices) == 0 {
		device := m.Device{
			Name:          name,
			Manufacturer:  manufacturer,
			Carrier:       carrier,
			Os:            os,
			Size:          size,
			Resolution:    resolution,
			Memory:        memory,
			DateOfRelease: dateOfRelease,
			Other:         other,
		}
		c.Txn.NewRecord(device)
		c.Txn.Create(&device)
		data.Device = device
		data.Success = true
	}
	return c.RenderJson(data)
}

/*
	Deviceのリストを取得
 	return data{sucess, devices}
*/
func (c Devices) GetList() revel.Result {
	data := struct {
		Success bool       `json:"success"`
		Devices []m.Device `json:"devices"`
	}{
		Success: false,
		Devices: []m.Device{},
	}

	var devices []m.Device
	c.Txn.Find(&devices)
	if len(devices) != 0 {
		data.Devices = devices
		data.Success = true
	}
	return c.RenderJson(data)
}

/*
	Deviceを特定のユーザーに貸し出す
	@param userId:ユーザ-ID
 	@param deviceId:端末ID
 	return data{sucess, device}
*/
func (c Devices) Borrow(user_id int64, device_id int64) revel.Result {
	data := struct {
		Success bool     `json:"success"`
		Device  m.Device `json:"device"`
	}{
		Success: false,
		Device:  m.Device{},
	}

	var users []m.User
	c.Txn.Find(&users, "id = ?", user_id)
	if len(users) != 0 {
		var devices []m.Device
		c.Txn.Find(&devices, "id = ?", device_id)
		if len(devices) != 0 {
			device := devices[0]
			device.User = users[0]
			data.Device = device
			data.Success = true
		}
	}
	return c.RenderJson(data)
}
