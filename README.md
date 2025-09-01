# Event Booking REST API (Golang)

A practice REST API for event booking, built with Go and Gin. It demonstrates user authentication, event management, and registration workflows using a SQLite database.

## Features

- **User Authentication**: Signup and login endpoints with JWT-based authentication.
- **Event Management**: Create, update, delete, and view events.
- **Event Registration**: Register and cancel registration for events.
- **Protected Routes**: Only authenticated users can create, update, delete, or register for events.
- **Password Hashing**: Secure password storage using bcrypt.
- **SQLite Database**: Lightweight, file-based database for easy setup.

## API Endpoints

### Auth

- `POST /signup` — Register a new user.
- `POST /login` — Login and receive a JWT token.

### Events

- `GET /events` — List all events.
- `GET /events/:id` — Get details of a single event.
- `POST /events` — Create a new event (auth required).
- `PUT /events/:id` — Update an event (auth required).
- `DELETE /events/:id` — Delete an event (auth required).

### Registration

- `POST /events/:id/register` — Register for an event (auth required).
- `DELETE /events/:id/register` — Cancel registration (auth required).

## Project Structure

```
event-booking-rest-api-golang/
├── api-test/         # HTTP request samples for testing endpoints
├── database/         # Database connection and table creation
├── middleware/       # JWT authentication middleware
├── models/           # Data models for users, events, registrations
├── routes/           # API route handlers
├── utils/            # Utility functions (JWT, password hashing)
├── main.go           # Application entry point
├── go.mod, go.sum    # Go dependencies
└── README.md         # Project documentation
```

## How It Works

- **Authentication**: JWT tokens are generated on login and required for protected endpoints. The middleware validates tokens and sets the user context.
- **Database**: Tables for users, events, and registrations are created automatically on startup.
- **Event Operations**: Authenticated users can manage events and register/cancel registrations.

## Getting Started

1. **Clone the repo**

   ```sh
   git clone https://github.com/yourusername/event-booking-rest-api-golang.git
   cd event-booking-rest-api-golang
   ```

2. **Install dependencies**

   ```sh
   go mod tidy
   ```

3. **Run the server**

   ```sh
   go run main.go
   ```

4. **Test endpoints**
   - Use the sample requests in `api-test/` with VS Code REST Client or Postman.

## Example Request

```http
POST http://localhost:8080/signup
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "yourpassword"
}
```

## Security

- Passwords are hashed before storing in the database.
- JWT tokens are signed and validated for all protected routes.

## Contributing

Feel free to fork and submit pull requests for improvements or new features!
