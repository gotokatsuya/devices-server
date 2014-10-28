package models

import (
	"time"
)

/**
 * 端末モデル
 */
type Device struct {
	Id            int64         `json:"id,omitempty"`
	Name          string        `json:"name,omitempty"`
	Manufacturer  string        `json:"manufacturer,omitempty"`
	Carrier       string        `json:"carrier,omitempty"`
	Os            string        `json:"os,omitempty"`
	Size          string        `json:"size,omitempty"`
	Resolution    string        `json:"resolution,omitempty"`
	Memory        string        `json:"memory,omitempty"`
	DateOfRelease int64         `json:"dataOfRelease,omitempty"`
	Other         string        `json:"other,omitempty"`
	CreatedAt     int64         `json:"createdAt,omitempty"`
	UpdatedAt     int64         `json:"updatedAt,omitempty"`
	DeletedAt     int64         `json:"deletedAt,omitempty"`
	User          User          `json:"user,omitempty"`         // One-To-One relationship (has one)
	UserId        int64         `json:"user_id,omitempty"`      // Foreign key of User
	DeviceStates  []DeviceState `json:"deviceStates,omitempty"` // One-To-Many relationship (has many)
}

func (d *Device) BeforeCreate() (err error) {
	d.CreatedAt = time.Now().Unix()
	d.UpdatedAt = time.Now().Unix()
	return
}

func (d *Device) BeforeUpdate() (err error) {
	d.UpdatedAt = time.Now().Unix()
	return
}

func (d *Device) BeforeDelete() (err error) {
	d.DeletedAt = time.Now().Unix()
	return
}
