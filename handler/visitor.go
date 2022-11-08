package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joshuaetim/quiz/domain/model"
	"github.com/joshuaetim/quiz/domain/repository"
	infrastructure "github.com/joshuaetim/quiz/infrastructure"
	"gorm.io/gorm"
)

type VisitorHandler interface {
	GetUserVisitor(*gin.Context)
	CreateUserVisitor(*gin.Context)
	GetAllUserVisitors(*gin.Context)
	GetAllStaffVisitors(*gin.Context)
	UpdateUserVisitor(*gin.Context)
	DeleteUserVisitor(*gin.Context)
}

type visitorHandler struct {
	repo repository.VisitorRepository
}

func NewVisitorHandler(db *gorm.DB) VisitorHandler {
	return &visitorHandler{
		repo: infrastructure.NewVisitorRepository(db),
	}
}

func (vh *visitorHandler) GetUserVisitor(ctx *gin.Context) {
	visitorID, err := strconv.Atoi(ctx.Param("id"))
	if uint(visitorID) <= 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request; issue with id parameter: %v" + err.Error()})
		return
	}
	userID := ctx.GetFloat64("userID")
	visitor, err := vh.repo.GetUserVisitor(uint(visitorID), uint(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "trouble fetching visitor: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": visitor.PublicVisitor()})

}

func (vh *visitorHandler) CreateUserVisitor(ctx *gin.Context) {
	var visitor model.Visitor
	if err := ctx.ShouldBindJSON(&visitor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "binding error: " + err.Error()})
		return
	}
	userID := ctx.GetFloat64("userID")
	visitor.UserID = uint(userID)

	visitor, err := vh.repo.AddVisitor(visitor)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error creating visitor: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": visitor})
}

func (vh *visitorHandler) GetAllUserVisitors(ctx *gin.Context) {
	userID := ctx.GetFloat64("userID")
	visitors, err := vh.repo.GetAllUserVisitor(uint(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "trouble fetching visitor: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": visitors})
}

func (vh *visitorHandler) GetAllStaffVisitors(ctx *gin.Context) {
	staffID, err := strconv.Atoi(ctx.Param("staffID"))
	if staffID <= 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request; issue with id parameter: %v" + err.Error()})
		return
	}
	userID := ctx.GetFloat64("userID")
	visitors, err := vh.repo.GetAllUserAndStaffVisitor(uint(userID), uint(staffID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "trouble fetching visitors: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": visitors})
}

func (vh *visitorHandler) UpdateUserVisitor(ctx *gin.Context) {
	// get input param
	visitorID, err := strconv.Atoi(ctx.Param("id"))
	if uint(visitorID) <= 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request; issue with id parameter: %v" + err.Error()})
		return
	}
	// bind visitor input
	var visitor model.Visitor
	if err := ctx.ShouldBindJSON(&visitor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "binding error: " + err.Error()})
		return
	}
	visitor.ID = uint(visitorID)

	userID := ctx.GetFloat64("userID")

	// check if visitor belongs to user
	_, err = vh.repo.GetUserVisitor(uint(visitorID), uint(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "record not found"})
		return
	}

	visitor, err = vh.repo.UpdateVisitor(visitor)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "problem updating visitor: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": visitor, "msg": "visitor updated"})
}

func (vh *visitorHandler) DeleteUserVisitor(ctx *gin.Context) {
	// get input param
	visitorID, err := strconv.Atoi(ctx.Param("id"))
	if uint(visitorID) <= 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request; issue with id parameter: %v" + err.Error()})
		return
	}
	userID := ctx.GetFloat64("userID")

	// check if visitor belongs to user
	_, err = vh.repo.GetUserVisitor(uint(visitorID), uint(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "record not found"})
		return
	}

	var visitor model.Visitor
	visitor.ID = uint(visitorID)
	err = vh.repo.DeleteVisitor(visitor)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cannot delete visitor: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "visitor deleted"})
}
