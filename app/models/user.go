package models

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
