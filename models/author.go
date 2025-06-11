// models/author.go
package models

import (
	"fmt"
	"gorm_one_to_one/config"
	"log"
)

type Author struct {
	ID        uint    `gorm:"primaryKey"`
	Name      string  `gorm:"size:100"`
	CountryID uint    // Foreign key to link with Country
	Country   Country `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// MigrateTables automatically migrates Author and Country tables
func MigrateTables() {
	config.DB.AutoMigrate(&Country{}, &Author{})
}

// models/author.go (continued)

// CreateAuthorWithCountry seeds an author with a country association
func CreateAuthorWithCountry(name string, countryName string) {
	country := Country{Name: countryName}
	author := Author{Name: name, Country: country}
	config.DB.Create(&author)
}

// GetAuthorWithCountry fetches an author along with their country
func GetAuthorWithCountry(authorID uint) {
	var author Author
	if err := config.DB.Preload("Country").First(&author, authorID).Error; err != nil {
		log.Println("Error fetching author:", err)
		return
	}
	fmt.Printf("Author: %s, Country: %s\n", author.Name, author.Country.Name)
}
