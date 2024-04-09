package finance

import "gorm.io/gorm"

// OrganizationAccount represents the account details for the whole organization.
type OrganizationAccount struct {
	gorm.Model
	Name           string          `gorm:"default:'dance_club'" json:"name"`
	MaxProfit      float64         `json:"max_profit"` // Added field for the maximum profit
	YearlyAccounts []YearlyAccount `gorm:"foreignKey:OrganizationID"`
}

// YearlyAccount represents the account details for a specific year.
type YearlyAccount struct {
	gorm.Model
	Year            int              `json:"year"`
	Profit          float64          `json:"profit"`
	Expenses        float64          `json:"expenses"`
	OrganizationID  uint             `json:"-"`
	MonthlyAccounts []MonthlyAccount `gorm:"foreignKey:YearlyAccountID"`
}

// MonthlyAccount represents the account details for a specific month.
type MonthlyAccount struct {
	gorm.Model
	Year            int     `json:"year"`
	Month           int     `json:"month"`
	Profit          float64 `json:"profit"`
	Expenses        float64 `json:"expenses"`
	YearlyAccountID uint    `json:"-"` // Excluded from JSON to prevent redundancy
}
