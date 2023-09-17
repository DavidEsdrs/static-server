package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/DavidEsdrs/template-server/connections"
	"github.com/DavidEsdrs/template-server/crud"
	"github.com/DavidEsdrs/template-server/models"
	"github.com/gorilla/mux"
)

type Handler func(http.ResponseWriter, *http.Request)

func createHandler(route string, ctx any) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(route)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]
	id, _ := strconv.Atoi(userID)
	user, err := crud.GetUser(id)
	if err != nil {
		createHandler("public/not-found.html", nil)(w, r)
		return
	}
	createHandler("public/user.html", user)(w, r)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	user, err := crud.GetUsers()
	if err != nil {
		createHandler("public/not-found.html", nil)(w, r)
		return
	}
	createHandler("public/users.html", user)(w, r)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Corpo da requisição inválido", http.StatusBadRequest)
		return
	}

	db, err := connections.Connect()

	if err != nil {
		http.Error(w, "Erro interno", http.StatusInternalServerError)
		return
	}

	result := db.Create(&user)

	if result.Error != nil {
		http.Error(w, "Erro criando registro", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Usuário criado")
}

func CreateUserTemplate(w http.ResponseWriter, r *http.Request) {
	createHandler("public/criar-usuario.html", nil)(w, r)
}
