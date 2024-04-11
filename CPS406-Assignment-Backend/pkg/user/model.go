package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name                    string  `json:"name"`
	Password                string  `json:"password"`
	Email                   string  `json:"email" gorm:"index;unique"`
	PhoneNumber             int     `json:"phone_number"`
	InAdvancePaymentCounter int     `json:"in_advance_payment_counter"`
	Balance                 int     `json:"balance"`
	Events                  []Event `gorm:"many2many:user_events;"` // Many-to-Many relationship
}

type Event struct {
	gorm.Model
	Name          string `json:"name" gorm:"index;unique"`
	CoachID       int    `json:"coach_id"` // Store email to fetch and associate Coach
	Location      string `json:"location"`
	Date          string `json:"date"`
	Time          string `json:"time"`
	Cost          int    `json:"cost"`
	Users         []User `gorm:"many2many:user_events;"` // Many-to-Many relationship
	EventExpenses int    `json:"event_expenses"`
	CoachExpenses int    `json:"coach_expenses"`
}

type ReceiveEvent struct {
	EventId   int    `json:"event_id"`
	UserEmail string `json:"email"`
}

type UserBalance struct {
	UserEmail        string `json:"email"`
	Amount           int    `json:"Amount"`
	InAdvancePayment bool   `json:"in_advance_payment"`
}
