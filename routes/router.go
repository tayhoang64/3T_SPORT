package routes

import (
	"3T_Shop/config"
	"3T_Shop/handler"
	"github.com/labstack/echo/v4"
	"os"
)

var (
	RoleHandler            handler.RoleHandler
	ProductCategoryHandler handler.ProductCategoryHandler
)

func InitRoutes() {
	config.LoadENV()
	router := echo.New()

	api := router.Group("/api")
	{
		role := api.Group("/role")
		{
			role.GET("/", RoleHandler.GetAll)
			role.POST("/", RoleHandler.Create)
		}

		productCategory := api.Group("/category")
		{
			productCategory.GET("/", ProductCategoryHandler.GetAll)
			productCategory.POST("/", ProductCategoryHandler.Create)
		}
	}

	router.Logger.Fatal(router.Start(":" + os.Getenv("PORT")))
}
