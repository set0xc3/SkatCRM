package web

import (
	"log"
	"net/http"
)

type Todo struct {
	ID    int
	Title string
	Done  bool
}

var todos = []Todo{
	{1, "Learn Go", false},
	{2, "Build a Todo App", false},
}

func TodoListHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	component := TodoList()
	err = component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in HelloWebHandler: %e", err)
	}
}
