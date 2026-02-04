package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ovitorvalente/gin-api-rest/internal/models"
)

func ListStudents(c *gin.Context) {
	students := []models.Student{
		{
			ID:        uuid.New(),
			Name:      "Vitor Valente",
			Age:       23,
			Major:     "Computer Science",
			CreatedAt: time.Now().AddDate(0, 1, 0),
			UpdatedAt: time.Now().AddDate(0, 1, 0),
		},
		{
			ID:        uuid.New(),
			Name:      "Alice",
			Age:       20,
			Major:     "Computer Science",
			CreatedAt: time.Now().AddDate(0, 1, 0),
			UpdatedAt: time.Now().AddDate(0, 1, 0),
		},
		{
			ID:        uuid.New(),
			Name:      "Bob",
			Age:       22,
			Major:     "Mathematics",
			CreatedAt: time.Now().AddDate(0, 2, 0),
			UpdatedAt: time.Now().AddDate(0, 2, 0),
		},
	}

	var filter models.StudentFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	var filteredStudents []models.Student

	for _, student := range students {
		if filter.ID != "" && student.ID.String() != filter.ID {
			continue
		}

		if filter.Name != "" &&
			!strings.Contains(
				strings.ToLower(student.Name),
				strings.ToLower(filter.Name),
			) {
			continue
		}

		filteredStudents = append(filteredStudents, student)
	}

	c.JSON(http.StatusOK, gin.H{
		"total":    len(filteredStudents),
		"students": filteredStudents,
	})
}
