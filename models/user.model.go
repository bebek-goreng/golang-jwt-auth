package models

import "time"

type User struct {
	Id        int64     `gorm:"type:primaryKey" json:"id"`
	FirstName string    `gorm:"type:varchar(200); not null" json:"first_name"`
	LastName  string    `gorm:"type:varchar(200)" json:"last_name"`
	Email     string    `gorm:"type:varchar(200); not null; unique" json:"email"`
	Password  string    `gorm:"type:varchar(200); not null" json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
