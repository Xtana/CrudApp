package transport

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	httpTransport "crudapp/internal/transport/http"
)

func RegisterRoutes(e *echo.Echo, db *sql.DB) {
	userHandler := httpTransport.NewUserHandler(db)

	e.POST("/users", userHandler.CreateUser)
	e.GET("/users", userHandler.GetAllUsers)
	e.GET("/users/:id", userHandler.GetUser)
	e.PUT("/users/:id", userHandler.UpdateUser)
	e.DELETE("/users/:id", userHandler.DeleteUser)
}