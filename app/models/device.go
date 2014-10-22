package models

import (
	"time"
)

type Device struct { // example user fields
	Id           int64
	Name         string
	Manufacturer string
	Carrier      string
	OS           string
	Size         string
	Resolution   string
	Memory       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time // for soft delete
}
