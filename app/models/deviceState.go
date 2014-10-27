package models

import (
	"time"
)

/**
 * 貸出、返却
 */
type DeviceState struct {
	Id        int64 `json:"id,omitempty"`
	DeviceId  int64 `json:"device_id,omitempty"`
	Action    bool  `json:"action,omitempty"`
	User      User  `json:"user,omitempty"`
	CreatedAt int64 `json:"createdAt,omitempty"`
}

func (d *DeviceState) BeforeCreate() (err error) {
	d.CreatedAt = time.Now().Unix()
	return
}
