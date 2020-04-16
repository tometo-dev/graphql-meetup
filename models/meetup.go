package models

type Meetup struct {
	ID          string `json:"id" gorm:"column:id; primary_key"`
	Name        string `json:"name" gorm:"column:name" validate:"required"`
	Description string `json:"description" gorm:"column:description" validate:"required"`
	UserID      string `json:"user_id" gorm:"column:user_id"`
}
