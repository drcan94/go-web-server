package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

// Global database connection
var conn *pgx.Conn

// User struct for form data
type User struct {
	ID        int
	FullName  string
	Email     string
	Password  string
	CreatedAt string
}

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to database
	connStr := os.Getenv("DATABASE_URL")
	conn, err = pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
	defer conn.Close(context.Background())

	// Create users table with updated fields
	_, err = conn.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			full_name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}

	// Define routes
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/register", handleRegister)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start server
	fmt.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	var dbPassword string
	err := conn.QueryRow(context.Background(),
		"SELECT password FROM users WHERE email = $1",
		email).Scan(&dbPassword)

	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if dbPassword == password { // Note: In production, use proper password hashing!
		fmt.Fprintf(w, "Login successful!")
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fullName := r.FormValue("fullname")
	email := r.FormValue("email")
	password := r.FormValue("password")

	_, err := conn.Exec(context.Background(),
		"INSERT INTO users (full_name, email, password) VALUES ($1, $2, $3)",
		fullName, email, password)

	if err != nil {
		http.Error(w, "Email already exists", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Registration successful!")
}