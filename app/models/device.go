package models

/**
 * 端末モデル
 */
type Device struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Manufacturer  string `json:"manufacturer,omitempty"`
	Carrier       string `json:"carrier,omitempty"`
	Os            string `json:"os,omitempty"`
	Size          string `json:"size,omitempty"`
	Resolution    string `json:"resolution,omitempty"`
	Memory        string `json:"memory,omitempty"`
	DateOfRelease int64  `json:"dataOfRelease,omitempty"`
	Other         string `json:"other,omitempty"`
	CreatedAt     int64  `json:"createdAt,omitempty"`
	UpdatedAt     int64  `json:"updatedAt,omitempty"`
	DeletedAt     int64  `json:"deletedAt,omitempty"`
	User          User   `json:"user,omitempty"`
}
