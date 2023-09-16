package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Username string
	Contact  string
}

var users = []User{
	{"David", "(81) 98434-1340"},
	{"José", "(81) 98321-3232"},
	{"Carla", "(87) 99156-7843"},
	{"Julia", "(28) 98642-9091"},
}

func main() {
	// configura a pasta que irá servir os arquivos estáticos de estilo
	static := http.FileServer(http.Dir("static"))

	// seta que as chamadas para a rota /static/ irão devolver os arquivos estáticos
	// desse diretório
	http.Handle("/static/", http.StripPrefix("/static/", static))

	// configura o handlers
	http.HandleFunc("/", createHandler("public/index.html"))
	http.HandleFunc("/usuarios", createHandler("public/users.html"))
	http.HandleFunc("/contatos", createHandler("public/contacts.html"))

	// inicia o servidor
	http.ListenAndServe(":3030", nil)
}

// 3_221_225_530

type Handler func(http.ResponseWriter, *http.Request)

func createHandler(route string) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(route)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, users)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
