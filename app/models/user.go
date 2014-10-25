package models

import (
	"time"
)

/**
 * ユーザーモデル
 */
type User struct {
	Id        int64  `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	CreatedAt int64  `json:"createdAt,omitempty"`
	UpdatedAt int64  `json:"updatedAt,omitempty"`
	DeletedAt int64  `json:"deletedAt,omitempty"`
}

func (u *User) BeforeCreate() (err error) {
	u.CreatedAt = time.Now().Unix()
	u.UpdatedAt = time.Now().Unix()
	return
}

func (u *User) BeforeUpdate() (err error) {
	u.UpdatedAt = time.Now().Unix()
	return
}

func (u *User) BeforeDelete() (err error) {
	u.DeletedAt = time.Now().Unix()
	return
}
