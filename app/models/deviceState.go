package models

import (
	"time"
)

/**
 * 貸出、返却
 */
type DeviceState struct {
	Id        int64 `json:"id,omitempty"`
	DeviceId  int64 `json:"device_id,omitempty"` // Foreign key for Device (belongs to)
	Action    bool  `json:"action,omitempty"`
	User      User  `json:"user,omitempty"`    // One-To-One relationship (has one)
	UserId    int64 `json:"user_id,omitempty"` // Foreign key of User
	CreatedAt int64 `json:"createdAt,omitempty"`
}

func (d *DeviceState) BeforeCreate() (err error) {
	d.CreatedAt = time.Now().Unix()
	return
}
