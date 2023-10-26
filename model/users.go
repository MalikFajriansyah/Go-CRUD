package model

type Users struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
