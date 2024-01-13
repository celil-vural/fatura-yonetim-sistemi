package models

import "gorm.io/gorm"

type Invoice struct {
	gorm.Model
	InvoiceTypeID uint `json:"invoiceTypeId" gorm:"not null"`
	Amount        int  `json:"amount" gorm:"type:int;not null"`
	IsPaid        bool `json:"isPaid" gorm:"type:bit;not null"`
	// Relations
	ApartmentID uint       `json:"apartmentId" gorm:"not null"`
	Apartment   *Apartment `json:"apartment" gorm:"foreignKey:ApartmentID;references:ID"`
}
type InvoiceType struct {
	gorm.Model
	Name string `json:"name" gorm:"type:varchar(50);not null"`
}
type InvoicePayment struct {
	gorm.Model
	InvoiceID uint `json:"invoiceId" gorm:"not null"`
	// Relations
	Invoice *Invoice `json:"invoice" gorm:"foreignKey:InvoiceID;references:ID"`
}
