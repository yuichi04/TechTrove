package routes

import (
	"go-task-app/handlers"

	"github.com/gorilla/mux"
)

// NewRouter creates a new mux router and registers the routes
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.ListTasksHandler).Methods("GET")
	r.HandleFunc("/add", handlers.AddTaskHandler).Methods("POST")
	r.HandleFunc("/delete", handlers.DeleteTaskHandler).Methods("GET")
	r.HandleFunc("/edit", handlers.EditTaskHandler).Methods("GET")
	r.HandleFunc("/update", handlers.UpdateTaskHandler).Methods("POST")
	return r
}
