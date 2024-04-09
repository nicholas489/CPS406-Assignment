package finance

import (
	"CPS406-Assignment-Backend/internal/util"
	"CPS406-Assignment-Backend/pkg/finance"
	"CPS406-Assignment-Backend/pkg/user"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"net/http"
)

// GetOrganizationAccount gets the organization account with all yearly accounts
func GetOrganizationAccount(writer http.ResponseWriter, request *http.Request, db *gorm.DB) {
	// Get the organization account
	var orgAccount finance.OrganizationAccount
	if err := db.First(&orgAccount).Error; err != nil {
		util.SendJSONError(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	// Get all yearly accounts to the organization account from getYearlyAccounts
	listOfYears, err := getYearlyAccounts(db, orgAccount)
	if err != nil {
		util.SendJSONError(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	// Add them to the response
	orgAccount.YearlyAccounts = listOfYears
	// Respond with the organization account
	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(orgAccount)
	if err != nil {
		return
	}
}

// getYearlyAccounts gets all yearly accounts to the organization account
func getYearlyAccounts(db *gorm.DB, orgAccount finance.OrganizationAccount) ([]finance.YearlyAccount, error) {
	// Get all yearly accounts to the organization account
	var yearlyAccounts []finance.YearlyAccount
	if err := db.Model(&orgAccount).Association("YearlyAccounts").Find(&yearlyAccounts); err != nil {
		return nil, err
	}
	println("yearlyAccounts", yearlyAccounts)
	return yearlyAccounts, nil
}

// GetYearlyAccount gets the yearly account by year and all monthly accounts
func GetYearlyAccount(writer http.ResponseWriter, request *http.Request, db *gorm.DB) {
	// Extract the year from the URL parameters
	year := chi.URLParam(request, "year")

	var yearlyAccount finance.YearlyAccount
	// Preload MonthlyAccounts while getting the YearlyAccount to avoid separate query for monthly accounts
	if err := db.Preload("MonthlyAccounts").Where("year = ?", year).First(&yearlyAccount).Error; err != nil {
		util.SendJSONError(writer, "No such year! Try 2014..=2024", http.StatusNotFound) // Use StatusNotFound for resources not found
		return
	}

	// No need to manually assign listOfMonths to yearlyAccount.MonthlyAccounts, as Preload has already done that
	println("yearlyAccount", yearlyAccount.MonthlyAccounts)
	// Respond with the yearly account in JSON format
	writer.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(yearlyAccount); err != nil {
		util.SendJSONError(writer, "Failed to encode yearly account to JSON", http.StatusInternalServerError)
		return
	}
}

// GetMonthlyAccount gets the monthly account by month and year
func GetMonthlyAccount(writer http.ResponseWriter, request *http.Request, db *gorm.DB) {
	// Get monthly account by id and year
	id := chi.URLParam(request, "month")
	year := chi.URLParam(request, "year")
	var monthlyAccount finance.MonthlyAccount
	// Get the monthly account from the database
	if err := db.First(&monthlyAccount, "month = ? AND year = ?", id, year).Error; err != nil {
		util.SendJSONError(writer, "No such month!,try 1..=12 or year,try 2014..=2024", http.StatusInternalServerError)
		return
	}
	// Respond with the monthly account
	writer.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(monthlyAccount)
	if err != nil {
		return
	}

}

// UpdateBalanceUser updates the balance of the user
func UpdateBalanceUser(writer http.ResponseWriter, request *http.Request, db *gorm.DB) {
	// Get the user ID from the URL parameters
	id := chi.URLParam(request, "id")
	// Get the user from the database
	var user2 user.User
	if err := db.First(&user2, id).Error; err != nil {
		util.SendJSONError(writer, "No such user!", http.StatusNotFound)
		return
	}
	// Decode the request body into a balance struct
	var balance user.UserBalance
	if err := json.NewDecoder(request.Body).Decode(&balance); err != nil {
		util.SendJSONError(writer, "Failed to decode balance", http.StatusBadRequest)
		return
	}
	// Update the user's balance
	user2.Balance -= balance.Balance
	if err := db.Save(&user2).Error; err != nil {
		util.SendJSONError(writer, "Failed to update balance", http.StatusInternalServerError)
		return
	}
	// Respond with the updated user
	writer.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(user2); err != nil {
		return
	}
}
