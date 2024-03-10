package model

import (
	"ToLet/database"
	"time"

	"gorm.io/gorm"
)

//	type Rental struct {
//		gorm.Model
//		// ID         uint      `gorm:"primary_key;auto_increment" json:"id"`
//		User_ID    uint      `json:"user_id"`
//		Room_ID    string    `json:"room_id"`
//		Start_Date time.Time `gorm:"required" json:"start_date"`
//		End_Date   time.Time `gorm:"required" json:"end_date"`
//		Total_Cost int       `gorm:"required" json:"total_cost"`
//		Payment_ID uint      `json:"payment_id"`
//		// CreatedAt  time.Time `json:"created_at"`
//		Room     Room      `gorm:"foreignkey:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"room"`
//		User     User      `gorm:"foreignkey:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"user"`
//		Payments []Payment `gorm:"foreignkey:Rental_ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"payments"`
//	}
type Rental struct {
	gorm.Model
	UserID     uint      `json:"user_id"`
	RoomID     string    `json:"room_id"`
	Start_Date time.Time `gorm:"required" json:"start_date"`
	End_Date   time.Time `gorm:"required" json:"end_date"`
	Status     string    `json:"status"`
	Total_Cost int       `gorm:"required" json:"total_cost"`
	User       User
	Room       Room
}

func (r *Rental) Save() (*Rental, error) {
	err := database.DB.Create(&r).Error
	if err != nil {
		return &Rental{}, err
	}
	return r, err
}
