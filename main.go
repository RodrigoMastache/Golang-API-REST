package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	//mapping
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record Video", Completed: false},
}

// Parameter "context" of type "gin.Context"
// "context" will contain a bunch of info about the incoming HTTP requests
func getTodos(context *gin.Context) {
	//supply the status of our incoming request and the json
	//will convert the data structure "todos" into json
	context.IndentedJSON(http.StatusOK, todos)
}

// Add new record
func addTodo(context *gin.Context) {
	var newTodo todo

	//We are taking whatever JSON on inside of our request body
	//and its gonna bind it to the "newTodo" variable
	//its gonna throw an error if json doesnt have the struct format (id, item, completed)
	//and its gonna be cataching the error inside the "err" var.
	if err := context.BindJSON(&newTodo); err != nil {
		//if error happens, we dont want to continue.
		return
	}

	todos = append(todos, newTodo)

	//return the new "todo"
	context.IndentedJSON(http.StatusCreated, newTodo)

}

// Function that the handler is going to utilize
func getTodo(context *gin.Context) {
	//Extract ID from URL (path parameter)
	id := context.Param("id")
	todo, err := getTodoById((id))

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "ToDo no encontrado."})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)

}

// Iterate over the array and find the correct "todo" wit that specific ID
// This is going to return the todo as well as an error, either one or another
func getTodoById(id string) (*todo, error) {
	//iterate over array and find specific "todo"
	for i, t := range todos {
		if t.ID == id {
			//return the "todo" and the error
			return &todos[i], nil
		}
	}

	//If we don't find it:
	return nil, errors.New("todo not found")
}

// Change "completed" status
func toggleTodoStatus(context *gin.Context) {

	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "ToDo no encontrado."})
		return
	}

	//Toggle status
	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)

}

func main() {
	//Create server
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodo)
	//Set the path
	router.Run("localhost:9090")
}
