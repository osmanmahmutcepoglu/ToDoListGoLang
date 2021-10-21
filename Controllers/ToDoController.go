package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
)

var (
	MsgTodoAdded              = "Yapılacak iş başarıyla eklendi."
	ErrTodoDescriptionIsEmpty = errors.New("Yapılacak işi giriniz.")
)

type ToDoList struct {  // Listeleme ve İndisleme işlemi için ToDoList Nesnesini oluşturuyoruz.
	Id          int    `json:"id"`
	Description string `json:"description"`
}

var toDoMemoryDatabase map[int]ToDoList //Veritabanı olarak bellek hafızasını kullanıyoruz.



type TodoListInput struct { //İnput'tan gelen değer için nesne oluşturuyoruz.
	TaskDescription string `json:"task_description" binding:"required"`
}

type TodoController interface { //TodoList controller'ımızın interfacesi
	GetTodoList(c *gin.Context)
	AddTodo(c *gin.Context)
}

type todoController struct { //Oluşturduğumuz ControllerBase'nin referansını tutuyoruz.
	base ControllerBase
}

func NewTodoController(basecontroller ControllerBase) TodoController {
	toDoMemoryDatabase = make(map[int]ToDoList, 0)
	return &todoController{
		base: basecontroller,
	}
}

// localhost:8080/api/v1/getTodoList
func (t todoController) GetTodoList(c *gin.Context) {
	todoList := make([]ToDoList, 0, len(toDoMemoryDatabase))

	for _, todo := range toDoMemoryDatabase {
		todoList = append(todoList, todo)
	}

	sort.Slice(todoList, func(i, j int) bool {
		return todoList[i].Id < todoList[j].Id
	})

	t.base.Data(c, http.StatusOK, todoList, "")
}

//localhost:8080/api/v1/addTodo
func (t todoController) AddTodo(c *gin.Context) {
	var todoListInput TodoListInput

	if err := c.ShouldBindJSON(&todoListInput); err != nil {
		t.base.Error(c, http.StatusBadRequest, ErrTodoDescriptionIsEmpty)
		return
	}

	todoId := len(toDoMemoryDatabase)
	addedTodo := ToDoList{
		Id:          todoId,
		Description: todoListInput.TaskDescription,
	}
	toDoMemoryDatabase[todoId] = addedTodo

	t.base.Data(c, http.StatusCreated, addedTodo, MsgTodoAdded)
}

