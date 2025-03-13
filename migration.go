package main

import (
	"akastra-mobile-api/src/infrastructure/database"
	"akastra-mobile-api/src/infrastructure/database/models/users"
	"fmt"
	"log"

	"gorm.io/gorm"
)

func main() {
	db, err := database.ConnectDB() 
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run the migration
	fmt.Println("Running database migration...")
	err = migrate(db)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	fmt.Println("Migration completed successfully.")
}

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&users.User{},
		&users.UserRole{},
	)
}
