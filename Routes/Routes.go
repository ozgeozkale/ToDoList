package Routes

import (
	"ToDoProject/Controllers"

	"github.com/labstack/echo/v4"
)

func Setup() *echo.Echo {

	e := echo.New()
	e.GET("ToDoApp/v1/", Controllers.Get)
	e.POST("ToDoApp/v1/Insert", Controllers.Insert)
	e.GET("ToDoApp/v1/GetByID", Controllers.GetTaskByID)
	e.PATCH("ToDoApp/v1/Update", Controllers.Update)

	// Start server
	e.Logger.Info(e.Start("localhost:8080"))
	return e
}
