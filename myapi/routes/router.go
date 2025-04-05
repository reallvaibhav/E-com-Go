package routes

import (
	"myapi/internal/product"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	product.RegisterRoutes(r)
}
