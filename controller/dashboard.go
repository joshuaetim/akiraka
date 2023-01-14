package controller

// controller handles methods that do not necessarily have repositories (DB)

import (
	infrastructure "github.com/joshuaetim/quiz/infrastructure"
	"gorm.io/gorm"
)

type DashboardController interface {
	GetUsersCount() int
}

func NewDashboardController(db *gorm.DB) DashboardController {
	return dashboardController{
		db: db,
	}
}

type dashboardController struct {
	db *gorm.DB
}

func (dash dashboardController) GetUsersCount() int {
	userDB := infrastructure.NewUserRepository(dash.db)
	return userDB.CountUsers()
}
