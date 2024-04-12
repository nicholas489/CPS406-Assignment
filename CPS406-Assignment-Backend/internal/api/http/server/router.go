package server

import (
	"CPS406-Assignment-Backend/internal/api/http/coach"
	"CPS406-Assignment-Backend/internal/api/http/finance"
	"CPS406-Assignment-Backend/internal/api/http/user"
	"CPS406-Assignment-Backend/internal/util"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"net/http"
)

func Server(r chi.Router, db *gorm.DB) {
	// Routes for the API
	// Route for the login
	r.Route("/login", func(r chi.Router) {
		// Login for the user
		r.Post("/user", func(writer http.ResponseWriter, request *http.Request) {
			user.PostLogin(writer, request, db)
		})
		// Login for the coach
		r.Post("/coach", func(writer http.ResponseWriter, request *http.Request) {
			coach.PostLogin(writer, request, db)
		})
	})
	// Route for the signup
	r.Route("/signup", func(r chi.Router) {
		// Signup for the user
		r.Post("/user", func(writer http.ResponseWriter, request *http.Request) {
			user.PostSignup(writer, request, db)
		})
		// Signup for the coach
		r.Post("/coach", func(writer http.ResponseWriter, request *http.Request) {
			coach.PostSignup(writer, request, db)
		})
	})
	// Route for the logout
	r.Route("/logout", func(r chi.Router) {
		r.Post("/", func(writer http.ResponseWriter, request *http.Request) {
			util.Logout(writer)
		})

	})
	// Route for the user
	r.Route("/user", func(r chi.Router) {
		r.Use(util.JwtMiddlewareUser)
		// Get the user by id
		r.Get("/{id}", func(writer http.ResponseWriter, request *http.Request) {
			user.GetUser(writer, request, db)

		})
		// Get the events that the user is in
		r.Get("/{id}/events", func(writer http.ResponseWriter, request *http.Request) {
			user.GetEvents(writer, request, db)
		})
		r.Get("/{id}/events/count", func(writer http.ResponseWriter, request *http.Request) {
			user.GetEventsCount(writer, request, db)
		})
		// Get all the users
		r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			user.GetAllUsers(writer, request, db)

		})
		// Delete the user from coach perspective
		r.With(util.JwtMiddlewareCoach).Delete("/{id}", func(writer http.ResponseWriter, request *http.Request) {
			coach.DeleteUser(writer, request, db)
		})

		r.Put("/{id}/pay", func(writer http.ResponseWriter, request *http.Request) {
			finance.PayBalanceUser(writer, request, db)
		})

	})
	// Route for the coach
	r.Route("/coach", func(r chi.Router) {
		// Get the coach by id
		r.Use(util.CombinedJwtMiddleware(util.JwtMiddlewareCoach, util.JwtMiddlewareAdmin))
		r.Get("/{id}", func(writer http.ResponseWriter, request *http.Request) {
			coach.GetCoach(writer, request, db)
		})
		// Get all the coaches
		r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			coach.GetAllCoaches(writer, request, db)
		})
		r.Get("/owed/{id}", func(writer http.ResponseWriter, request *http.Request) {
			coach.GetOwed(writer, request, db)
		})
		r.Post("/{id}/pay", func(writer http.ResponseWriter, request *http.Request) {
			finance.PayOwedCoach(writer, request, db)
		})
	})
	// Route for the event
	r.Route("/event", func(r chi.Router) {
		// Get all the events
		r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			coach.GetEvents(writer, request, db)
		})
		// Make an event
		r.With(util.JwtMiddlewareCoach).Post("/", func(writer http.ResponseWriter, request *http.Request) {
			coach.PostEvent(writer, request, db)
		})

		// Get an event by id

		r.Get("/{id}", func(writer http.ResponseWriter, request *http.Request) {
			coach.GetEvent(writer, request, db)
		})
		// Join an event
		r.With(util.JwtMiddlewareUser).Post("/join", func(writer http.ResponseWriter, request *http.Request) {
			user.JoinEvent(writer, request, db)
		})

		// Leave an event
		r.With(util.JwtMiddlewareUser).Delete("/leave", func(writer http.ResponseWriter, request *http.Request) {
			user.LeaveEvent(writer, request, db)
		})

	})
	//Route for the auth
	r.Route("/auth", func(r chi.Router) {
		// Check the cookie
		r.Post("/session", func(writer http.ResponseWriter, request *http.Request) {
			util.CheckCookie(writer, request, db)
		})
	})
	// Route for the finance
	r.Route("/finance", func(r chi.Router) {
		// Get all the years of finance
		r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			finance.GetOrganizationAccount(writer, request, db)
		})
		// Get the finance by year
		r.Get("/{year}", func(writer http.ResponseWriter, request *http.Request) {
			finance.GetYearlyAccount(writer, request, db)
		})
		// Get the finance by month
		r.Get("/{year}/{month}", func(writer http.ResponseWriter, request *http.Request) {
			finance.GetMonthlyAccount(writer, request, db)
		})
	})
}
