package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID       uint `gorm:"primaryKey"`
	Username string
	Email    string
}

// connects to PostgreSQL database and returns the database connection
func connectToPostgreSQL() (*gorm.DB, error) {
	// PostgreSQL connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_DB_NAME"),
		os.Getenv("PG_DB_SSLMODE"))

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// creates a new database user
func createUser(db *gorm.DB, user User) error {
	result := db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// retrieves a user from the database
func getUser(db *gorm.DB, id uint) (*User, error) {
	var user User
	result := db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func main() {

	db, err := connectToPostgreSQL()
	if err != nil {
		fmt.Println("Error connecting to PostgreSQL")
		return
	}
	fmt.Println("Connected to PostgreSQL:", db.Name())

	// Perform database migration
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}

	// Create a new user
	newUser := &User{Username: "john_doe", Email: "john.doe@example.com"}
	err = createUser(db, *newUser)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("User created successfully")

	// Get a user
	user, err := getUser(db, 1)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("User:", user)
}
