package model

type Base struct {
	Id        string `json:"id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt *int64 `json:"updated_at"`
	DeletedAt *int64 `json:"deleted_at"`
}
