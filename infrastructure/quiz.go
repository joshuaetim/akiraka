package infrastructure

import (
	"github.com/joshuaetim/quiz/domain/model"
	"github.com/joshuaetim/quiz/domain/repository"
	"gorm.io/gorm"
)

type quizRepo struct {
	db *gorm.DB
}

func NewQuizRepository(db *gorm.DB) repository.QuizRepository {
	return &quizRepo{
		db: db,
	}
}

var _ repository.StaffRepository = &staffRepo{}

func (r *quizRepo) AddQuiz(quiz model.Quiz) (model.Quiz, error) {
	return quiz, r.db.Create(&quiz).Error
}

func (r *quizRepo) GetAllQuiz() ([]model.Quiz, error) {
	var quiz []model.Quiz
	return quiz, r.db.Find(&quiz).Error
}

func (r *quizRepo) GetQuiz(id uint) (model.Quiz, error) {
	var quiz model.Quiz
	return quiz, r.db.First(&quiz, "id = ?", id).Error
}

func (r *quizRepo) GetQuizBySession(id string) ([]model.Quiz, error) {
	var quiz []model.Quiz
	return quiz, r.db.Find(&quiz, "session_id = ?", id).Error
}
