package db

import (
	"CPS406-Assignment-Backend/pkg/coach"
	"CPS406-Assignment-Backend/pkg/finance"
	"CPS406-Assignment-Backend/pkg/user"
	"gorm.io/gorm"
	"log"
	"math/rand"
)

func SeedDatabase(db *gorm.DB) {
	// Start transaction
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		log.Fatalf("Failed to start transaction: %v", tx.Error)
	}

	// Seed coaches
	if err := seedCoaches(tx); err != nil {
		tx.Rollback()
		log.Fatalf("Failed to seed coaches: %v", err)
	}

	// Seed events
	if err := seedEvents(tx); err != nil {
		tx.Rollback()
		log.Fatalf("Failed to seed events: %v", err)
	}

	// Seed users
	if err := seedUsers(tx); err != nil {
		tx.Rollback()
		log.Fatalf("Failed to seed users: %v", err)
	}

	// Enroll users in events
	if err := enrollUsersInEvents(tx); err != nil {
		tx.Rollback()
		log.Fatalf("Failed to enroll users in events: %v", err)
	}
	// Populate finance
	if err := populateFinance(tx); err != nil {
		tx.Rollback()
		log.Fatalf("Failed to populate finance: %v", err)
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		log.Fatalf("Failed to commit transaction: %v", err)
	}
}

func seedCoaches(tx *gorm.DB) error {
	coaches := []coach.Coach{
		{Name: "CoachMike", Email: "mike@example.com", PhoneNumber: 1234567890, Password: "pass123", Owed: 1234},
		{Name: "CoachAnna", Email: "anna@example.com", PhoneNumber: 1234567891, Password: "pass456", Owed: 5678},
	}

	for _, c := range coaches {
		var count int64
		tx.Model(&coach.Coach{}).Where("email = ?", c.Email).Count(&count)
		if count == 0 {
			if err := tx.Create(&c).Error; err != nil {
				return err
			}
			log.Printf("Created coach: %s", c.Email)
		}
	}
	return nil
}

func seedEvents(tx *gorm.DB) error {
	events := []user.Event{
		{Name: "Morning Yoga", CoachID: 1, Location: "Central Park", Date: "2024-04-11", Time: "06:00:00", Cost: 10, EventExpenses: 1200, CoachExpenses: 1234},
		{Name: "Evening Run", CoachID: 2, Location: "Riverside", Date: "2024-04-12", Time: "06:00:00", Cost: 5, EventExpenses: 2500, CoachExpenses: 5678},
		{Name: "Advanced Cycling", CoachID: 1, Location: "West Side Highway", Date: "2024-04-13", Time: "09:00:00", Cost: 15, EventExpenses: 1800, CoachExpenses: 1234},
		{Name: "Boxing Class", CoachID: 2, Location: "Downtown Gym", Date: "2024-04-14", Time: "11:00:00", Cost: 20, EventExpenses: 2200, CoachExpenses: 5678},
		{Name: "Soccer Training", CoachID: 1, Location: "North Field", Date: "2024-04-15", Time: "14:00:00", Cost: 10, EventExpenses: 2500, CoachExpenses: 1234},
		{Name: "Basketball Workshop", CoachID: 2, Location: "City Arena", Date: "2024-04-16", Time: "16:00:00", Cost: 12, EventExpenses: 2000, CoachExpenses: 5678},
		{Name: "Swimming 101", CoachID: 1, Location: "Community Pool", Date: "2024-04-17", Time: "07:00:00", Cost: 18, EventExpenses: 3000, CoachExpenses: 1234},
	}

	for _, e := range events {
		var count int64
		tx.Model(&user.Event{}).Where("name = ? AND coach_email = ?", e.Name, e.CoachID).Count(&count)
		if count == 0 {
			if err := tx.Create(&e).Error; err != nil {
				return err
			}
			log.Printf("Created event: %s", e.Name)
		}
	}
	return nil
}

func seedUsers(tx *gorm.DB) error {
	users := []user.User{
		{Name: "John Doe", Email: "john.doe@example.com", PhoneNumber: 1234567892, Password: "secure123", Balance: 100},
		{Name: "Jane Smith", Email: "jane.smith@example.com", PhoneNumber: 1234567893, Password: "secure456", Balance: 150},
	}

	for _, u := range users {
		var count int64
		tx.Model(&user.User{}).Where("email = ?", u.Email).Count(&count)
		if count == 0 {
			if err := tx.Create(&u).Error; err != nil {
				return err
			}
			log.Printf("Created user: %s", u.Email)
		}
	}
	return nil
}

func enrollUsersInEvents(tx *gorm.DB) error {
	// Example: Enroll all users in all events
	// This is a simple approach; you might want to customize the logic
	// based on your application's needs (e.g., enroll users in specific events)

	var users []user.User
	if err := tx.Find(&users).Error; err != nil {
		return err
	}

	var events []user.Event
	if err := tx.Find(&events).Error; err != nil {
		return err
	}

	for _, u := range users {
		for _, e := range events {
			// Check if the user is already enrolled in the event to avoid duplicates
			var count int64
			tx.Table("user_events").Where("user_id = ? AND event_id = ?", u.ID, e.ID).Count(&count)
			if count == 0 {
				// Associate user with event
				if err := tx.Model(&u).Association("Events").Append(&e); err != nil {
					return err
				}
				log.Printf("Enrolled user %s in event %s", u.Email, e.Name)
			}
		}
	}
	return nil
}

func randomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// populateFinance populates the database with an OrganizationAccount, associated YearlyAccounts, and MonthlyAccounts
func populateFinance(tx *gorm.DB) error {
	orgAccount := finance.OrganizationAccount{Name: "dance_club"}
	if err := tx.Create(&orgAccount).Error; err != nil {
		return err // Return error if creation fails
	}

	// Seed yearly accounts from 2014 to 2024
	for year := 2014; year <= 2024; year++ {
		yearlyAccount := finance.YearlyAccount{
			Year:           year,
			Profit:         0, // Will be calculated
			Expenses:       0, // Will be calculated
			OrganizationID: orgAccount.ID,
		}

		// Generate data for each month
		for month := 1; month <= 12; month++ {
			monthlyProfit := randomFloat64(500, 5000)   // Random profit
			monthlyExpenses := randomFloat64(300, 4000) // Random expenses

			// Create and append monthly accounts
			monthlyAccount := finance.MonthlyAccount{
				Year:            year,
				Month:           month,
				Profit:          monthlyProfit,
				Expenses:        monthlyExpenses,
				YearlyAccountID: yearlyAccount.ID,
			}

			yearlyAccount.MonthlyAccounts = append(yearlyAccount.MonthlyAccounts, monthlyAccount)

			// Update yearly profit and expenses
			yearlyAccount.Profit += monthlyProfit
			yearlyAccount.Expenses += monthlyExpenses
		}

		// Save the yearly account along with its monthly accounts
		if err := tx.Create(&yearlyAccount).Error; err != nil {
			return err // Return error if creation fails
		}
	}

	return nil // Return nil if successful
}
