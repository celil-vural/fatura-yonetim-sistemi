package apartmentDtos

import (
	"fatura-yonetim-sistemi/entity/models"
	"gorm.io/gorm"
	"time"
)

type ApartmentDtoForCreate struct {
	ApartmentNumber     string `json:"apartmentNumber"`
	ApartmentType       string `json:"apartmentType"`
	ApartmentFloor      string `json:"apartmentFloor"`
	ApartmentBlock      rune   `json:"apartmentBlock"`
	ApartmentStatus     bool   `json:"apartmentStatus"`
	ApartmentRoomsCount int    `json:"apartmentRoomsCount"`
	ApartmentArea       int    `json:"apartmentArea"`
	OwnerID             string `json:"ownerId" gorm:"type:varchar(50);not null"`
	HirerID             string `json:"hirerId" gorm:"type:varchar(50)"`
}
type ApartmentDtoForUpdateHirer struct {
	ID      uint   `json:"id"`
	HirerID string `json:"hirerId" gorm:"type:varchar(50)"`
}
type ApartmentDtoForManagerGet struct {
	ID                  uint           `json:"id"`
	ApartmentNumber     string         `json:"apartmentNumber"`
	ApartmentType       string         `json:"apartmentType"`
	ApartmentFloor      string         `json:"apartmentFloor"`
	ApartmentBlock      rune           `json:"apartmentBlock"`
	ApartmentStatus     bool           `json:"apartmentStatus"`
	ApartmentRoomsCount int            `json:"apartmentRoomsCount"`
	ApartmentArea       int            `json:"apartmentArea"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `json:"deleted_at"`
	OwnerID             string         `json:"ownerId" `
	HirerID             string         `json:"hirerId"`
}

func (a ApartmentDtoForManagerGet) FromApartment(apartment *models.Apartment) {
	a.ID = apartment.ID
	a.ApartmentNumber = apartment.ApartmentNumber
	a.ApartmentType = apartment.ApartmentType
	a.ApartmentFloor = apartment.ApartmentFloor
	a.ApartmentBlock = apartment.ApartmentBlock
	a.ApartmentStatus = apartment.ApartmentStatus
	a.ApartmentRoomsCount = apartment.ApartmentRoomsCount
	a.ApartmentArea = apartment.ApartmentArea
	a.CreatedAt = apartment.CreatedAt
	a.UpdatedAt = apartment.UpdatedAt
	a.DeletedAt = apartment.DeletedAt
	a.OwnerID = apartment.OwnerID
	a.HirerID = apartment.HirerID
}
func (a ApartmentDtoForCreate) ToApartment() *models.Apartment {
	var apartment models.Apartment
	apartment.ApartmentNumber = a.ApartmentNumber
	apartment.ApartmentType = a.ApartmentType
	apartment.ApartmentFloor = a.ApartmentFloor
	apartment.ApartmentBlock = a.ApartmentBlock
	apartment.ApartmentStatus = a.ApartmentStatus
	apartment.ApartmentRoomsCount = a.ApartmentRoomsCount
	apartment.ApartmentArea = a.ApartmentArea
	apartment.OwnerID = a.OwnerID
	apartment.HirerID = a.HirerID
	return &apartment
}
