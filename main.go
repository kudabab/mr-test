package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kudabab/mr-test/controller"
	"github.com/kudabab/mr-test/db"
	"github.com/kudabab/mr-test/service"
)

var (
	userService    service.UserService       = service.New()
	userController controller.UserController = controller.New(userService)
)

func main() {

	db.InitDB()

	router := gin.Default()

	router.GET("/users", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, userController.FindAll())
	})
	router.POST("/users", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, userController.Save(ctx))
	})
	router.Run(":8080")
}
