package model

type User struct {
	*BaseModel
	userId      int    `json:"user_id"`
	userName    string `json:"user_name"`
	password    string `json:"password"`
	phoneNumber string `json:"phone_number"`
	shareMode   int    `json:"share_mode"`
	sex         int    `json:"gender"`
	desc        string `json:"desc"`
	role        int    `json:"role"`
}

