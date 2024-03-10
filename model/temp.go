package model

// type User struct {
// 	gorm.Model
// 	ID            uint   `gorm:"primary_key;auto_increment" json:"id"`
// 	FirstName     string `gorm:"required" json:"first_name"`
// 	LastName      string `gorm:"required" json:"last_name"`
// 	Gender        string `gorm:"required" json:"gender"`
// 	Email         string `gorm:"required" json:"email"`
// 	ContactNumber string `gorm:"required" json:"contact_number"`
// 	Password      string `gorm:"required" json:"password"`
// 	Role          string `gorm:"required" json:"role"`
// 	RoleID        uint   `gorm:"not null;DEFAULT:3" json:"role_id"`
// 	ValidKey      string `json:"valid_key"`
// 	// CreatedAt      time.Time `json:"created_at"`
// 	Rentals        []Rental  `gorm:"foreignKey:User_ID" json:"rentals"`
// 	Payments       []Payment `gorm:"foreignKey:User_ID" json:"payments"`
// 	Customer_Rooms []Room    `gorm:"foreignKey:User_ID" json:"customer_rooms"`
// 	Pg_Admin_Room  []Room    `gorm:"foreignKey:Pg_Admin_ID" json:"admin_rooms"`
// }

// type Room struct {
// 	gorm.Model
// 	ID             uint   `gorm:"primary_key;auto_increment" json:"id"`
// 	User_ID        uint   `json:"user_id"`
// 	Pg_Admin_ID    uint   `json:"pg_admin_id"`
// 	RoomType       string `grom:"required" json:"room_type"`
// 	Location       string `grom:"required" json:"room_location"`
// 	Cost           int    `grom:"required" json:"room_cost"`
// 	Security_Cost  int    `grom:"required" json:"room_security_cost"`
// 	Capacity       string `grom:"required" json:"room_capacity"`
// 	Year           int    `grom:"required" json:"room_age"`
// 	Feature        string `grom:"required" json:"room_feature"`
// 	Description    string `grom:"required" json:"room_description"`
// 	Status         string `grom:"required" json:"room_status"`
// 	AvailableRooms uint   `grom:"required" json:"available_room"`
// 	Rented_ID      uint   `json:"rental_id"`
// 	// CreatedAt   time.Time `json:"created_at"`
// 	Customers []User    `gorm:"foreignkey:ID;" json:"customers"`
// 	Admins    User      `gorm:"foreignkey:ID;" json:"admin"`
// 	Rentals   []Rental  `gorm:"foreignkey:Room_ID;" json:"rentals"`
// 	Payments  []Payment `gorm:"foreignKey:Room_ID" json:"payments"`
// }

// type Rental struct {
// 	gorm.Model
// 	ID         uint      `gorm:"primary_key;auto_increment" json:"id"`
// 	User_ID    uint      `json:"user_id"`
// 	Room_ID    string    `json:"room_id"`
// 	Start_Date time.Time `gorm:"required" json:"start_date"`
// 	End_Date   time.Time `gorm:"required" json:"end_date"`
// 	Total_Cost int       `gorm:"required" json:"total_cost"`
// 	Payment_ID uint      `json:"payment_id"`
// 	// CreatedAt  time.Time `json:"created_at"`
// 	Room     Room      `gorm:"foreignkey:ID" json:"room"`
// 	User     User      `gorm:"foreignkey:ID" json:"user"`
// 	Payments []Payment `gorm:"foreignkey:Rental_ID" json:"payments"`
// }

// type Payment struct {
// 	gorm.Model
// 	ID          uint      `gorm:"primary_key;auto_increment" json:"id"`
// 	Rental_ID   uint      `json:"rental_id"`
// 	User_ID     uint      `json:"user_id"`
// 	Room_ID     uint      `json:"room_id"`
// 	Amount      uint      `grom:"required" json:"room_amount"`
// 	PaymentDate time.Time `grom:"required" json:"payment_date"`
// 	Rental      Rental    `json:"rental"`
// 	User        User      `json:"user"`
// 	Room        Room      `json:"room"`
// }

// 1. Room 1:
//    - RoomType: "PG"
//    - Gender: "Girls"
//    - Location: "ABC Street, City"
//    - Cost: 5000
//    - SecurityCost: 2000
//    - Capacity: "2 sharing"
//    - Year: 2
//    - Feature: "AC, Attached Bathroom"
//    - Description: "Spacious rooms available for girls"
//    - Status: "Ready to move"
//    - AvailableRooms: 3

// 2. Room 2:
//    - RoomType: "Independent"
//    - Gender: "Boys"
//    - Location: "XYZ Road, Town"
//    - Cost: 8000
//    - SecurityCost: 3000
//    - Capacity: "3 sharing"
//    - Year: 1
//    - Feature: "Balcony, Furnished"
//    - Description: "Newly constructed independent rooms"
//    - Status: "Under construction"
//    - AvailableRooms: 5

// 3. Room 3:
//    - RoomType: "PG"
//    - Gender: "Girls"
//    - Location: "DEF Avenue, Suburb"
//    - Cost: 4500
//    - SecurityCost: 1800
//    - Capacity: "4 sharing"
//    - Year: 3
//    - Feature: "WiFi, Security Guard"
//    - Description: "Well-maintained PG for girls"
//    - Status: "Available from next month"
//    - AvailableRooms: 2

// 4. Room 4:
//    - RoomType: "Independent"
//    - Gender: "Boys"
//    - Location: "PQR Lane, Village"
//    - Cost: 6000
//    - SecurityCost: 2500
//    - Capacity: "2 sharing"
//    - Year: 5
//    - Feature: "Parking, Garden"
//    - Description: "Peaceful independent rooms for boys"
//    - Status: "Vacant"
//    - AvailableRooms: 4

// 5. Room 5:
//    - RoomType: "PG"
//    - Gender: "Mixed"
//    - Location: "LMN Street, Downtown"
//    - Cost: 5500
//    - SecurityCost: 2200
//    - Capacity: "3 sharing"
//    - Year: 4
//    - Feature: "24/7 Water Supply"
//    - Description: "Comfortable PG available for mixed gender"
//    - Status: "Fully occupied"
//    - AvailableRooms: 0
