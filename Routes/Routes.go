package Routes

import (
	"quinAI/Controllers"

	"github.com/labstack/echo/v4"
)

func Setup() *echo.Echo {

	e := echo.New()
	e.GET("/", Controllers.Get)
	e.POST("/Insert", Controllers.Insert)
	e.GET("/GetByID", Controllers.GetTaskByID)
	e.PATCH("/Update", Controllers.Update)

	// Start server
	e.Logger.Info(e.Start("localhost:8080"))
	return e
}
