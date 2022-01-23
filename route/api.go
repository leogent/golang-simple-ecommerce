package route

import (
	"net/http"
	"simple-ecommerce/handler"

	"github.com/gin-gonic/gin"
)

func RunApi(address string) error {
	userHandler := handler.NewUserHandler()

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome to Our Store")
	})

	apiRoutes := r.Group("/api")
	userRoutes := apiRoutes.Group("/user")
	{
		userRoutes.POST("/register", userHandler.AddUser)
		userRoutes.POST("/signin", userHandler.SignInUser)
	}

	return r.Run(address)
}
