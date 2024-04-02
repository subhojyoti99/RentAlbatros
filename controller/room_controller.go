package controller

import (
	"ToLet/database"
	"ToLet/model"
	"ToLet/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateRoom(context *gin.Context) {

	tx := database.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	user := util.CurrentUser(context)

	if user.Role != "admin" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied. Only Admins allowed."})
		return
	}

	var input model.Room

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	newRoom := model.Room{
		OwnerID:        user.ID,
		RoomType:       input.RoomType,
		Gender:         input.Gender,
		Location:       input.Location,
		Cost:           input.Cost,
		SecurityCost:   input.SecurityCost,
		Capacity:       input.Capacity,
		Year:           input.Year,
		Feature:        input.Feature,
		Description:    input.Description,
		Status:         input.Status,
		AvailableRooms: input.AvailableRooms,
	}

	addedRoom, err := newRoom.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// database.DB.Create(&newRoom)

	// Fetch owner information and set it in the response
	owner, err := model.FindUserById(user.ID) // You need to implement GetUserByID function in your model package
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch owner information"})
		return
	}

	addedRoom.Owner = owner

	context.JSON(http.StatusOK, gin.H{"new_room": addedRoom})

	tx.Commit()
}

func GetAllRooms(context *gin.Context) {
	tx := database.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var rooms []model.Room

	err := database.DB.Preload("Owner").Find(&rooms).Error
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"Rooms": rooms})

	tx.Commit()
}

func GetRoom(context *gin.Context) {
	tx := database.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var room model.Room

	id, _ := strconv.Atoi(context.Param("id"))

	err := database.DB.Where("id=?", id).Preload("Owner").Preload("Owner.OwnedRooms").Preload("Owner.Rentals").Find(&room).Error
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"Room": room})

	tx.Commit()
}
