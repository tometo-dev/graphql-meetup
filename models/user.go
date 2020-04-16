package models

type User struct {
	ID       string    `json:"id" gorm:"column:id; primary_key"`
	Username string    `json:"username" gorm:"column:username" validate:"required"`
	Email    string    `json:"email" gorm:"column:email" validate:"required"`
}
