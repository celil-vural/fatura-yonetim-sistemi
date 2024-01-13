package models

import "gorm.io/gorm"

type Apartment struct {
	gorm.Model
	ApartmentNumber     string `json:"apartmentNumber" gorm:"type:varchar(50);not null;unique"`
	ApartmentType       string `json:"apartmentType" gorm:"type:varchar(50);not null"`
	ApartmentFloor      string `json:"apartmentFloor" gorm:"type:varchar(50);not null"`
	ApartmentBlock      rune   `json:"apartmentBlock" gorm:"type:varchar(1);not null"`
	ApartmentStatus     bool   `json:"apartmentStatus" gorm:"type:bit;not null"`
	ApartmentRoomsCount int    `json:"apartmentRoomsCount" gorm:"type:int;not null"`
	ApartmentArea       int    `json:"apartmentArea" gorm:"type:int;not null"`
	// Relations
	OwnerID  string    `json:"ownerId" gorm:"type:varchar(50);not null"`
	Owner    *User     `json:"owner" gorm:"foreignKey:OwnerID;references:ID"`
	HirerID  string    `json:"hirerId" gorm:"type:varchar(50)"`               // nullable
	Hirer    *User     `json:"hirer" gorm:"foreignKey:HirerID;references:ID"` // nullable
	Dues     []Dues    `json:"dues" gorm:"foreignKey:ApartmentID;references:ID"`
	Invoices []Invoice `json:"invoices" gorm:"foreignKey:ApartmentID;references:ID"`
}
