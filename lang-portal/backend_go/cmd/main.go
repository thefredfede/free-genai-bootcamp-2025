package main

import (
	"lang-portal-backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Register routes
	r.GET("/api/dashboard/last_study_session", handlers.GetLastStudySession)
	r.GET("/api/dashboard/study_progress", handlers.GetStudyProgress)
	r.GET("/api/dashboard/quick-stats", handlers.GetQuickStats)
	r.GET("/api/dashboard/mastery_progress", handlers.GetMasteryProgress)
	r.GET("/api/words", handlers.GetWords)
	r.GET("/api/words/:id", handlers.GetWordByID)
	r.GET("/api/groups", handlers.GetGroups)
	r.GET("/api/groups/:id", handlers.GetGroupByID)
	r.GET("/api/groups/:id/words", handlers.GetGroupWords)
	r.GET("/api/groups/:id/study_sessions", handlers.GetGroupStudySessions)
	r.GET("/api/study_activities/:id", handlers.GetStudyActivityByID)
	r.GET("/api/study_activities/:id/study_sessions", handlers.GetStudyActivityStudySessions)
	r.GET("/api/study_sessions", handlers.GetStudySessions)
	r.GET("/api/study_sessions/:id", handlers.GetStudySessionByID)
	r.GET("/api/study_sessions/:id/words", handlers.GetStudySessionWords)
	r.POST("/api/reset_history", handlers.ResetHistory)
	r.POST("/api/full_reset", handlers.FullReset)
	r.POST("/api/study_sessions", handlers.CreateStudySession)
	r.POST("/api/study_sessions/:id/review", handlers.ReviewStudySession)
	r.POST("/api/study_sessions/:id/words/:word_id/review", handlers.ReviewWord)
	r.POST("/api/study_activities", handlers.CreateStudyActivity)

	r.Run() // listen and serve on 0.0.0.0:8080
}
