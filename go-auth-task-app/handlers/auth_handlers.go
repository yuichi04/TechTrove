package handlers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	store = sessions.NewCookieStore([]byte("somthing-very-secret"))
	users = map[string]string{}
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	user := session.Values["user"]
	tmpl, _ := template.ParseFiles("templates/home.html")
	tmpl.Execute(w, user)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, _ := template.ParseFiles("templates/login.html")
		tmpl.Execute(w, nil)
	} else {
		username := r.FormValue("username")
		password := r.FormValue("password")
		if users[username] == password {
			session, _ := store.Get(r, "session")
			session.Values["user"] = username
			session.Save(r, w)
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		}
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	delete(session.Values, "user")
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusFound)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, _ := template.ParseFiles("templates/register.html")
		tmpl.Execute(w, nil)
	} else {
		username := r.FormValue("username")
		password := r.FormValue("password")
		if username != "" && password != "" {
			users[username] = password
			http.Redirect(w, r, "/login", http.StatusFound)
		} else {
			http.Error(w, "Please provide a username and password", http.StatusBadRequest)
		}
	}
}
