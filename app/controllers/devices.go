package controllers

import (
	m "devices-server/app/models"
	"github.com/revel/revel"
)

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
	dateOfRelease int64,
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
	Deviceを更新
 	@param device_id:ID
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
func (c Devices) Update(device_id int64,
	name string,
	manufacturer string,
	carrier string,
	os string,
	size string,
	resolution string,
	memory string,
	dateOfRelease int64,
	other string) revel.Result {
	data := struct {
		Success bool     `json:"success"`
		Device  m.Device `json:"device"`
	}{
		Success: false,
		Device:  m.Device{},
	}

	var devices []m.Device
	c.Txn.Find(&devices, "id = ?", device_id)
	if len(devices) != 0 {
		device := devices[0]
		device.Name = name
		device.Manufacturer = manufacturer
		device.Carrier = carrier
		device.Os = os
		device.Size = size
		device.Resolution = resolution
		device.Memory = memory
		device.DateOfRelease = dateOfRelease
		device.Other = other
		c.Txn.Save(&device)
		data.Device = device
		data.Success = true
	}
	return c.RenderJson(data)
}

/*
	Deviceのリストを取得
 	return data{sucess, devices}
*/
func (c Devices) List() revel.Result {
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
		for i := 0; i < len(devices); i++ {
			device := devices[i]
			device = m.FindUser(c.Txn, device)
			device = m.FindDeviceStates(c.Txn, device)
			devices[i] = device
		}
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
			user := users[0]
			device.UserId = user.Id
			device.User = user
			device.State = true
			device.DeviceStates = c.FindAfterCreateDeviceState(user, device, true)
			c.Txn.Save(&device)

			data.Device = device
			data.Success = true
		}
	}
	return c.RenderJson(data)
}

/*
	ユーザーがDeviceを返却する
	@param userId:ユーザ-ID
 	@param deviceId:端末ID
 	return data{sucess, device}
*/
func (c Devices) Return(user_id int64, device_id int64) revel.Result {
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
			user := users[0]
			device.UserId = user.Id
			device.User = user
			device.State = false
			device.DeviceStates = c.FindAfterCreateDeviceState(user, device, false)
			c.Txn.Save(&device)

			data.Device = device
			data.Success = true
		}
	}
	return c.RenderJson(data)
}

/*
	履歴を追加する
	@param deviceStates:端末の貸し出し履歴
 	@param user:ユーザ-
 	@param device_id:端末ID
 	return deviceStates
*/
func (c Devices) FindAfterCreateDeviceState(user m.User, device m.Device, state bool) []m.DeviceState {
	deviceState := m.DeviceState{
		State:    state,
		DeviceId: device.Id,
		UserId:   user.Id,
	}
	c.Txn.NewRecord(deviceState)
	c.Txn.Create(&deviceState)

	var device_states []m.DeviceState
	c.Txn.Model(&device).Related(&device_states)

	for i := 0; i < len(device_states); i++ {
		device_state := device_states[i]
		var user m.User
		c.Txn.Model(&device_state).Related(&user)
		device_state.User = user
		device_states[i] = device_state
	}

	return device_states
}
