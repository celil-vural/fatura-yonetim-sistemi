package apartmentDtos

import "fatura-yonetim-sistemi/entity/models"

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
	ID      string `json:"id"`
	HirerID string `json:"hirerId" gorm:"type:varchar(50)"`
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
