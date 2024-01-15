package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// template
type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

// data
var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Code", Completed: false},
}

// get
func getTodos(context *gin.Context) {

	context.IndentedJSON(http.StatusOK, todos)

}

// post
func addTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

// handler
func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
	}

	context.IndentedJSON(http.StatusOK, todo)
}

// patch
func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
	}

	//flip boolean values - toggle todo
	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}

// list ids
func getTodoById(id string) (*todo, error) {

	//iterate through todos array
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

// main function
func main() {

	//server
	router := gin.Default()

	router.GET("/todo", getTodos)
	router.GET("/todo/:id", getTodo)
	router.PATCH("/todo/:id", toggleTodoStatus)
	router.POST("/todo", addTodo)
	router.Run("localhost:7777")
}
