package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ovitorvalente/gin-api-rest/internal/models"
)

func ListCourses(c *gin.Context) {
	courses := []models.Course{
		{
			ID:          uuid.New(),
			Name:        "Introduction to Go",
			Description: "Learn the basics of Go programming language.",
			Level:       "Beginner",
			Duration:    40,
			CreatedAt:   time.Now().AddDate(0, -1, 0),
			UpdatedAt:   time.Now().AddDate(0, -1, 0),
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"courses": courses,
	})
}
