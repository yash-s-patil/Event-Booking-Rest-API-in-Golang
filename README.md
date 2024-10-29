# Event-Booking-Rest-API-in-Golang
A Go-powered "Event Booking" REST API that allows users to browse, create, update, and delete events, as well as manage user registrations and authentication.

## API Endpoints
![API Endpoints](https://raw.githubusercontent.com/yash-s-patil/Event-Booking-Rest-API-in-Golang/main/restapi.png)

- GET /events
  - Retrieve a list of available events.
  - Authentication: Not required.

- GET /events/{id}
  - Retrieve details of a specific event by its ID.
  - Authentication: Not required.
 
- POST /events
  - Create a new bookable event.
  - Authentication: Required.
 
- PUT /events/{id}
  - Update an existing event.
  - Authentication: Required (only by the event creator).
 
- DELETE /events/{id}
  - Delete an existing event.
  - Authentication: Required (only by the event creator).
 
### User Management

- POST /signup
  - Register a new user.
  - Authentication: Not required.

- POST /login
  - Authenticate a user and obtain an authentication token (JWT).
  - Authentication: Not required.
 
### Event Registration

- POST /events/{id}/register
  - Register a user for an event.
  - Authentication: Required.
 
- DELETE /events/{id}/register
  - Cancel a user's registration for an event.
  - Authentication: Required.
 
## Authentication

The API uses JSON Web Tokens (JWT) for authentication. Users must include the token in the Authorization header of requests to protected endpoints.

Example:
```
Authorization: Bearer <your-token>
```

## Setup And Installation

### Install Dependencies
- Clone the repository
```
https://github.com/yash-s-patil/Event-Booking-Rest-API-in-Golang.git
cd Event-Booking-Rest-API-in-Golang
```

### Install required packages:
- Gin (HTTP web framework written in Go)
```
go get -u github.com/gin-gonic/gin
```
- SQLite driver for Go 
```
go get modernc.org/sqlite
```
- bcrypt for password hashing
```
go get -u golang.org/x/crypto
```
- JSON Web Tokens for Go
```
go get -u github.com/golang-jwt/jwt/v5
```
Set up environment variables (e.g., database configuration, JWT secret).

- Run the application
```
go run main.go
```

## Setup And Installation (Alternative - Using Dockerfile)
- Clone the repository
```
git clone https://github.com/yash-s-patil/Event-Booking-Rest-API-in-Golang.git
cd Event-Booking-Rest-API-in-Golang
```

- Build the Docker Image
```
docker build -t event-booking-api .
```

- Run the Docker Container
```
docker run -d -p 8080:8080 event-booking-api
```

### Testing the API
You can use an HTTP client like VSCode's "REST Client" extension or Postman to test the API endpoints. Here's how you can test the different endpoints:

- Signup
```
POST http://localhost:8080/signup
Content-Type: application/json

{
    "email": "test3@example.com",
    "password": "test"
}
```

- Login
```
POST http://localhost:8080/login
Content-Type: application/json

{
    "email": "test2@example.com",
    "password": "test2"
}
```
Copy the returned token, you'll need it for the authenticated endpoints.

- Create Event
```
POST http://localhost:8080/events
Content-Type: application/json
Authorization: your_auth_token

{
    "name": "Test event 1",
    "description": "test Event 1",
    "location": "test location 1",
    "dateTime": "2024-01-01T15:30:00.000Z"
}
```

- Get Events
```
GET http://localhost:8080/events
```

- Get Event by ID
```
GET http://localhost:8080/events/1
```

- Update Event
```
PUT http://localhost:8080/events/1
Content-Type: application/json
Authorization: your_auth_token

{
    "name": "Event Update",
    "description": "update Event",
    "location": "Updated Location",
    "dateTime": "2024-02-01T15:30:00.00Z"
}
```

- Delete Event
```
DELETE http://localhost:8080/events/1
Authorization: your_auth_token
```

- Register for Event
```
POST http://localhost:8080/events/1/register
Authorization: your_auth_token
```

- Cancel Registration
```
DELETE http://localhost:8080/events/1/register
Authorization: your_auth_token
```



