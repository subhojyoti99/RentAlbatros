package model

import (
	"ToLet/database"

	"gorm.io/gorm"
)

//	type Room struct {
//		gorm.Model
//		ID             uint   `gorm:"primary_key;auto_increment" json:"id"`
//		User_ID        uint   `json:"user_id"`
//		Pg_Admin_ID    uint   `json:"pg_admin_id"`
//		RoomType       string `grom:"required" json:"room_type"`
//		Location       string `grom:"required" json:"room_location"`
//		Cost           int    `grom:"required" json:"room_cost"`
//		Security_Cost  int    `grom:"required" json:"room_security_cost"`
//		Capacity       string `grom:"required" json:"room_capacity"`
//		Year           int    `grom:"required" json:"room_age"`
//		Feature        string `grom:"required" json:"room_feature"`
//		Description    string `grom:"required" json:"room_description"`
//		Status         string `grom:"required" json:"room_status"`
//		AvailableRooms uint   `grom:"required" json:"available_room"`
//		Rented_ID      uint   `json:"rental_id"`
//		// CreatedAt   time.Time `json:"created_at"`
//		Customers []User    `gorm:"foreignKey:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"customers"`
//		Admins    User      `gorm:"foreignKey:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"admin"`
//		Rentals   []Rental  `gorm:"foreignKey:Room_ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"rentals"`
//		Payments  []Payment `gorm:"foreignKey:Room_ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"payments"`
//	}
// type Room struct {
// 	gorm.Model
// 	UserID         uint   `json:"user_id"`
// 	PgAdminID      uint   `json:"pg_admin_id"`
// 	RoomType       string `gorm:"required" json:"room_type"`
// 	Location       string `gorm:"required" json:"room_location"`
// 	Cost           int    `gorm:"required" json:"room_cost"`
// 	SecurityCost   int    `gorm:"required" json:"room_security_cost"`
// 	Capacity       string `gorm:"required" json:"room_capacity"`
// 	Year           int    `gorm:"required" json:"room_age"`
// 	Feature        string `gorm:"required" json:"room_feature"`
// 	Description    string `gorm:"required" json:"room_description"`
// 	Status         string `gorm:"required" json:"room_status"`
// 	AvailableRooms uint   `gorm:"required" json:"available_room"`
// 	RentedID       uint   `json:"rentalid"`
// }

type Room struct {
	gorm.Model
	OwnerID uint `json:"owner_id"`
	// CustomerID     uint   `json:"customer_id"`
	RoomType       string `gorm:"required" json:"room_type"`   //like pg or independent room
	Gender         string `gorm:"required" json:"gender_type"` // girls or boys
	Location       string `gorm:"required" json:"room_location"`
	Cost           int    `gorm:"required" json:"room_cost"`
	SecurityCost   int    `gorm:"required" json:"room_security_cost"`
	Capacity       string `gorm:"required" json:"room_capacity"`
	Year           int    `gorm:"required" json:"room_age"`     //age of the room or pg
	Feature        string `gorm:"required" json:"room_feature"` // 1bhk or 2bhk
	Description    string `gorm:"required" json:"room_description"`
	Status         string `gorm:"required" json:"room_status"` // ready to move or not
	AvailableRooms uint   `gorm:"required" json:"available_room"`
	// RentedID       uint   `json:"rental_id"`
	Owner     User   `gorm:"foreignKey:OwnerID" json:"owner"`
	Customers []User `gorm:"many2many:room_customers:customer_rooms" json:"customers"`
	Rentals   []Rental
}

func (r *Room) Save() (*Room, error) {
	err := database.DB.Create(&r).Error
	if err != nil {
		return &Room{}, err
	}
	return r, err
}

// func FindRooms()
