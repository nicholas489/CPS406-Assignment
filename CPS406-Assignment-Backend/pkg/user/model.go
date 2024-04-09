package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string  `json:"name"`
	Password    string  `json:"password"`
	Email       string  `json:"email" gorm:"index;unique"`
	PhoneNumber int     `json:"phone_number"`
	Balance     int     `json:"balance"`
	Events      []Event `gorm:"many2many:user_events;"` // Many-to-Many relationship
}

type Event struct {
	gorm.Model
	Name       string `json:"name" gorm:"index;unique"`
	CoachEmail string `json:"coach_email"` // Store email to fetch and associate Coach
	Location   string `json:"location"`
	Cost       int    `json:"cost"`
	Users      []User `gorm:"many2many:user_events;"` // Many-to-Many relationship
}

type ReceiveEvent struct {
	EventId   string `json:"event_id"`
	UserEmail string `json:"email"`
}

type UserBalance struct {
	UserEmail string `json:"email"`
	Balance   int    `json:"balance"`
}
