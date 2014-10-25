package models

import "time"

/**
 * 端末モデル
 */
type Device struct {
	Id            int64     `json:"id,omitempty"`
	Name          string    `json:"name,omitempty"`
	Manufacturer  string    `json:"manufacturer,omitempty"`
	Carrier       string    `json:"carrier,omitempty"`
	Os            string    `json:"os,omitempty"`
	Size          string    `json:"size,omitempty"`
	Resolution    string    `json:"resolution,omitempty"`
	Memory        string    `json:"memory,omitempty"`
	DateOfRelease time.Time `json:"dataOfRelease,omitempty"`
	Other         string    `json:"other,omitempty"`
	CreatedAt     time.Time `json:"createdAt,omitempty"`
	UpdatedAt     time.Time `json:"updatedAt,omitempty"`
	DeletedAt     time.Time `json:"deletedAt,omitempty"`
	User          User      `json:"user,omitempty"`
}
