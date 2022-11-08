package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joshuaetim/quiz/domain/model"
	"github.com/joshuaetim/quiz/domain/repository"
	infrastructure "github.com/joshuaetim/quiz/infrastructure"
	"gorm.io/gorm"
)

type StaffHandler interface {
	GetUserStaff(*gin.Context)
	CreateUserStaff(*gin.Context)
	GetAllStaffByUser(*gin.Context)
	UpdateUserStaff(*gin.Context)
	DeleteUserStaff(*gin.Context)
}

type staffHandler struct {
	repo repository.StaffRepository
}

func NewStaffHandler(db *gorm.DB) StaffHandler {
	return &staffHandler{
		repo: infrastructure.NewStaffRepository(db),
	}
}

func (sh *staffHandler) CreateUserStaff(ctx *gin.Context) {
	var staff model.Staff
	if err := ctx.ShouldBindJSON(&staff); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "binding error: " + err.Error()})
		return
	}
	userID := ctx.GetFloat64("userID")
	staff.UserID = uint(userID)

	// TODO: check for empty fields
	staff, err := sh.repo.AddStaff(staff)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": staff})
}

func (sh *staffHandler) GetUserStaff(ctx *gin.Context) {
	userID := ctx.GetFloat64("userID")
	fmt.Println(userID)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "value error: " + err.Error()})
		return
	}
	staff, err := sh.repo.GetUserStaff(uint(id), uint(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "trouble fetching staff: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": staff.PublicStaff()})
}

func (sh *staffHandler) GetAllStaffByUser(ctx *gin.Context) {
	userID := ctx.GetFloat64("userID")
	staff, err := sh.repo.GetAllStaffByUser(uint(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "problem fetching user; " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": staff})
}

func (sh *staffHandler) UpdateUserStaff(ctx *gin.Context) {
	userID := ctx.GetFloat64("userID")
	staffID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "problem with input data: " + err.Error()})
		return
	}
	var staff model.Staff
	if err := ctx.ShouldBindJSON(&staff); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	staff.ID = uint(staffID)

	// first confirm the user has the staff
	_, err = sh.repo.GetUserStaff(uint(staffID), uint(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "trouble updating staff: " + err.Error()})
		return
	}

	staff, err = sh.repo.UpdateStaff(staff)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "problem updating staff; " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": staff, "msg": "staff updated"})
}

func (sh *staffHandler) DeleteUserStaff(ctx *gin.Context) {
	userID := ctx.GetFloat64("userID")
	staffID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "problem with input data: " + err.Error()})
		return
	}
	// first confirm the user has the staff
	_, err = sh.repo.GetUserStaff(uint(staffID), uint(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "trouble updating staff: " + err.Error()})
		return
	}

	var staff model.Staff
	staff.ID = uint(staffID)
	err = sh.repo.DeleteStaff(staff)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "problem deleting staff; " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "staff deleted"})
}
