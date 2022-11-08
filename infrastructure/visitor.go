package infrastructure

import (
	"github.com/joshuaetim/quiz/domain/model"
	"github.com/joshuaetim/quiz/domain/repository"
	"gorm.io/gorm"
)

type visitorRepo struct {
	db *gorm.DB
}

func NewVisitorRepository(db *gorm.DB) repository.VisitorRepository {
	return &visitorRepo{
		db: db,
	}
}

var _ repository.VisitorRepository = &visitorRepo{}

func (r *visitorRepo) AddVisitor(visitor model.Visitor) (model.Visitor, error) {
	return visitor, r.db.Create(&visitor).Error
}

func (r *visitorRepo) GetVisitor(id uint) (model.Visitor, error) {
	var visitor model.Visitor
	return visitor, r.db.First(&visitor, id).Error
}

func (r *visitorRepo) GetUserVisitor(id, userid uint) (model.Visitor, error) {
	var visitor model.Visitor
	return visitor, r.db.First(&visitor, "id = ? AND user_id = ?", id, userid).Error
}

func (r *visitorRepo) GetAllVisitor() ([]model.Visitor, error) {
	var visitors []model.Visitor
	return visitors, r.db.Find(&visitors).Error
}

func (r *visitorRepo) GetAllUserVisitor(userid uint) ([]model.Visitor, error) {
	var visitors []model.Visitor
	return visitors, r.db.Find(&visitors, "user_id = ?", userid).Error
}

func (r *visitorRepo) GetAllUserAndStaffVisitor(user_id uint, staff_id uint) ([]model.Visitor, error) {
	var visitors []model.Visitor
	return visitors, r.db.Find(&visitors, "user_id = ? AND staff_id = ?", user_id, staff_id).Error
}

func (r *visitorRepo) UpdateVisitor(visitor model.Visitor) (model.Visitor, error) {
	// exists?
	var visitorModel = visitor
	if err := r.db.First(&visitorModel).Error; err != nil {
		return visitor, err
	}
	err := r.db.Model(&visitor).Updates(&visitor).Error
	visitor, _ = r.GetVisitor(visitor.ID)
	return visitor, err
}

func (r *visitorRepo) DeleteVisitor(visitor model.Visitor) error {
	// exists?
	if err := r.db.First(&visitor).Error; err != nil {
		return err
	}
	return r.db.Delete(&visitor).Error
}
