package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID             string    `json:"id" gorm:"type:varchar(50);primaryKey"`
	Name           string    `json:"name" gorm:"type:varchar(50);not null"`
	Surname        string    `json:"surname" gorm:"type:varchar(50);not null"`
	Phone          string    `json:"phone" gorm:"type:varchar(50);unique;not null"`
	IdentityNumber string    `json:"identityNumber" gorm:"type:varchar(11);unique;not null;check:length(11)"`
	Email          string    `json:"email" gorm:"type:varchar(50);unique;not null;"`
	Password       string    `json:"password" gorm:"not null"`
	SessionActive  bool      `json:"session_active" gorm:"default:false"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime;not null"`
	DeletedAt      time.Time `json:"deleted_at" gorm:"default:null"`
	// Relations
	OwnedApartments  []Apartment `json:"ownedApartments" gorm:"foreignKey:OwnerID;references:ID"`
	RentedApartments []Apartment `json:"rentedApartments" gorm:"foreignKey:HirerID;references:ID"`
	SentMessages     []Message   `json:"sentMessages" gorm:"foreignKey:UserID;references:ID"`
}
type Manager struct {
	gorm.Model
	UserID string `json:"user_id" gorm:"type:varchar(50);primaryKey"`
	User   *User  `json:"user" gorm:"foreignKey:UserID;references:ID"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New().String()
	return
}
