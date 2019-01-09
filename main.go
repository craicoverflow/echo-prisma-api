package main

import (
	"context"
	"net/http"

	prisma "github.com/craicoverflow/echo-prisma-api/generated/prisma-client"
	"github.com/labstack/echo"
)

type responseError struct {
	Message string `json:"message"`
}

func main() {
	e := echo.New()

	e.POST("/users", createUser)
	e.GET("users", getUsers)
	e.GET("/users/:id", getUserByID)
	e.DELETE("/users/:id", deleteUserWithID)

	e.Logger.Fatal(e.Start(":9000"))
}

func createUser(c echo.Context) error {
	client := prisma.New(nil)
	ctx := context.TODO()

	name := "Enda"
	newUser, err := client.CreateUser(prisma.UserCreateInput{
		Name: name,
	}).Exec(ctx)

	if err != nil {
		return c.JSON(400, responseError{err.Error()})
	}

	return c.JSON(http.StatusCreated, newUser)
}

func getUserByID(c echo.Context) error {
	client := prisma.New(nil)
	ctx := context.TODO()

	id := c.Param("id")

	user, err := client.User(prisma.UserWhereUniqueInput{
		ID: &id,
	}).Exec(ctx)

	if err != nil {
		return c.JSON(404, responseError{"User not found"})
	}

	return c.JSON(http.StatusOK, user)
}

func deleteUserWithID(c echo.Context) error {
	client := prisma.New(nil)
	ctx := context.TODO()

	id := c.Param("id")

	user, err := client.DeleteUser(prisma.UserWhereUniqueInput{
		ID: &id,
	},
	).Exec(ctx)

	if err != nil {
		return c.JSON(404, responseError{"User not found"})
	}

	return c.JSON(http.StatusOK, user)
}

func getUsers(c echo.Context) error {
	client := prisma.New(nil)
	ctx := context.TODO()

	users, err := client.Users(nil).Exec(ctx)

	if err != nil {
		return c.JSON(404, responseError{"No users found"})
	}

	return c.JSON(http.StatusOK, users)
}
