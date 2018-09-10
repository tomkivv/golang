package handler

import (
	"github.com/labstack/echo"
	"github.com/vtomkiv/golang.api/api"
	"net/http"
	"strconv"
)

type TaskController struct {
	TaskService api.TaskService
}

func (tc *TaskController) FindTask(c echo.Context) error {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	user, _:= tc.TaskService.GetTask(id)

	if user != nil {

		 return c.JSON(http.StatusOK, user)
	 }else {

	 	return echo.ErrNotFound
	 }
}

func (tc *TaskController) CreateTask(c echo.Context) error {

	var task api.Task

	c.Bind(&task)

	id, _:= tc.TaskService.CreateTask(&task)

	createdTask, _ := tc.TaskService.GetTask(id)

	return c.JSON(http.StatusCreated, createdTask)
}

func (tc *TaskController) UpdateTask(c echo.Context) error {
	var task api.Task

	c.Bind(&task)

	updatedTask, _ := tc.TaskService.UpdateTask(&task)

	return c.JSON(http.StatusOK, updatedTask)
}

func (tc *TaskController) DeleteTask(c echo.Context) error {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	tc.TaskService.DeleteTask(id)

	return c.NoContent(http.StatusNoContent)
}