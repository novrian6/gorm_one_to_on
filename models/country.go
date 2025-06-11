// models/country.go
package models

type Country struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:100;unique"`
}
