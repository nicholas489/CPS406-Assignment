package coach

import (
	"CPS406-Assignment-Backend/internal/util"
	"CPS406-Assignment-Backend/pkg/coach"
	"CPS406-Assignment-Backend/pkg/jwtM"
	"CPS406-Assignment-Backend/pkg/user"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"net/http"
)

// PostEvent creates an event in the database
func PostEvent(writer http.ResponseWriter, request *http.Request, db *gorm.DB) {
	// Parse the request body
	var event, existingEvent user.Event
	err := json.NewDecoder(request.Body).Decode(&event)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	// If the event already exists
	result := db.First(&existingEvent, "name = ?", event.Name)
	if result.Error == nil {
		util.SendJSONError(writer, "Event already exists", http.StatusConflict)
		return
	}
	// Make event and putting an empty user list
	event.Users = []user.User{}
	// Create the event in the database
	db.Create(&event)
	//Update the coach expenses
	var coach coach.Coach
	db.First(&coach, "id = ?", event.CoachID)
	coach.Owed += event.CoachExpenses
	db.Save(&coach)
	// Send the event as a response and set the status code to 201 (Created)
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(event)
}

// GetEvent gets an event from the database
func GetEvent(writer http.ResponseWriter, request *http.Request, db *gorm.DB) {
	// Get the id from the url
	id := chi.URLParam(request, "id")

	// Get the event from the database and save it in the event variable
	var event user.Event
	db.First(&event, "id = ?", id)
	// Getting all the users that are in the event
	err := db.Model(&event).Association("Users").Find(&event.Users)
	// Print all the users in the event
	if err != nil {
		util.SendJSONError(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the event as a response
	json.NewEncoder(writer).Encode(event)
}

// GetEvents gets all the events from the database
func GetEvents(writer http.ResponseWriter, request *http.Request, db *gorm.DB) {
	// Get all the events from the database
	var events []user.Event
	db.Find(&events)
	// Send the events as a response
	json.NewEncoder(writer).Encode(events)
}

// PostLogin logs in a coach
func PostLogin(writer http.ResponseWriter, request *http.Request, db *gorm.DB) {
	// Parse the request body
	var login coach.Coach
	err := json.NewDecoder(request.Body).Decode(&login)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	// Check if the coach exists in the database
	var coach coach.Coach
	result := db.First(&coach, "email = ?", login.Email)
	if result.Error != nil {
		util.SendJSONError(writer, "Coach not found", http.StatusNotFound)
		return
	}
	// If the coach exists, check if the passwords match
	if coach.Password != login.Password {
		util.SendJSONError(writer, "Invalid password", http.StatusUnauthorized)
		return
	}
	privileges := util.SetPrivileges(jwtM.CustomClaims{Privileges: jwtM.Privileges{Coach: true, User: true}})
	// Generate a JWT token
	token, err := util.GenerateJWT(coach.Email, privileges)
	if err != nil {
		util.SendJSONError(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	util.SetTokenAsCookie(writer, token)
	// Send the token as a response
	writer.WriteHeader(http.StatusOK)
	// Send the coach as a response
	json.NewEncoder(writer).Encode(coach)
}

// PostSignup signs up a coach
func PostSignup(writer http.ResponseWriter, request *http.Request, db *gorm.DB) {
	// Parse the request body
	var coach coach.Coach
	err := json.NewDecoder(request.Body).Decode(&coach)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	// If the coach already exists
	result := db.First(&coach, "email = ?", coach.Email)
	if result.Error == nil {
		util.SendJSONError(writer, "Coach already exists", http.StatusConflict)
		return
	}
	// Generate a JWT token
	privileges := util.SetPrivileges(jwtM.CustomClaims{Privileges: jwtM.Privileges{Coach: true, User: true}})
	token, err := util.GenerateJWT(coach.Email, privileges)
	if err != nil {
		util.SendJSONError(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	util.SetTokenAsCookie(writer, token)
	// Create the coach in the database
	db.Create(&coach)
	// Send the coach as a response and set the status code to 201 (Created)
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(coach)
}

// DeleteUser deletes a user from the database
func DeleteUser(writer http.ResponseWriter, request *http.Request, db *gorm.DB) {
	// Get the email from the url

	id := chi.URLParam(request, "id")

	// Get the user from the database
	var user user.User
	db.First(&user, "id = ?", id)

	// Delete the user from the database
	db.Delete(&user)

	// Send the user as a response
	json.NewEncoder(writer).Encode(user)
}

// GetCoach gets a coach from the database by id
func GetCoach(writer http.ResponseWriter, request *http.Request, db *gorm.DB) {
	var coach coach.Coach
	// Get the id from the url
	id := chi.URLParam(request, "id")
	// Get the coach from the database
	db.First(&coach, "id = ?", id)
	// Send the coach as a response
	json.NewEncoder(writer).Encode(coach)
}

// GetAllCoaches gets all the coaches from the database
func GetAllCoaches(writer http.ResponseWriter, request *http.Request, db *gorm.DB) {
	// Get all the coaches from the database
	var coaches []coach.Coach
	// Find all the coaches
	db.Find(&coaches)
	// Send the coaches as a response
	json.NewEncoder(writer).Encode(coaches)
}

func GetOwed(writer http.ResponseWriter, request *http.Request, db *gorm.DB) {
	type Owed struct {
		Name string `json:"name"`
		Owed int    `json:"owed"`
	}
	var coach coach.Coach
	// Get the id from the url
	id := chi.URLParam(request, "id")
	// Get the coach from the database
	db.First(&coach, "id = ?", id)
	// Get the amount owed
	var owed Owed
	owed.Name = coach.Name
	owed.Owed = coach.Owed
	// Send the owed as a response
	json.NewEncoder(writer).Encode(owed)
}
