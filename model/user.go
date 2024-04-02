package model

import (
	"ToLet/database"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//	type User struct {
//		gorm.Model
//		ID            uint   `gorm:"primary_key;auto_increment" json:"id"`
//		FirstName     string `gorm:"required" json:"first_name"`
//		LastName      string `gorm:"required" json:"last_name"`
//		Gender        string `gorm:"required" json:"gender"`
//		Email         string `gorm:"required" json:"email"`
//		ContactNumber string `gorm:"required" json:"contact_number"`
//		Password      string `gorm:"required" json:"password"`
//		Role          string `gorm:"required" json:"role"`
//		RoleID        uint   `gorm:"not null;DEFAULT:3" json:"role_id"`
//		ValidKey      string `json:"valid_key"`
//		// CreatedAt      time.Time `json:"created_at"`
//		Rentals        []Rental  `gorm:"foreignKey:User_ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"rentals"`
//		Payments       []Payment `gorm:"foreignKey:User_ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"payments"`
//		Customer_Rooms []Room    `gorm:"foreignKey:User_ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"customer_rooms"`
//		Pg_Admin_Room  []Room    `gorm:"foreignKey:Pg_Admin_ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"admin_rooms"`
//	}
// type User struct {
// 	ID             uint   `gorm:"primary_key; auto_increment"`
// 	FirstName      string `gorm:"required" json:"first_name"`
// 	LastName       string `gorm:"required" json:"last_name"`
// 	Gender         string `gorm:"required" json:"gender"`
// 	Email          string `gorm:"required" json:"email"`
// 	ContactNumber  string `gorm:"required" json:"contact_number"`
// 	Password       string `gorm:"required" json:"password"`
// 	Role           string `gorm:"required" json:"role"`
// 	RoleID         uint   `gorm:"not null;DEFAULT:3" json:"role_id"`
// 	ValidKey       string `json:"valid_key"`
// 	Customer_Rooms []Room `gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"customer_rooms"`
// }

type User struct {
	gorm.Model
	ID                  uint   `gorm:"primary_key; auto_increment"`
	FirstName           string `gorm:"required" json:"first_name"`
	LastName            string `gorm:"required" json:"last_name"`
	Gender              string `gorm:"required" json:"gender"`
	Email               string `gorm:"unique;not null" json:"email"`
	ContactNumber       string `gorm:"required" json:"contact_number"`
	Password            string `gorm:"required" json:"password"`
	Role                string `gorm:"required" json:"role"`
	RoleID              uint   `gorm:"not null;DEFAULT:3" json:"role_id"`
	ValidKey            string `json:"valid_key"`
	Address             string `json:"address"`
	ImgURL              string `json:"img_url"`
	registration_source string `gorm:"not null"`
	Rentals             []Rental
	Payments            []Payment
	OwnedRooms          []Room `gorm:"foreignKey:OwnerID"`
}

func (u *User) Save() (*User, error) {
	err := database.DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, err
}

// Generate encrypted password
func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	return nil
}

// Get all users
func FindUsers(User *[]User) (err error) {
	err = database.DB.Find(User).Error
	if err != nil {
		return err
	}
	return nil
}

// Get user by email
func FindUserByEmail(email string) (User, error) {
	var user User
	err := database.DB.Where("email=?", email).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// Get user by contact
func FindUserByContact(contact_number string) (User, error) {
	var user User
	err := database.DB.Where("contact_number=?", contact_number).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// Validate user password
func (user *User) ValidateUserPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

// Get user by id
func FindUserById(id uint) (User, error) {
	var user User
	err := database.DB.Where("id=?", id).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// Get user by id
func FindUser(user *User, id int) (err error) {
	err = database.DB.Where("id = ?", id).First(user).Error
	if err != nil {
		return err
	}
	return nil
}

// Update user
func UpdateUser(user *User) (err error) {
	err = database.DB.Omit("password").Save(user).Error
	if err != nil {
		return err
	}
	return nil
}
