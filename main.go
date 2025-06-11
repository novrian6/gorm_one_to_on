// main.go
package main

import (
	"gorm_one_to_one/config"
	"gorm_one_to_one/models"
)

func main() {
	// Connect to the database
	config.ConnectDatabase()

	// Migrate the schema for Author and Country tables
	models.MigrateTables()

	// Seed an author with a country
	models.CreateAuthorWithCountry("J.K. Rowling", "United Kingdom")

	// Retrieve and print the author with their country
	models.GetAuthorWithCountry(1)
}
