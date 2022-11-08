package factory

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/joshuaetim/quiz/domain/model"
	infrastructure "github.com/joshuaetim/quiz/infrastructure"
	"gorm.io/gorm"
)

func SeedStaff(dbInstance *gorm.DB, user uint) (model.Staff, error) {
	staff := model.Staff{
		Firstname: gofakeit.FirstName(),
		Lastname:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
		Phone:     gofakeit.Phone(),
		Address:   gofakeit.Address().Address,
		Status:    "active",
		UserID:    user,
	}
	sr := infrastructure.NewStaffRepository(dbInstance)
	s, err := sr.AddStaff(staff)
	return s, err
}
