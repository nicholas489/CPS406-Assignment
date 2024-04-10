# CPS406-Assignment-Backend
## Tech Stack
- ### GoLang
- ### GORM
- ### SQLite3
- ### JWT

## File Tree Structure

Below is the structure of the project, explaining the purpose of each file and directory:

```plaintext
CPS406-Assignment-Backend/
├── cmd/                           # Command-line interface applications
│   └── app/                       # Main application entry point
│       └── main.go                # The main Go file where execution begins
├── internal/                      # Private application and library code
│   └── api/                       # Internal API-related code
│       └── http/                  # HTTP transport layer specific code
│           ├── coach/             # Coach-related HTTP handlers
│           │   └── coach.go       # Handles coach-related requests
│           ├── finance/           # Finance-related HTTP handlers
│           │   └── finance.go     # Handles finance-related requests
│           ├── server/            # Server configuration and setup
│           │   └── router.go      # Routes and server setup
│           └── user/              # User-related HTTP handlers
│               └── user.go        # Handles user-related requests
│   └── db/                        # Database configuration and initialization
│       └── db.go                  # Database connection setup and GORM setup
│       └── seed.go                # Seed data for testing
│   └── util/                      # Utility functions used across the application
│       ├── sort.go                # Utility functions for sorting
│       └── util.go                # General utility functions
├── pkg/                           # Library code that's ok to use by external applications
│   ├── coach/                     # Domain model definitions for coach
│   │   └── model.go               # Coach model definition
│   ├── finance/                   # Domain model definitions for finance
│   │   └── model.go               # Finance model definition
│   ├── event/                     # Domain model definitions for event
│   │   └── model.go               # Event model definition
│   ├── jwtM/                      # Domain model definitions related to JWT middleware
│   │   └── model.go               # JWT middleware model definition
│   ├── login/                     # Domain model definitions for login functionality
│   │   └── model.go               # Login model definition
│   └── user/                      # Domain model definitions for user
│       └── model.go               # User model definition
└── .env                           # Environment variables and configuration settings
└── go.mod                         # Go module definition
└── README                         # This file

```

## Data Models

The following tables describe the data models used in the application and their respective fields:

### Coach

| Field       | Type   | JSON Key       | GORM Annotation | Description                              |
|-------------|--------|----------------|-----------------|------------------------------------------|
| ID          | uint   | -              | -               | The primary key (auto-incremented).      |
| CreatedAt   | time   | -              | -               | Timestamp of creation.                   |
| UpdatedAt   | time   | -              | -               | Timestamp of last update.                |
| DeletedAt   | time   | -              | -               | Timestamp of deletion (if soft deleted). |
| Name        | string | `name`         | -               | Unique username for the coach.           |
| Email       | string | `email`        | `index;unique`  | Unique email address for the coach.      |
| PhoneNumber | int    | `phone_number` | -               | Contact phone number for the coach.      |
| Password    | string | `password`     | -               | Hashed password for the coach's account. |
| Owed        | int    | `owed`         | -               | Amount owed to the coach.                | 

### Event

| Field      | Type        | JSON Key      | GORM Annotation | Description                              |
|------------|-------------|---------------|-----------------|------------------------------------------|
| ID         | uint        | -             | -               | The primary key (auto-incremented).      |
| CreatedAt  | time        | -             | -               | Timestamp of creation.                   |
| UpdatedAt  | time        | -             | -               | Timestamp of last update.                |
| DeletedAt  | time        | -             | -               | Timestamp of deletion (if soft deleted). |
| Name       | string      | `name`        | `index;unique`  | Unique name for the event.               |
| CoachEmail | string      | `coach_email` | -               | Email of the coach hosting the event.    |
| Location   | string      | `location`    | -               | Location where the event will be held.   |
| Cost       | int         | `cost`        | -               | Cost to attend the event.                |
| Users      | []user.User | `users`       | -               | List of users attending the event.       |
| EventExpenses | int         | `event_expenses` | - | Cost of running the event.               |
| CoachExpenses | int         | `coach_expenses` | - | Amount owed to coach.                    |

### User

| Field       | Type   | JSON Key       | GORM Annotation | Description                              |
|-------------|--------|----------------|-----------------|------------------------------------------|
| ID          | uint   | -              | -               | The primary key (auto-incremented).      |
| CreatedAt   | time   | -              | -               | Timestamp of creation.                   |
| UpdatedAt   | time   | -              | -               | Timestamp of last update.                |
| DeletedAt   | time   | -              | -               | Timestamp of deletion (if soft deleted). |
| Name        | string | `name`         | -               | Full name of the user.                   |
| Password    | string | `password`     | -               | Hashed password for the user's account.  |
| Email       | string | `email`        | `index;unique`  | Unique email address for the user.       |
| PhoneNumber | int    | `phone_number` | -               | Contact phone number for the user.       |
| Balance     | int    | `balance`      | -               | Balance amount for the user's account.   |
| EventID     | uint   | `event_id`     | -               | Foreign key relating to an Event.        |

## Seeded Data for Testing

Below are tables with sample data that is inserted into the database by the seeder function for testing purposes.

### Coaches

| UserName  | Email                | PhoneNumber | Password   | Owed       |
|-----------|----------------------|-------------|------------|------------|
| CoachMike | mike@example.com     | 1234567890  | pass123    | 1234       |
| CoachAnna | anna@example.com     | 1234567891  | pass456    | 5678       |

### Events

| Name          | Coach Email       | Location      | Cost | Event Expenses | Coach Expenses |
|---------------|-------------------|---------------|------|----------------|----------------|
| Morning Yoga  | mike@example.com  | Central Park  | 10   | 1200           | 1234           |
| Evening Run   | anna@example.com  | Riverside     | 5    | 2500           | 5678           |

### Users

| Name       | Email                  | PhoneNumber | Password  | Balance |
|------------|------------------------|-------------|-----------|---------|
| John Doe   | john.doe@example.com   | 1234567892  | secure123 | 100     |
| Jane Smith | jane.smith@example.com | 1234567893  | secure456 | 150     |

This data is meant for initial development and testing only and should not be used in production environments.


## API Routes Documentation

Below is the detailed table of the API routes, their descriptions, middleware, and the required JSON structure for requests (where applicable).


| Method | Endpoint         | Description                                                                 | Middleware                  | Required JSON Structure                                                                                |
|--------|------------------|-----------------------------------------------------------------------------|-----------------------------|--------------------------------------------------------------------------------------------------------|
| POST   | `api/login/user` | Logs in a user                                                              | None                        | `{ "email": "user@example.com", "password": "password123" }`                                           |
| POST   | `api/login/coach` | Logs in a coach                                                             | None                        | `{ "email": "coach@example.com", "password": "password123" }`                                          |
| POST   | `api/signup/user` | Signs up a new user                                                         | None                        | `{ "name": "John Doe", "email": "newuser@example.com", "password": "password123" }`                    |
| POST   | `api/signup/coach` | Signs up a new coach                                                        | None                        | `{ "name": "Jane Doe", "email": "newcoach@example.com", "password": "password123" }`                   |
| GET    | `api/user/{id}`  | Retrieves a specific user by ID                                             | `JwtMiddlewareUser`         | N/A                                                                                                    |
| GET    | `api/user/`      | Retrieves all users                                                         | `JwtMiddlewareUser`         | N/A                                                                                                    |
| POST   | `api/event/join` | Allows a user to join an event                                              | `JwtMiddlewareUser`         | `{ "event_name": "Yoga Class", "user_email": "user@example.com" }`                                     |
| POST   | `api/event/`     | Allows a coach to create an event                                           | `JwtMiddlewareCoach`        | `{ "name": "Yoga Class", "location": "Park", "cost": 10, event_expenses: 1200, coach_expenses: 1234 }` |
| GET    | `api/event/{id}` | Retrieves a specific event by name                                          | None                        | N/A                                                                                                    |
| GET    | `api/event/`     | Retrieves all events                                                        | None                        | N/A                                                                                                    |
| DELETE | `api/user/{id}`  | Deletes the user of the corresponding id                                    | None                        | N/A                                                                                                    |
| POST   | `api/auth/session` | Returns contents of cookie                                                  | None                        | N/A                                                                                                    |
| POST   | `api/auth/logout` | Logs out the user                                                           | None                        | N/A                                                                                                    |
| GET    | `api/finance` | Retrieves all of the yearly financial data                                  | `JwtMiddlewareCoach` | N/A                                                                                                    | 
| GET    | `api/finance/{year}` | Retrieves all of the financial data for a specific year                     | `JwtMiddlewareCoach` | N/A                                                                                                    |
| GET    | `api/finance/{year}/{month}` | Retrieves all of the financial data for a specific month in a specific year | `JwtMiddlewareCoach` | N/A                                                                                                    |
| PUT    | `api/user/{id}/pay` | Allows a user to pay                                                        | `JwtMiddlewareUser` | `{ "amount": int, "email": "example@example.com", "in_advance_payment": bool }`                        |
| GET    | `api/coach/owed/{id}`     | Retrieves the amount owed to a specific coach                                      | `JwtMiddlewareCoach`            | N/A                                                                                                    |

## Notes

- The `/{id}`, `/{year}` and `/{month}` in the endpoints are placeholders for dynamic values specific to each request.
- `Required JSON Structure` applies only to POST requests. Ensure the request body matches the structure for successful API calls.
- Middleware descriptions:
    - `JwtMiddlewareUser`: Checks for a valid JWT token for a user.
    - `JwtMiddlewareCoach`: Checks for a valid JWT token for a coach.
- Ensure to replace placeholder values like `user@example.com`, `password123`, etc., with actual information when making requests.


