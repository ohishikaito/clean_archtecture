package infrastructure

import (
	"github.com/gin-gonic/gin"

	"app/app/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	router := gin.Default()
	userController := controllers.NewUserController(NewSqlHandler())

	router.POST("/users", func(c *gin.Context) {
		userController.Create(c)
	})
	router.GET("/users", func(c *gin.Context) {
		userController.Index(c)
	})
	router.GET("/users/:id", func(c *gin.Context) {
		userController.Show(c)
	})
	router.DELETE("/users/:id", func(c *gin.Context) {
		userController.Delete(c)
	})

	Router = router
}
