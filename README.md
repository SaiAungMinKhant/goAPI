# goAPI

A lightweight Go API service that provides account-related functionality, including coin balance queries.

## Features

- RESTful API with Chi router
- Middleware for authorization
- Structured error handling
- JSON response formatting
- Logging with Logrus

## Project Structure

```
goAPI/
├── api/              # API models and error handling
├── cmd/              # Entry points for the application
│   └── api/          # Main API server
├── internal/         # Private application code
│   ├── handlers/     # HTTP request handlers
│   ├── middleware/   # Custom middleware
│   └── tools/        # Utility functions
```

## Getting Started

### Prerequisites

- Go 1.24 or higher

### Installation

1. Clone the repository
   ```
   git clone https://github.com/SaiAungMinKhant/goAPI.git
   ```

2. Navigate to the project directory
   ```
   cd goAPI
   ```

3. Install dependencies
   ```
   go mod download
   ```

### Running the API

1. Start the API server
   ```
   go run cmd/api/main.go
   ```

2. The API will be available at `http://localhost:8000`

## API Endpoints

### GET /account/coins

Get the coin balance for a user account.

Query Parameters:
- `username`: The username of the account

Response:
```json
{
  "code": 200,
  "balance": 100
}
```

## Error Handling

The API returns structured error responses:

```json
{
  "code": 400,
  "message": "Error message"
}
```

## Dependencies

- [Chi](https://github.com/go-chi/chi) - Lightweight, idiomatic and composable router
- [Logrus](https://github.com/sirupsen/logrus) - Structured logger
- [Gorilla Schema](https://github.com/gorilla/schema) - Form/query parameter decoder

## License

This project is licensed under the MIT License - see the LICENSE file for details.