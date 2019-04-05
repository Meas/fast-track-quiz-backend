package repositories

import (
	"github.com/meas/fast-track-quiz-backend/helpers"
	"github.com/meas/fast-track-quiz-backend/models"
)

// GetQuizResult calculates the result of the user session, TODO: check if session is repeated and throw error
func GetQuizResult(userSession models.SessionSubmission, correctAnswers []models.SubmittedAnswer, allSubmissions *[]models.Score) models.Result {
	var result models.Result
	result.Points = getScore(userSession.Answers, correctAnswers)
	result.BetterThan, result.BetterOrEqualThan = getComparisons(result.Points, allSubmissions)

	*allSubmissions = append(*allSubmissions, models.Score{
		ID:     userSession.ID,
		Points: result.Points,
	})

	return result
}

func getScore(userAnswers []models.SubmittedAnswer, correctAnswers []models.SubmittedAnswer) int {
	var score int
	for _, userAnswer := range userAnswers {
		for _, correctAnswer := range correctAnswers {
			if userAnswer.QuestionID != correctAnswer.QuestionID {
				continue
			}
			if userAnswer.AnswerID == correctAnswer.AnswerID {
				score++
			}
			break
		}
	}
	return score
}

func getComparisons(userPoints int, allSubmissions *[]models.Score) (float32, float32) {
	var betterThan float32
	var betterOrEqualThen float32
	var numberOfSubmissions = float32(len(*allSubmissions))
	if numberOfSubmissions == 0 {
		return 1, 1
	}
	for _, submission := range *allSubmissions {
		if userPoints >= submission.Points {
			betterOrEqualThen++
			if userPoints > submission.Points {
				betterThan++
			}
		}
	}
	return helpers.RoundUp(betterThan / numberOfSubmissions), helpers.RoundUp(betterOrEqualThen / numberOfSubmissions)
}
