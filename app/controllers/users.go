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
	Userを更新
	@param user_id:ID
 	@param username:ユーザーネーム
 	return data{sucess, user}
*/
func (c Users) Update(user_id int64, username string) revel.Result {
	data := struct {
		Success bool   `json:"success"`
		User    m.User `json:"user"`
	}{
		Success: false,
		User:    m.User{},
	}

	var users []m.User
	c.Txn.Find(&users, "id = ?", user_id)
	if len(users) != 0 {
		user := users[0]
		user.Name = username
		c.Txn.Save(&user)
		data.User = user
		data.Success = true
	}
	return c.RenderJson(data)
}

/*
	Userのリストを取得
 	return data{sucess, users}
*/
func (c Users) List() revel.Result {
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
