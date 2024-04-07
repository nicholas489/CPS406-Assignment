package user

import (
	"CPS406-Assignment-Backend/internal/util"
	"CPS406-Assignment-Backend/pkg/jwtM"
	"CPS406-Assignment-Backend/pkg/login"
	"CPS406-Assignment-Backend/pkg/user"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// GetAllUsers gets all the users from the database
func GetAllUsers(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.Header().Set("Content-Type", "text/plain")
	var users []user.User
	// get all the users from the database
	db.Find(&users)
	// send the users as a response list without password
	for i := range users {
		users[i].Password = ""
	}
	json.NewEncoder(w).Encode(users)
}

// GetUser gets a user from the database by id
func GetUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.Header().Set("Content-Type", "text/plain")
	// get the id from the url
	id := chi.URLParam(r, "id")
	// get the user from the database and save it in the user variable
	var user user.User
	db.Debug().Find(&user, "id = ?", id)
	// send the user as a response
	json.NewEncoder(w).Encode(user)
}
func GetEvents(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.Header().Set("Content-Type", "text/plain")
	// get all the events from the database
	// Get the id from the url
	id := chi.URLParam(r, "id")
	// Get the event list from user's list of events
	var user2 user.User
	db.First(&user2, "id = ?", id)
	var events []user.Event // Corrected variable declaration for events slice
	db.Model(&user2).Association("Events").Find(&events)
	// send the events as a response list
	json.NewEncoder(w).Encode(events)


// GetEvents gets all the events from the database that the user is in
func GetEvents(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.Header().Set("Content-Type", "text/plain")
	// get all the events from the database
	// Get the id from the url
	id := chi.URLParam(r, "id")
	// Get the event list from user's list of events
	var user2 user.User
	db.First(&user2, "id = ?", id)
	var events []user.Event // Corrected variable declaration for events slice
	db.Model(&user2).Association("Events").Find(&events)
	// send the events as a response list
	json.NewEncoder(w).Encode(events)

}

// PostLogin logs in the user


func PostLogin(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// Parse the request body
	var l login.Login
	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		util.SendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Check if the user exists in the database
	var user user.User
	result := db.First(&user, "email = ?", l.Email)
	// If the user does not exist, send a JSON error message
	if result.Error != nil {
		util.SendJSONError(w, "User not found", http.StatusNotFound)
		return
	}
	// If the user exists, check if the passwords match
	if user.Password != l.Password {
		util.SendJSONError(w, "Invalid password"+"password given: "+l.Password, http.StatusUnauthorized)
		return
	}
	// Generate a JWT token
	privileges := util.SetPrivileges(jwtM.CustomClaims{Privileges: jwtM.Privileges{User: true}})
	tokenString, err := util.GenerateJWT(user.Email, privileges)
	// If the token generation fails, send a JSON error message
	if err != nil {
		util.SendJSONError(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}
	// Set the token as a cookie
	util.SetTokenAsCookie(w, tokenString)
	// Send the user details
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message": "Login successful",
		"email":   user.Email,
		"name":    user.Name,
		"id":      strconv.Itoa(int(user.ID)),
	}
	json.NewEncoder(w).Encode(response)
}

// PostSignup creates a new user
func PostSignup(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// Parse the request body
	var u user.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		util.SendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Check if the user already exists
	var user user.User
	result := db.First(&user, "email = ? ", u.Email)

	// If the user already exists, send an error message
	if result.Error == nil {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}
	// If the user does not exist, create a new user
	db.Create(&u)
	// Generate a JWT token
	privileges := util.SetPrivileges(jwtM.CustomClaims{Privileges: jwtM.Privileges{User: true}})
	// If the token generation fails, send a JSON error message
	if err != nil {
		util.SendJSONError(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}
	// Set the token as a cookie
	tokenString, err := util.GenerateJWT(u.Email, privileges)
	util.SetTokenAsCookie(w, tokenString)
	// Send the user details and token as a response and set status code to 201 (Created)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message": "Signup successful",
		"email":   u.Email,
		"name":    u.Name,
		"id":      strconv.Itoa(int(u.ID)),
	}
	json.NewEncoder(w).Encode(response)
}

// JoinEvent adds a user to an event
func JoinEvent(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// Parse the request body
	var ue user.ReceiveEvent
	if err := json.NewDecoder(r.Body).Decode(&ue); err != nil {
		util.SendJSONError(w, err.Error()+"ee", http.StatusBadRequest)
		return
	}

	// Start a transaction to add the user to the event

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// Get the event from the database
	var e user.Event
	if result := tx.Preload("Users").Where("id = ?", ue.EventId).First(&e); result.Error != nil {
		tx.Rollback()
		util.SendJSONError(w, "Event not found", http.StatusNotFound)
		return
	}
	// Get the user from the database
	var u user.User
	if result := tx.Where("email = ?", ue.UserEmail).First(&u); result.Error != nil {
		tx.Rollback()
		util.SendJSONError(w, "User not found", http.StatusNotFound)
		return
	}
	// Check if the user is already in the event
	for _, user := range e.Users {
		if user.Email == u.Email {
			tx.Rollback()
			util.SendJSONError(w, "User already in event", http.StatusConflict)
			return
		}
	}
	// Add the user to the event and subtract the cost from the user's balance
	u.Balance -= e.Cost
	if result := tx.Save(&u); result.Error != nil {
		tx.Rollback()
		util.SendJSONError(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	// Append the user to the event's Users slice
	if err := tx.Model(&e).Association("Users").Append(&u); err != nil {
		tx.Rollback()
		util.SendJSONError(w, "Failed to add user to event", http.StatusInternalServerError)
		return
	}

	// Append the event to the user's Events slice
	u.Events = append(u.Events, e)
	// Save the user

	if err := tx.Save(&u).Error; err != nil {
		tx.Rollback()
		util.SendJSONError(w, "Failed to add event to user's list", http.StatusInternalServerError)
		return
	}

	// Commit the transaction

	if err := tx.Commit().Error; err != nil {
		util.SendJSONError(w, "Transaction commit error", http.StatusInternalServerError)
		return
	}
	// Send a success message
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message": "User joined event successfully",
		"event":   e.Name,
		"email":   u.Email,
		"id":      strconv.Itoa(int(u.ID)),
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		// Log the error of failing to encode or send the response
		fmt.Println("Error sending response:", err)
	}
}


// LeaveEvent removes a user from an event and adds the cost back to the user's balance
func LeaveEvent(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// Parse the request body

	var ue user.ReceiveEvent
	if err := json.NewDecoder(r.Body).Decode(&ue); err != nil {
		util.SendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Start a transaction to remove the user from the event

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Get the event from the database

	var e user.Event
	if result := tx.Preload("Users").Where("id = ?", ue.EventId).First(&e); result.Error != nil {
		tx.Rollback()
		util.SendJSONError(w, "Event not found", http.StatusNotFound)
		return
	}

	// Get the user from the database

	var u user.User
	if result := tx.Where("email = ?", ue.UserEmail).First(&u); result.Error != nil {
		tx.Rollback()
		util.SendJSONError(w, "User not found", http.StatusNotFound)
		return
	}
	// Check if the user is in the event

	for i, user2 := range e.Users {
		if user2.Email == u.Email {
			// Remove the user from the event's Users slice
			e.Users = append(e.Users[:i], e.Users[i+1:]...)
			// Save the event

			if result := tx.Save(&e); result.Error != nil {
				tx.Rollback()
				util.SendJSONError(w, result.Error.Error(), http.StatusInternalServerError)
				return
			}

			// Add the cost back to the user's balance
			u.Balance += e.Cost
			// Save the user

			if result := tx.Save(&u); result.Error != nil {
				tx.Rollback()
				util.SendJSONError(w, result.Error.Error(), http.StatusInternalServerError)
				return
			}

			// Also remove the event from the user's list of events

			if err := tx.Model(&u).Association("Events").Delete(&e); err != nil {
				tx.Rollback()
				util.SendJSONError(w, "Failed to remove event from user", http.StatusInternalServerError)
				return
			}

			// Commit the transaction
			tx.Commit()
			// Send a success message

			w.Header().Set("Content-Type", "application/json")
			response := map[string]string{
				"message": "User left event successfully",
				"event":   e.Name,
				"email":   u.Email,
			}

			// Send the response

			if err := json.NewEncoder(w).Encode(response); err != nil {
				// Log the error of failing to encode or send the response
				fmt.Println("Error sending response:", err)
			}
			return
		}
	}

	// If the user is not in the event, send an error message
	tx.Rollback()
	util.SendJSONError(w, "User not in event", http.StatusNotFound)

