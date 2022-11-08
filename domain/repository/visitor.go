package repository

import "github.com/joshuaetim/quiz/domain/model"

type VisitorRepository interface {
	AddVisitor(model.Visitor) (model.Visitor, error)
	GetVisitor(uint) (model.Visitor, error)
	GetUserVisitor(uint, uint) (model.Visitor, error)
	GetAllVisitor() ([]model.Visitor, error)
	GetAllUserVisitor(uint) ([]model.Visitor, error)
	GetAllUserAndStaffVisitor(uint, uint) ([]model.Visitor, error)
	UpdateVisitor(model.Visitor) (model.Visitor, error)
	DeleteVisitor(model.Visitor) error
}
