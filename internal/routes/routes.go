package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ovitorvalente/gin-api-rest/internal/controllers"
)

func HandlerRequests() {
	r := gin.Default()

	r.GET("/students", controllers.ListStudents)
	r.GET("/courses", controllers.ListCourses)
	port := ":5000"
	r.Run(port)
}
