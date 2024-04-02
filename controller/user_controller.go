package controller

import (
	"ToLet/database"
	"ToLet/model"
	"ToLet/util"
	"errors"
	"fmt"
	"io"
	"net/http"

	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func Register(context *gin.Context) {
	ownerKey := os.Getenv("OWNER_VALID_KEY")
	adminKey := os.Getenv("ADMIN_VALID_KEY")

	tx := database.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var input model.User

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the role is valid
	if input.Role != "owner" && input.Role != "admin" && input.Role != "customer" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role specified"})
		return
	}

	// Set the role ID based on the input
	var roleID uint
	switch input.Role {
	case "owner":
		roleID = 1
	case "admin":
		roleID = 2
	case "customer":
		roleID = 3
	}

	owner := model.User{
		FirstName:     input.FirstName,
		LastName:      input.LastName,
		Gender:        input.Gender,
		Email:         input.Email,
		ContactNumber: input.ContactNumber,
		Password:      input.Password,
		Role:          input.Role,
		RoleID:        roleID,
		ValidKey:      input.ValidKey,
		Address:       input.Address,
	}

	admin := model.User{
		FirstName:     input.FirstName,
		LastName:      input.LastName,
		Gender:        input.Gender,
		Email:         input.Email,
		ContactNumber: input.ContactNumber,
		Password:      input.Password,
		Role:          input.Role,
		RoleID:        roleID,
		ValidKey:      input.ValidKey,
		Address:       input.Address,
	}

	customer := model.User{
		FirstName:     input.FirstName,
		LastName:      input.LastName,
		Gender:        input.Gender,
		Email:         input.Email,
		ContactNumber: input.ContactNumber,
		Password:      input.Password,
		Role:          input.Role,
		RoleID:        roleID,
		ValidKey:      input.ValidKey,
		Address:       input.Address,
	}

	existingEmail, _ := model.FindUserByEmail(input.Email)
	if input.Email == existingEmail.Email {
		context.JSON(http.StatusConflict, gin.H{"error": "This email is already registered."})
		return
	}

	existingContact, _ := model.FindUserByContact(input.ContactNumber)
	if input.ContactNumber == existingContact.ContactNumber {
		context.JSON(http.StatusConflict, gin.H{"error": "This contact number is already registered."})
		return
	}

	customer.Rentals = []model.Rental{}
	customer.Payments = []model.Payment{}
	customer.OwnedRooms = []model.Room{}

	for i := range admin.OwnedRooms {
		admin.OwnedRooms[i].OwnerID = admin.ID
	}

	// for i := range customer.Rentals {
	// 	customer.Rentals[i].User_ID = customer.ID
	// }
	// for j := range customer.Payments {
	// 	customer.Payments[j].User_ID = customer.ID
	// }
	// for k := range customer.CustomerRooms {
	// 	customer.CustomerRooms[k].UserID = customer.ID
	// }

	// database.DB.Preload("User").First(&admin.User_Rooms)

	// fmt.Println("hfiowehfiouhwe--iofhweiohf", ownerKey)
	if input.Role == "owner" && input.ValidKey == ownerKey {
		database.DB.Create(&owner)
		// context.JSON(http.StatusCreated, gin.H{"owner": owner})
		jwt, err := util.GenerateJWT(owner)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{"token": jwt, "user": owner, "message": "Successfully logged in"})
	} else if input.Role == "admin" && input.ValidKey == adminKey {
		database.DB.Create(&admin)
		// context.JSON(http.StatusCreated, gin.H{"admin": admin})
		jwt, err := util.GenerateJWT(admin)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{"token": jwt, "user": admin, "message": "Successfully logged in"})
	} else if input.Role == "customer" {
		database.DB.Create(&customer)
		// context.JSON(http.StatusCreated, gin.H{"customer": customer})
		jwt, err := util.GenerateJWT(customer)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{"token": jwt, "user": customer, "message": "Successfully logged in"})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Enter correct valid key.."})
	}

	tx.Commit()
}

// User Login
func Login(context *gin.Context) {

	tx := database.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var input model.Login

	if err := context.ShouldBindJSON(&input); err != nil {
		var errorMessage string
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			validationError := validationErrors[0]
			if validationError.Tag() == "required" {
				errorMessage = fmt.Sprintf("%s not provided", validationError.Field())
			}
		}
		context.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	user, err := model.FindUserByEmail(input.Email)

	if len(user.Email) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
		fmt.Println("err", err)
		return
	}

	err = user.ValidateUserPassword(input.Password)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt, err := util.GenerateJWT(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": jwt, "user": user, "message": "Successfully logged in"})

	tx.Commit()
}

// get all users
func GetUsers(context *gin.Context) {

	tx := database.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var user []model.User
	err := model.FindUsers(&user)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"users": user})

	tx.Commit()
}

// get user by id
func GetUserById(context *gin.Context) {

	tx := database.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	id, _ := strconv.Atoi(context.Param("id"))
	var user model.User
	// err := model.FindUser(&user, id)
	// if err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		context.AbortWithStatus(http.StatusNotFound)
	// 		return
	// 	}

	// 	context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	// 	return
	// }

	err := database.DB.Preload("OwnedRooms").Preload("OwnedRooms.Owner").First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.AbortWithStatus(http.StatusNotFound)
			return
		}
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// for i := range user.OwnedRooms {
	// 	user.OwnedRooms[i].OwnerID = user.ID
	// }

	// for i := range user.OwnedRooms {
	// 	// err := database.DB.Model(&user.OwnedRooms[i]).Preload("owner").Error
	// 	err := database.DB.Preload("owner").Find(&user.OwnedRooms[i]).Error
	// 	if err != nil {
	// 		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// }

	context.JSON(http.StatusOK, gin.H{"user": user})

	tx.Commit()
}

// update user
func UpdateUser(context *gin.Context) {

	tx := database.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	curUser := util.CurrentUser(context)

	// var user model.User
	// if err := context.ShouldBindJSON(&user); err != nil {
	// 	context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// id, _ := strconv.Atoi(context.Param("id"))

	user, err := model.FindUserById(curUser.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.AbortWithStatus(http.StatusNotFound)
			return
		}
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	fmt.Println("du2udud", user)

	if user.ID != curUser.ID {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied."})
		return
	}

	var input model.User
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = curUser.ID
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Gender = input.Gender
	user.Email = input.Email
	user.ContactNumber = input.ContactNumber
	user.Password = input.Password
	user.Role = input.Role
	user.ValidKey = input.ValidKey
	user.Address = input.Address

	fmt.Println("d+++u2udud", user)

	// err = model.UpdateUser(&user)
	// err = database.DB.Omit("password").Save(&user).Error

	// if err != nil {
	// 	tx.Rollback()
	// 	context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	// 	return
	// }
	context.JSON(http.StatusOK, user)

	tx.Commit()
}

func UpdateTheUser(context *gin.Context) {
	tx := database.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	curUser := util.CurrentUser(context)

	var input model.User
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := model.FindUserById(curUser.ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User not found."})
		return
	}

	if user.ID != curUser.ID {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied. You can only update your own details."})
		return
	}

	user.FirstName = input.FirstName
	user.LastName = input.LastName
	// user.Email = curUser.Email
	// user.ContactNumber = curUser.ContactNumber
	user.Gender = input.Gender
	// user.Role = curUser.Role
	// user.RoleID = curUser.RoleID
	user.Address = input.Address
	user.Gender = input.Gender

	err = database.DB.Omit("password", "email", "role", "role_id", "contact_number", "valid_key").Save(&user).Error
	if err != nil {
		tx.Rollback()
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"updatedUser": user})

	tx.Commit()

}

func CurrentUserGet(context *gin.Context) {
	tx := database.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	user := util.CurrentUser(context)
	context.JSON(http.StatusOK, gin.H{"User": user})
	tx.Commit()
}

// func AddImageURL(context *gin.Context) {
// 	tx := database.DB.Begin()

// 	defer func() {
// 		if r := recover(); r != nil {
// 			tx.Rollback()
// 		}
// 	}()
// 	curUser := util.CurrentUser(context)

// 	// Fetch the user from the database
// 	user, err := model.FindUserById(curUser.ID)
// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": "User not found."})
// 		return
// 	}

// 	// Check if the logged-in user is the owner of the profile
// 	if user.ID != curUser.ID {
// 		context.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied. You can only update your own details."})
// 		return
// 	}

// 	// Bind the JSON request to the input model
// 	var input struct {
// 		ImgURL string `json:"img_url"`
// 	}
// 	if err := context.ShouldBindJSON(&input); err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Update the user's image URL
// 	user.ImgURL = input.ImgURL

// 	// Save the updated user to the database
// 	if err := model.UpdateUser(&user); err != nil {
// 		tx.Rollback()
// 		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user."})
// 		return
// 	}

// 	tx.Commit()
// 	context.JSON(http.StatusOK, gin.H{"updatedUser": user})
// }

func UploadProfileImage(c *gin.Context) {
	tx := database.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Get the current user
	curUser := util.CurrentUser(c)

	// Fetch the user from the database
	user, err := model.FindUserById(curUser.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found."})
		return
	}

	// Check if the logged-in user is the owner of the profile
	if user.ID != curUser.ID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied. You can only update your own details."})
		return
	}

	// Retrieve the uploaded file from the form
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	// Create a unique filename for the uploaded file
	filename := strconv.Itoa(int(curUser.ID)) + "_" + file.Filename

	// Save the uploaded file to the server
	if err := c.SaveUploadedFile(file, "uploads/"+filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	fmt.Println("user", user)

	// Update the user's ImgURL with the filename or URL of the uploaded image
	// In this example, we assume the image is stored in a directory named "uploads"
	user.ImgURL = filename

	// Save the updated user to the database
	err = database.DB.Omit("password", "email", "role", "role_id", "contact_number", "valid_key").Save(&user).Error
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"updatedUser": user})

	// updatedUser, err := user.Save()
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user"})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "img_url": user.ImgURL})
	tx.Commit()
}

func UploadProfileImageNew(c *gin.Context) {
	tx := database.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Get the current user
	curUser := util.CurrentUser(c)

	// Fetch the user from the database
	user, err := model.FindUserById(curUser.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found."})
		return
	}

	// Check if the logged-in user is the owner of the profile
	if user.ID != curUser.ID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied. You can only update your own details."})
		return
	}

	// Parse the multipart form data
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer form.RemoveAll()

	// Retrieve the uploaded file from the form
	fileHeaders := form.File["image"]
	if len(fileHeaders) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	// Retrieve the first file (assuming single file upload)
	fileHeader := fileHeaders[0]

	// Open the uploaded file
	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer file.Close()

	// Create a unique filename for the uploaded file
	filename := strconv.Itoa(int(curUser.ID)) + "_" + fileHeader.Filename

	// Save the uploaded file to the server
	dst, err := os.Create("/uploads/" + filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create file"})
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Update the user's ImgURL with the filename or URL of the uploaded image
	// In this example, we assume the image is stored in a directory named "uploads"
	user.ImgURL = filename

	// Save the updated user to the database
	err = database.DB.Omit("password", "email", "role", "role_id", "contact_number", "valid_key").Save(&user).Error
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"updatedUser": user})

	tx.Commit()
}

// func UpdateTheUserProfile(context *gin.Context) {
// 	tx := database.DB.Begin()

// 	defer func() {
// 		if r := recover(); r != nil {
// 			tx.Rollback()
// 		}
// 	}()

// 	curUser := util.CurrentUser(context)

// 	var input struct {
// 		model.User
// 		Image *multipart.FileHeader `form:"image"`
// 	}

// 	if err := context.ShouldBindJSON(&input); err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	user, err := model.FindUserById(curUser.ID)
// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": "User not found."})
// 		return
// 	}

// 	if user.ID != curUser.ID {
// 		context.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied. You can only update your own details."})
// 		return
// 	}

// 	user.FirstName = input.FirstName
// 	user.LastName = input.LastName
// 	user.Gender = input.Gender
// 	user.Address = input.Address

// 	// Check if the user uploaded a new image
// 	if input.Image != nil {
// 		// Retrieve the uploaded file from the form
// 		file, err := input.Image.Open()
// 		if err != nil {
// 			context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to open uploaded file"})
// 			return
// 		}
// 		defer file.Close()

// 		// Create a unique filename for the uploaded file
// 		filename := strconv.Itoa(int(curUser.ID)) + "_" + input.Image.Filename

// 		// Save the uploaded file to the server
// 		if err := context.SaveUploadedFile(input.Image, "uploads/"+filename); err != nil {
// 			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
// 			return
// 		}

// 		// Update the user's ImgURL with the filename or URL of the uploaded image
// 		user.ImgURL = "/uploads/" + filename
// 	}

// 	if err := database.DB.Omit("password", "email", "role", "role_id", "contact_number", "valid_key").Save(&user).Error; err != nil {
// 		tx.Rollback()
// 		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	context.JSON(http.StatusOK, gin.H{"updatedUser": user})

// 	tx.Commit()
// }
