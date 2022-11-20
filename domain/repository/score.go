package repository

import "github.com/joshuaetim/quiz/domain/model"

type ScoreRepository interface {
	AddScore(model.Score) (model.Score, error)
	GetAllScores() ([]model.Score, error)
	GetScoresBySession(string) ([]model.Score, error)
	GetScore(uint) (model.Score, error)
}
