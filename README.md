# Simple Web Login System

A minimalist web application built with Go that implements a basic user authentication system with login and registration functionality. The application uses PostgreSQL for data storage and serves a clean, responsive UI.

## Features

- User Registration
- User Login
- Responsive UI
- PostgreSQL Database Integration
- Static File Serving
- Environment Variable Configuration

## Prerequisites

- Go 1.24 or higher
- PostgreSQL database
- Git

## Installation

1. Clone the repository:
```bash
git clone [your-repository-url]
cd [repository-name]
```

2. Install dependencies:
```bash
go mod download
```

3. Create a `.env` file in the root directory with the following content:
```env
DATABASE_URL=postgres://username:password@localhost:5432/dbname
```
Replace `username`, `password`, and `dbname` with your PostgreSQL credentials.

## Project Structure

```
.
├── main.go           # Main application file
├── static/           # Static files directory
│   └── style.css     # CSS styles
├── templates/        # HTML templates
│   └── index.html    # Main page template
├── go.mod           # Go modules file
├── go.sum           # Go modules checksum
└── .env             # Environment variables (create this)
```

## Running the Application

1. Make sure PostgreSQL is running and accessible with the credentials in your `.env` file.

2. Start the application:
```bash
go run main.go
```

3. Open your web browser and navigate to:
```
http://localhost:8080
```

## Database Schema

The application automatically creates a `users` table with the following structure:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## Security Notes

This is a basic implementation and should not be used in production without the following improvements:

- Password hashing implementation
- Session management
- CSRF protection
- HTTPS implementation
- Input validation
- Rate limiting

## Dependencies

- github.com/jackc/pgx/v5 - PostgreSQL driver
- github.com/joho/godotenv - Environment variable management
- github.com/lib/pq - PostgreSQL library

## Contributing

Feel free to submit issues and enhancement requests.

## License

[Your chosen license] 