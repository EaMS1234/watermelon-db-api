package crud

import (
	"api/banco"
	"log"

	"encoding/json"
	"net/http"
	"strconv"
)

type Usuario struct{
	Nome_de_usuario string `json:"nome"`
	E_mail string `json:"email"`
	Senha string `json:"senha"`
	ID_usuario int `json:"id"`
}

func (Usuario) TableName() string {
	return "Usuario"
}


func GetUsuario(w http.ResponseWriter, r *http.Request) {
	log.Output(1, "GET ID_usuario = %v" + r.PathValue("id"))

	db := banco.Banco()
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {panic(err)}

	var usuario Usuario
	db.First(&usuario, "ID_usuario = ?", id)

	usuario.Senha = ""

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usuario)
}

func PostUsuario(w http.ResponseWriter, r *http.Request) {
	log.Output(1, "POST Usuario")

	db := banco.Banco()

	var usuario Usuario;

	usuario.Nome_de_usuario = r.FormValue("nome")
	usuario.E_mail = r.FormValue("email")
	usuario.Senha = r.FormValue("senha")

	db.Create(&usuario)
}

func GetUsuarioTodos(w http.ResponseWriter, r *http.Request) {
	log.Output(1, "GET TODOS Usuario")

	db := banco.Banco()

	var usuarios []Usuario
	db.Find(&usuarios)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usuarios)
}

