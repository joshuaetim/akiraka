package handler

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joshuaetim/quiz/common"
	"github.com/joshuaetim/quiz/domain/model"
	"github.com/joshuaetim/quiz/domain/repository"
	infrastructure "github.com/joshuaetim/quiz/infrastructure"
	"gorm.io/gorm"
)

const (
	ErrorAlreadySubmitted = "ERR_REPEAT"
)

type Answer struct {
	UserID    uint   `json:"userID"`
	SessionID string `json:"sessionID"`
	Answers   string `json:"answers"`
}

type QuizHandler interface {
	CreateQuiz(*gin.Context)
	GetAllQuiz(*gin.Context)
	GetQuizBySession(*gin.Context)
	GradeQuiz(*gin.Context)
	UploadQuestions(*gin.Context)
}

type quizHandler struct {
	repo      repository.QuizRepository
	scoreRepo repository.ScoreRepository
	userRepo  repository.UserRepository
}

func NewQuizHandler(db *gorm.DB) QuizHandler {
	return &quizHandler{
		repo:      infrastructure.NewQuizRepository(db),
		scoreRepo: infrastructure.NewScoreRepository(db),
		userRepo:  infrastructure.NewUserRepository(db),
	}
}

func (sh *quizHandler) CreateQuiz(ctx *gin.Context) {
	var quiz model.Quiz
	if err := ctx.ShouldBindJSON(&quiz); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "binding error: " + err.Error()})
		return
	}
	userID := ctx.GetFloat64("userID")
	quiz.StaffID = uint(userID)

	// TODO: check for empty fields
	quiz, err := sh.repo.AddQuiz(quiz)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": quiz})
}

func (sh *quizHandler) GetAllQuiz(ctx *gin.Context) {
	// userID := ctx.GetFloat64("userID")
	quiz, err := sh.repo.GetAllQuiz()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "problem fetching quiz; " + err.Error()})
		return
	}
	var quizPublic []model.Quiz
	for _, q := range quiz {
		quizPublic = append(quizPublic, *q.PublicQuiz())
	}
	ctx.JSON(http.StatusOK, gin.H{"data": quizPublic})
}

func (sh *quizHandler) GetQuizBySession(ctx *gin.Context) {
	sessionID := ctx.Query("sessionID")
	if sessionID == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "sessionID cannot be empty"})
		return
	}
	// check if user has submitted quiz
	uID := uint(ctx.GetFloat64("userID"))
	user, err := sh.userRepo.GetUser(uID)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	scoreFromDB, err := sh.scoreRepo.GeneralSearch(map[string]interface{}{
		"matric":     user.Matric,
		"session_id": sessionID,
	})
	if err == nil || scoreFromDB.Matric != "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "you have already submitted", "key": ErrorAlreadySubmitted})
		return
	}
	quiz, err := sh.repo.GetQuizBySession(sessionID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "problem fetching quiz; " + err.Error()})
		return
	}
	var quizPublic []model.Quiz
	for _, q := range quiz {
		quizPublic = append(quizPublic, *q.PublicQuiz())
	}
	ctx.JSON(http.StatusOK, gin.H{"data": quizPublic})
}

func (qh *quizHandler) GradeQuiz(ctx *gin.Context) {
	var queryAnswer Answer
	if err := ctx.ShouldBindJSON(&queryAnswer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "binding error: " + err.Error()})
		return
	}
	queryAnswer.Answers = strings.Trim(queryAnswer.Answers, " ")
	if queryAnswer.Answers == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "answers should be filled"})
		return
	}
	var answerMap map[string]string
	if err := json.Unmarshal([]byte(queryAnswer.Answers), &answerMap); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "marshal error: " + err.Error()})
		return
	}
	score := 0

	for question, selected := range answerMap {
		questionID, err := strconv.ParseUint(question, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		quiz, err := qh.repo.GetQuiz(uint(questionID))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(selected, quiz.Answer)
		if quiz.Answer == selected {
			score++
		}
	}
	percent := (float64(score) / float64(len(answerMap))) * 100
	percent = math.Ceil((percent * 100) / 100)

	queryAnswer.UserID = uint(ctx.GetFloat64("userID"))
	err := AddScore(queryAnswer, ctx, qh.scoreRepo, qh.userRepo, percent)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": percent})
}

func (qh *quizHandler) UploadQuestions(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error: file not uploaded: " + err.Error()})
		return
	}
	f, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error opening file: " + err.Error()})
		return
	}
	reader := csv.NewReader(f)
	type quiz struct {
		Question string
		Options  string
		Answer   string
	}
	data := []quiz{}
	// var resultData [][]string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error reading CSV file: " + err.Error()})
			return
		}
		quizData := quiz{}
		for i, value := range record {
			switch i {
			case 0:
				quizData.Question = value
			case 1:
				quizData.Options = value
			case 2:
				_, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": "wrong type: option must be a number" + err.Error()})
					return
				}
				quizData.Answer = value
			}
		}
		data = append(data, quizData)
	}
	sessionID := common.RandStringBytes(5)
	fmt.Printf("%+v\n", data)
	for _, quizInfo := range data {
		quizModel := model.Quiz{
			Question:  quizInfo.Question,
			Options:   quizInfo.Options,
			Answer:    quizInfo.Answer,
			StaffID:   1,
			SessionID: sessionID,
		}
		_, err := qh.repo.AddQuiz(quizModel)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error creating quiz" + err.Error()})
			return
		}
	}
}

func AddScore(queryAnswer Answer, ctx *gin.Context, scoreRepo repository.ScoreRepository, userRepo repository.UserRepository, percent float64) error {
	uID := queryAnswer.UserID

	user, err := userRepo.GetUser(uint(uID))
	if err != nil {
		return err
	}
	sessionID := queryAnswer.SessionID
	if strings.Trim(sessionID, " ") == "" {
		return fmt.Errorf("session ID must be provided")
	}

	_, err = scoreRepo.AddScore(model.Score{
		Matric:    user.Matric,
		Score:     fmt.Sprintf("%.2f", percent),
		SessionID: sessionID,
	})
	if err != nil {
		return err
	}
	return nil
}
