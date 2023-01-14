package model

import "gorm.io/gorm"

type Score struct {
	gorm.Model
	Matric    string `json:"matric"`
	Score     string `json:"score"`
	SessionID string `json:"sessionID"`
}

func (score *Score) PublicScore() *Score {
	scoreModel := &Score{
		Matric:    score.Matric,
		Score:     score.Score,
		SessionID: score.SessionID,
	}
	scoreModel.ID = score.ID
	scoreModel.CreatedAt = score.CreatedAt
	return scoreModel
}

func (Score) TableName() string {
	return "score"
}
