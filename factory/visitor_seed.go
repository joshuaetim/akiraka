package factory

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/joshuaetim/quiz/domain/model"
	infrastructure "github.com/joshuaetim/quiz/infrastructure"
	"gorm.io/gorm"
)

func SeedVisitor(dbInstance *gorm.DB, user uint, staff uint) (model.Visitor, error) {
	visitor := model.Visitor{
		Firstname:  gofakeit.FirstName(),
		Lastname:   gofakeit.LastName(),
		Email:      gofakeit.Email(),
		Phone:      gofakeit.Phone(),
		Occupation: gofakeit.JobTitle(),
		Note:       gofakeit.Sentence(30),
		StaffID:    staff,
		UserID:     user,
	}
	vr := infrastructure.NewVisitorRepository(dbInstance)
	visitor, err := vr.AddVisitor(visitor)

	return visitor, err
}
