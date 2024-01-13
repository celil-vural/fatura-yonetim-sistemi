package models

import "gorm.io/gorm"

type Dues struct {
	gorm.Model
	Amount int `json:"amount" gorm:"type:int;not null"`
	// Relations
	ApartmentID uint       `json:"apartmentId" gorm:"not null"`
	Apartment   *Apartment `json:"apartment" gorm:"foreignKey:ApartmentID;references:ID"`
}
type DuesPayment struct {
	gorm.Model
	DuesID uint `json:"duesId" gorm:"not null"`
	// Relations
	Dues *Dues `json:"dues" gorm:"foreignKey:DuesID;references:ID"`
}
