package model

import "gorm.io/gorm"

type Quiz struct {
	gorm.Model
	Question  string `json:"question"`
	Options   string `json:"options"`
	Answer    string `json:"answer"`
	StaffID   uint   `json:"staff,omitempty"`
	SessionID string `json:"sessionID"`
}

func (q *Quiz) PublicQuiz() *Quiz {
	quiz := &Quiz{
		Question: q.Question,
		Options:  q.Options,
		StaffID:  q.StaffID,
	}
	quiz.ID = q.ID
	quiz.CreatedAt = q.CreatedAt
	return quiz
}

func (Quiz) TableName() string {
	return "quiz"
}
