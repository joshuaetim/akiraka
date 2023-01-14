package infrastructure

import (
	"github.com/joshuaetim/quiz/domain/model"
	"github.com/joshuaetim/quiz/domain/repository"
	"gorm.io/gorm"
)

type staffRepo struct {
	db *gorm.DB
}

func NewStaffRepository(db *gorm.DB) repository.StaffRepository {
	return &staffRepo{
		db: db,
	}
}

var _ repository.StaffRepository = &staffRepo{}

func (r *staffRepo) AddStaff(staff model.Staff) (model.Staff, error) {
	return staff, r.db.Create(&staff).Error
}

func (r *staffRepo) GetStaff(id uint) (model.Staff, error) {
	var staff model.Staff
	return staff, r.db.First(&staff, id).Error
}

func (r *staffRepo) GetAllStaff() ([]model.Staff, error) {
	var staff []model.Staff
	return staff, r.db.Find(&staff).Error
}

func (r *staffRepo) UpdateStaff(staff model.Staff) (model.Staff, error) {
	// exists?
	var sModel model.Staff = staff
	if err := r.db.First(&sModel).Error; err != nil {
		return staff, err
	}
	err := r.db.Model(&staff).Updates(&staff).Error
	staff, _ = r.GetStaff(staff.ID)
	return staff, err
}

func (r *staffRepo) DeleteStaff(staff model.Staff) error {
	// exists?
	if err := r.db.First(&staff).Error; err != nil {
		return err
	}
	return r.db.Delete(&staff).Error
}

func (r *staffRepo) GeneralSearch(query map[string]interface{}) (model.Staff, error) {
	staff := model.Staff{}
	return staff, r.db.Where(query).First(&staff).Error
}
