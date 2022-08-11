package model

type UserAccount struct {
	*BaseModel
	userId     int    `json:"user_id"`
	platformId int    `json:"platform"`
	password   string `json:"password"`
}
