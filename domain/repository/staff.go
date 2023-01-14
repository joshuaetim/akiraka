package repository

import "github.com/joshuaetim/quiz/domain/model"

type StaffRepository interface {
	AddStaff(model.Staff) (model.Staff, error)
	GetStaff(uint) (model.Staff, error)
	GetAllStaff() ([]model.Staff, error)
	UpdateStaff(model.Staff) (model.Staff, error)
	DeleteStaff(model.Staff) error
	GeneralSearch(map[string]interface{}) (model.Staff, error)
}
