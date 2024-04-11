package util

import (
	"CPS406-Assignment-Backend/pkg/coach"
	"CPS406-Assignment-Backend/pkg/jwtM"
	"CPS406-Assignment-Backend/pkg/user"
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"net/http"
	"os"
	"time"
)

// SetTokenAsCookie sets the JWT token as a cookie
func SetTokenAsCookie(w http.ResponseWriter, token string) {
	// Create a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",                   // Cookie name
		Value:    token,                          // JWT token
		Path:     "/",                            // Cookie path
		Expires:  time.Now().Add(24 * time.Hour), // Cookie expiration time
		HttpOnly: true,                           // Accessible only by the web server, helps mitigate XSS
		Secure:   true,                           // Ensure cookie is sent over HTTPS only
		SameSite: http.SameSiteStrictMode,        // Strict mode for CSRF mitigation
	})
}

// GenerateJWT generates a JWT token
func GenerateJWT(username string, privileges jwtM.Privileges) (string, error) {
	// Create an instance of CustomClaims
	claims := jwtM.CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			// Expires 30 days from now
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
			Issuer:    "go_app",
		},
		Username:   username,
		Privileges: privileges,
	}

	// Create a new JWT token using the HS256 signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with your secret key
	secret := os.Getenv("SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		println(err.Error())
		return "", err
	}

	return tokenString, nil
}

// SendJSONError sends a JSON error message
func SendJSONError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

// SetPrivileges sets the privileges for authentication
func SetPrivileges(user jwtM.CustomClaims) jwtM.Privileges {
	privileges := jwtM.Privileges{}
	if user.Privileges.Admin {
		privileges.Admin = true
	}
	if user.Privileges.User {
		privileges.User = true
	}
	if user.Privileges.Coach {
		privileges.Coach = true
	}
	return privileges
}

// JwtMiddlewareUser is a middleware that checks if the user is authenticated
func JwtMiddlewareUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the Cookie
		cookie, err := r.Cookie("auth_token")
		tokenString := cookie.Value
		if err != nil {
			SendJSONError(w, "No token provided", http.StatusUnauthorized)
			return
		}
		if tokenString == "" {
			SendJSONError(w, "No token provided", http.StatusUnauthorized)
			return
		}

		// Parse the token

		token, err := jwt.ParseWithClaims(tokenString, &jwtM.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil {
			SendJSONError(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Check if the token is valid
		if !token.Valid {
			SendJSONError(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		if claims, ok := token.Claims.(*jwtM.CustomClaims); ok {
			if !claims.Privileges.User {
				SendJSONError(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

// JwtMiddlewareAdmin is a middleware that checks if the user is an admin
func JwtMiddlewareAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the Authorization header
		cookie, err := r.Cookie("auth_token")
		tokenString := cookie.Value
		if tokenString == "" {
			SendJSONError(w, "No token provided", http.StatusUnauthorized)
			return
		}

		// Parse the token

		token, err := jwt.ParseWithClaims(tokenString, &jwtM.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil {
			SendJSONError(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Check if the token is valid
		if !token.Valid {
			SendJSONError(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		if claims, ok := token.Claims.(*jwtM.CustomClaims); ok {
			if !claims.Privileges.Admin {
				SendJSONError(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

// JwtMiddlewareCoach is a middleware that checks if the user is a coach
func JwtMiddlewareCoach(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the Authorization header
		cookie, err := r.Cookie("auth_token")
		tokenString := cookie.Value
		if tokenString == "" {
			SendJSONError(w, "No token provided", http.StatusUnauthorized)
			return
		}

		// Parse the token

		token, err := jwt.ParseWithClaims(tokenString, &jwtM.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil {
			SendJSONError(w, "Invalid token, problem with token"+err.Error()+tokenString, http.StatusUnauthorized)
			return
		}

		// Check if the token is valid
		if !token.Valid {
			SendJSONError(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		if claims, ok := token.Claims.(*jwtM.CustomClaims); ok {
			if !claims.Privileges.Coach {
				SendJSONError(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

// CombinedJwtMiddleware is a middleware that combines the admin and coach middleware
func CombinedJwtMiddleware(adminMiddleware, coachMiddleware func(http.Handler) http.Handler) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Attempt to use the admin middleware
			adminPassed := false
			adminMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				adminPassed = true
				next.ServeHTTP(w, r)
			})).ServeHTTP(w, r)

			// If admin check didn't pass, try the coach middleware
			if !adminPassed {
				coachMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					next.ServeHTTP(w, r)
				})).ServeHTTP(w, r)
			}
		})
	}
}

// CheckCookie checks the cookie for the JWT token
func CheckCookie(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// Get the token from the Cookie
	cookie, err := r.Cookie("auth_token")
	if err != nil {
		SendJSONError(w, "No token provided", http.StatusUnauthorized)
		return
	}
	// Check if the token is empty
	if cookie.Value == "" {
		SendJSONError(w, "No token provided", http.StatusUnauthorized)
		return
	}
	// Parse the token
	tokenString := cookie.Value
	token, err := jwt.ParseWithClaims(tokenString, &jwtM.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	// Check if the token is valid
	if err != nil {
		SendJSONError(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	if claims, ok := token.Claims.(*jwtM.CustomClaims); ok && token.Valid {
		// Make a response struct to remove the issuer and expiration time
		type Response struct {
			Email      string          `json:"email"`
			Id         uint            `json:"id"`
			Privileges jwtM.Privileges `json:"privileges"`
		}
		// Send the claims as a response
		var response Response
		if claims.Privileges.Coach {
			var coach coach.Coach
			db.First(&coach, "email = ?", claims.Username)
			response.Id = coach.ID
		} else {
			var user user.User
			db.First(&user, "email = ?", claims.Username)
			response.Id = user.ID
		}

		response.Email = claims.Username
		response.Privileges = claims.Privileges
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}
}

// Logout logs out the user
func Logout(w http.ResponseWriter) {
	// Remove the cookie by setting it to an empty value and an expiration time in the past
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-1 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})
	// Send a 200 OK response
	w.WriteHeader(http.StatusOK)
}
