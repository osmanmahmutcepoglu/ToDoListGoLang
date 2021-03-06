package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTodoController(t *testing.T) {

	toDoListController := NewTodoController(&controllerBase{})
	gin.SetMode(gin.TestMode)

	t.Run("addTodo", func(t *testing.T) {
		t.Run("Error when sending empty body", func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			toDoListController.AddTodo(c)

			assert.Equal(t, http.StatusBadRequest, w.Code)
    
			resBody := Attachment{}
			json.NewDecoder(w.Body).Decode(&resBody)

			assert.EqualValues(t, ErrTodoDescriptionIsEmpty.Error(), resBody.Message)
		})
		t.Run("Added todo successfully", func(t *testing.T) {
			dummyTaskDescription := "dummy todo"
			reqBody := TodoListInput{TaskDescription: dummyTaskDescription}

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			payload, _ := json.Marshal(&reqBody)
			request := httptest.NewRequest(http.MethodPost, "/api/v1/addTodo", bytes.NewBuffer(payload))
			c.Request = request

			toDoListController.AddTodo(c)

			assert.Equal(t, http.StatusCreated, w.Code)

			resBody := Attachment{}
			json.NewDecoder(w.Body).Decode(&resBody)

			assert.Equal(t, MsgTodoAdded, resBody.Message)
			assert.Equal(t, toDoMemoryDatabase[0].Id, 0)
			assert.Equal(t, toDoMemoryDatabase[0].Description, dummyTaskDescription)
		})
	})

	t.Run("getTodoList", func(t *testing.T) {
		t.Run("Get empty list when no todo is available", func(t *testing.T) {
			toDoMemoryDatabase = make(map[int]ToDoList, 0)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			toDoListController.GetTodoList(c)

			assert.Equal(t, http.StatusOK, w.Code)

			resBody := Attachment{}
			json.NewDecoder(w.Body).Decode(&resBody)

			assert.Empty(t, resBody.Data)
		})
		t.Run("Get todos ascending order by their ids", func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			for i := 0; i < 5; i++ {
				toDoMemoryDatabase[i] = ToDoList{
					Id:          i,
					Description: fmt.Sprintf("Dummy-Task-%d", i),
				}
			}

			toDoListController.GetTodoList(c)

			assert.Equal(t, http.StatusOK, w.Code)

			resBody := Attachment{}
			json.NewDecoder(w.Body).Decode(&resBody)

			for i := 0; i < 5; i++ {
				val, _ := resBody.Data.([]interface{})
				cvr, convertOk := val[i].(map[string]interface{})

				firstId := cvr["id"]
				firstDescription := cvr["description"]

				assert.True(t, convertOk)
				assert.EqualValues(t, i, firstId)
				assert.Equal(t, fmt.Sprintf("Dummy-Task-%d", i), firstDescription)
			}
		})
	})
}
