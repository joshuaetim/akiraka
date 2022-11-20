package infrastructure

import (
	"github.com/joshuaetim/quiz/domain/model"
	"github.com/joshuaetim/quiz/domain/repository"
	"gorm.io/gorm"
)

type scoreRepo struct {
	db *gorm.DB
}

func NewScoreRepository(db *gorm.DB) repository.ScoreRepository {
	return &scoreRepo{
		db: db,
	}
}

var _ repository.ScoreRepository = &scoreRepo{}

func (r *scoreRepo) AddScore(score model.Score) (model.Score, error) {
	return score, r.db.Create(&score).Error
}

func (r *scoreRepo) GetAllScores() ([]model.Score, error) {
	var scores []model.Score
	return scores, r.db.Find(&scores).Error
}

func (r *scoreRepo) GetScore(id uint) (model.Score, error) {
	var score model.Score
	return score, r.db.First(&score, "id = ?", id).Error
}

func (r *scoreRepo) GetScoresBySession(id string) ([]model.Score, error) {
	var scores []model.Score
	return scores, r.db.Find(&scores, "session_id = ?", id).Error
}
