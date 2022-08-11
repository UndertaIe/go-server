package model


type BaseModel struct {
	createdAt time.time `json:"created_at"`
	modifiedAt time.time `json:"modified_at"`
	deletedAt time.time `json:"deleted_at"`
	is_deleted bool `json:"is_deleted"`
}