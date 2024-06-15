package routes

import (
	"go-auth-task-app/handlers"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("GET", "POST")
	r.HandleFunc("/logout", handlers.LogoutHandler).Methods("GET")
	r.HandleFunc("/register", handlers.RegisterHandler).Methods("GET", "POST")
	// r.HandleFunc("/tasks", handlers.ListTasksHandler).Methods("GET")
	// r.HandleFunc("/tasks/add", handlers.AddTaskHandler).Methods("POST")
	// r.HandleFunc("/tasks/delete", handlers.DeleteTaskHandler).Methods("GET")
	// r.HandleFunc("/tasks/edit", handlers.EditTaskHandler).Methods("GET")
	// r.HandleFunc("/tasks/update", handlers.UpdateTaskHandler).Methods("POST")
	return r
}
