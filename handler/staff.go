package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshuaetim/quiz/domain/model"
	"github.com/joshuaetim/quiz/domain/repository"
	infrastructure "github.com/joshuaetim/quiz/infrastructure"
	"gorm.io/gorm"
)

type StaffHandler interface {
	RegisterStaff(*gin.Context)
	UpdateUserStaff(*gin.Context)
}

type staffHandler struct {
	repo repository.StaffRepository
}

func NewStaffHandler(db *gorm.DB) StaffHandler {
	return &staffHandler{
		repo: infrastructure.NewStaffRepository(db),
	}
}

func (sh *staffHandler) RegisterStaff(ctx *gin.Context) {
	var staff model.Staff
	if err := ctx.ShouldBindJSON(&staff); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "binding error: " + err.Error()})
		return
	}

	// TODO: check for empty fields
	staff, err := sh.repo.AddStaff(staff)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": staff})
}

func (sh *staffHandler) UpdateUserStaff(ctx *gin.Context) {
	staffID := ctx.GetFloat64("userID")
	var staff model.Staff
	if err := ctx.ShouldBindJSON(&staff); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	staff.ID = uint(staffID)

	staff, err := sh.repo.UpdateStaff(staff)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "problem updating staff; " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": staff, "msg": "staff updated"})
}
