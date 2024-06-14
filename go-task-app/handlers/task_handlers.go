package handlers

import (
	"go-task-app/models"
	"html/template"
	"net/http"
	"strconv"
	"sync"
)

var (
	tasks  []models.Task
	nextID int
	mu     sync.Mutex
)

func ListTasksHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/list.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, tasks)
}

func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		title := r.FormValue("title")
		if title != "" {
			mu.Lock()
			nextID++
			tasks = append(tasks, models.Task{ID: nextID, Title: title})
			mu.Unlock()
		}
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func EditTaskHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	// Find the task with the given ID
	var taskToEdit *models.Task
	for i := range tasks {
		if tasks[i].ID == id {
			taskToEdit = &tasks[i]
			break
		}
	}
	if taskToEdit == nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/edit.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, taskToEdit)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		idStr := r.FormValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid task ID", http.StatusBadRequest)
			return
		}

		title := r.FormValue("title")
		if title == "" {
			http.Error(w, "Title cannot be empty", http.StatusBadRequest)
			return
		}

		mu.Lock()
		defer mu.Unlock()
		for i := range tasks {
			if tasks[i].ID == id {
				tasks[i].Title = title
				break
			}
		}
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
