package db

import (
	"CPS406-Assignment-Backend/pkg/coach"
	"CPS406-Assignment-Backend/pkg/user"
	"gorm.io/gorm"
	"log"
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

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		log.Fatalf("Failed to commit transaction: %v", err)
	}
}

func seedCoaches(tx *gorm.DB) error {
	coaches := []coach.Coach{
		{Name: "CoachMike", Email: "mike@example.com", PhoneNumber: 1234567890, Password: "pass123"},
		{Name: "CoachAnna", Email: "anna@example.com", PhoneNumber: 1234567891, Password: "pass456"},
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
		{Name: "Morning Yoga", CoachEmail: "mike@example.com", Location: "Central Park", Cost: 10},
		{Name: "Evening Run", CoachEmail: "anna@example.com", Location: "Riverside", Cost: 5},
	}

	for _, e := range events {
		var count int64
		tx.Model(&user.Event{}).Where("name = ? AND coach_email = ?", e.Name, e.CoachEmail).Count(&count)
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
