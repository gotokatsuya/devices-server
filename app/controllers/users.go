package controllers

import "github.com/revel/revel"
import m "devices-server/app/models"

type Users struct {
	GormController
}

/*
	Userを作成
 	@param username:ユーザーネーム
 	return data{sucess, user}
*/
func (c Users) Create(username string) revel.Result {
	data := struct {
		Success bool   `json:"success"`
		User    m.User `json:"user"`
	}{
		Success: false,
		User:    m.User{},
	}

	var users []m.User
	c.Txn.Find(&users, "name = ?", username)
	if len(users) == 0 {
		user := m.User{Name: username}
		c.Txn.NewRecord(user)
		c.Txn.Create(&user)
		data.User = user
		data.Success = true
	}
	return c.RenderJson(data)
}

/*
	Userのリストを取得
 	return data{sucess, users}
*/
func (c Users) GetList() revel.Result {
	data := struct {
		Success bool     `json:"success"`
		Users   []m.User `json:"users"`
	}{
		Success: false,
		Users:   []m.User{},
	}

	var users []m.User
	c.Txn.Find(&users)
	if len(users) != 0 {
		data.Users = users
		data.Success = true
	}
	return c.RenderJson(data)
}
