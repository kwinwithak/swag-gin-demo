package handlers

import (
	"net/http"
	"swag-gin-demo/models"

	"github.com/gin-gonic/gin"
)

// todo slice to seed todo list data
var todoList = []models.Todo{
	{"1", "Learn Go"},
	{"2", "Build an API with Go"},
	{"3", "Document the API with swag"},
}

// @Summary get all items in the todo list
// @Description this is a description of sorts, so to speak, as it were
// @ID get-all-todos
// @Produce json
// @Success 200 {object} models.Todo
// @Router /todo [get]
func GetAllTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todoList)
}

// @Summary get a todo item by ID
// @Description this is a pretty verbose description showing
// @Description that multi line descriptions work and that swag
// @Description conveniently merges the text into a single string.
// @ID get-todo-by-id
// @Produce json
// @Param id path string true "todo ID"
// @Success 200 {object} models.Todo
// @Failure 404 {object} models.Message
// @Router /todo/{id} [get]
func GetTodoByID(c *gin.Context) {
	ID := c.Param("id")

	// loop through todoList and return item with matching ID
	for _, todo := range todoList {
		if todo.ID == ID {
			c.JSON(http.StatusOK, todo)
			return
		}
	}

	// return error message if todo is not found
	r := models.Message{"todo not found"}
	c.JSON(http.StatusNotFound, r)
}
