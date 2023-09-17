package main

import (
	"log"
	"net/http"

	"github.com/DavidEsdrs/template-server/connections"
	"github.com/DavidEsdrs/template-server/handlers"
	"github.com/gorilla/mux"
)

type User struct {
	Username string
	Contact  string
}

// var users = []User{
// 	{"David", "(81) 98434-1340"},
// 	{"José", "(81) 98321-3232"},
// 	{"Carla", "(87) 99156-7843"},
// 	{"Julia", "(28) 98642-9091"},
// }

func main() {
	_, err := connections.Connect()

	if err != nil {
		log.Fatal("error while connecting with database")
	}

	r := mux.NewRouter()

	// configura o handlers
	// r.HandleFunc("/", createHandler("public/index.html", users))
	r.HandleFunc("/usuarios", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/usuarios/criar", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/usuarios/criacao", handlers.CreateUserTemplate).Methods("GET")
	r.HandleFunc("/usuarios/{userID:[0-9]+}", handlers.GetUser)
	// r.HandleFunc("/contatos", createHandler("public/contacts.html", users))

	// configura a pasta que irá servir os arquivos estáticos de estilo
	static := http.FileServer(http.Dir("static"))

	// seta que as chamadas para a rota /static/ irão devolver os arquivos estáticos
	// desse diretório
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", static))

	// Create an HTTP server with the router
	server := &http.Server{
		Addr:    ":3030",
		Handler: r, // Set the router as the handler
	}

	// Start the server
	log.Fatal(server.ListenAndServe())
}
