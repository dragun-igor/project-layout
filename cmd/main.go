package main

import (
	"github.com/dragun-igor/project-layout/internal/domains/users"
	"github.com/labstack/echo/v4"
)

func main() {
	repo := &users.Repo{}
	service := users.NewService(repo)
	handler := users.NewHTTPHandler(service)

	echo.New().POST("", handler.Create)
}
