package model

import (
	"ToLet/database"
	"time"

	"gorm.io/gorm"
)

//	type Payment struct {
//		gorm.Model
//		ID          uint      `gorm:"primary_key;auto_increment" json:"id"`
//		Rental_ID   uint      `json:"rental_id"`
//		User_ID     uint      `json:"user_id"`
//		Room_ID     uint      `json:"room_id"`
//		Amount      uint      `grom:"required" json:"room_amount"`
//		PaymentDate time.Time `grom:"required" json:"payment_date"`
//		Rental      Rental    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"rental"`
//		User        User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"user"`
//		Room        Room      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"room"`
//	}
type Payment struct {
	gorm.Model
	UserID      uint      `json:"user_id"`
	RentalID    uint      `json:"rental_id"`
	RoomID      uint      `json:"room_id"`
	Amount      uint      `grom:"required" json:"room_amount"`
	PaymentDate time.Time `grom:"required" json:"payment_date"`
	User        User
	Rental      Rental `gorm:"foreignKey:RentalID"`
}

func (p *Payment) Save() (*Payment, error) {
	err := database.DB.Create(&p).Error
	if err != nil {
		return &Payment{}, err
	}
	return p, err
}
