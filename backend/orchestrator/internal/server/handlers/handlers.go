package handlers

import (
	"net/http"
	mdb "orchstrator/pkg/db/mongodb"

	"github.com/labstack/echo/v4"
)

type Router struct {
	MongoDB *mdb.Mongo
}

func NewRouter(MongoDB *mdb.Mongo) *Router {
	return &Router{MongoDB: MongoDB}
}

// Just Ping function
func (R *Router) GetPing(c echo.Context) error {
	return c.JSON(http.StatusTeapot, "Sorry, I'm a teapot")
}

func (R *Router) PostCalculate(c echo.Context) error {
	// TODO: Принять выражение, разбить на задачи, сохранить
	return c.JSON(http.StatusTeapot, "Sorry, I'm a teapot")
}

func (R *Router) GetExpressions(c echo.Context) error {
	// TODO: Вернуть все выражения
	return c.JSON(http.StatusTeapot, "Sorry, I'm a teapot")
}

func (R *Router) GetExpressionById(c echo.Context) error {
	// TODO: Вернуть выражение по ID
	return c.JSON(http.StatusTeapot, "Sorry, I'm a teapot")
}

func (R *Router) GetTask(c echo.Context) error {
	// TODO: Выдать задачу
	return c.JSON(http.StatusTeapot, "Sorry, I'm a teapot")
}

func (R *Router) PostTask(c echo.Context) error {
	// TODO: Принять задачу
	return c.JSON(http.StatusTeapot, "Sorry, I'm a teapot")
}
