package Controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"quinAI/Models"
	"quinAI/Service"
)

var err error

func Get(c echo.Context) error {
	var parameters Models.InputModel
	c.Bind(&parameters)
	err, returnmodel := Service.Get()
	if err != nil {
		return c.String(http.StatusNotFound, "")
	} else {
		return c.JSON(http.StatusOK, returnmodel)
	}
}

func Insert(c echo.Context) error {
	c.Request().Header.Set("Content-Type", echo.MIMEApplicationJSONCharsetUTF8)

	var parameters Models.InputModel
	c.Bind(&parameters)
	err, returnmodel := Service.Insert(parameters)
	if err != nil {
		return c.String(http.StatusNotFound, "")
	} else {
		return c.JSON(http.StatusOK, returnmodel)
	}
}
func Update(c echo.Context) error {
	c.Request().Header.Set("Content-Type", echo.MIMEApplicationJSONCharsetUTF8)

	var parameters Models.InputModel
	c.Bind(&parameters)
	err, returnmodel := Service.Update(parameters)
	if err != nil {
		return c.String(http.StatusNotFound, "")
	} else {
		return c.JSON(http.StatusOK, returnmodel)
	}
}

func GetTaskByID(c echo.Context) error {
	c.Request().Header.Set("Content-Type", echo.MIMEApplicationJSONCharsetUTF8)
	u := new(Models.IdModel)
	if err := c.Bind(u); err != nil {
		return err
	}

	id := Models.IdModel{Id: u.Id}
	err, returnmodel := Service.GetTaskByID(id)
	if err != nil {
		return c.String(http.StatusNotFound, "")
	} else {
		return c.JSON(http.StatusOK, returnmodel)
	}
}
