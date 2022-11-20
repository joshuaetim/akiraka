package repository

import "github.com/joshuaetim/quiz/domain/model"

type QuizRepository interface {
	AddQuiz(model.Quiz) (model.Quiz, error)
	GetAllQuiz() ([]model.Quiz, error)
	GetQuizBySession(string) ([]model.Quiz, error)
	GetQuiz(uint) (model.Quiz, error)
}
