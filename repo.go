package main

import (
	"fmt"
	"time"
)

var currentId int

var todos Todos

// Give us some seed data
func init() {
	RepoCreateTodo(Todo{Name: "Write presentation", Due: time.Now()})
	RepoCreateTodo(Todo{Name: "Host meetup", Due: time.Now()})
}

func RepoFindTodo(id int) Todo {
	for _, todo := range todos {
		if todo.Id == id {
			return todo
		}
	}
	// return empty if not found
	return Todo{}
}

func RepoCreateTodo(todo Todo) Todo {
	currentId += 1
	todo.Id = currentId
	todos = append(todos, todo)
	return todo
}

func RepoDestroyTodo(id int) error {
	for i, todo := range todos {
		if todo.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Could not find Todo id of %d to delete", id)
}
