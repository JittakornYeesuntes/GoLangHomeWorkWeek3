package main 

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Status string `json:"status"`
}

var todos = make(map[string]*Todo)

func postTodosHandler(c *gin.Context) {
	t := Todo{}

	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	i := len(todos)
	i++
	id := strconv.Itoa(i)
	t.ID = id
	todos[id] = &t

	c.JSON(http.StatusCreated, todos[id])
}

func getTodoByIDHandler(c *gin.Context) {
	id := c.Param("id")
	t, ok := todos[id]
	if !ok {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	c.JSON(http.StatusOK, t)
}

func getTodoHandler(c *gin.Context) {
	allTodos := []*Todo{}
	for _, item := range todos {
		allTodos = append(allTodos, item)
	}

	c.JSON(http.StatusOK, allTodos)
}

func putTodoByIDHandler(c *gin.Context) {
	id := c.Param("id")
	t := Todo{}

	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	t.ID = id
	todos[id] = &t
	c.JSON(http.StatusOK, todos[id])
}

func deleteTodoByIDHandler(c *gin.Context) {
	id := c.Param("id")
	delete(todos, id)
	c.JSON(http.StatusOK, gin.H{
		"status" : "success",
	})
}


func main() {
	router := gin.Default()

	group := router.Group("/api")
	group.POST("/todos", postTodosHandler)
	group.GET("/todos/:id", getTodoByIDHandler)
	group.GET("/todos", getTodoHandler)
	group.PUT("/todos/:id", putTodoByIDHandler)
	group.DELETE("/todos/:id", deleteTodoByIDHandler)

	router.Run(":1234")
}