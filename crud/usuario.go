package crud

import (
	"api/banco"
	"fmt"
	"log"

	"encoding/json"
	"net/http"
	"strconv"
)

type Usuario struct{
	Nome_de_usuario string `json:"nome"`
	E_mail string `json:"email"`
	Senha string `json:"senha"`
	ID_usuario int `json:"id" gorm:"primaryKey"`
}

func (Usuario) TableName() string {
	return "Usuario"
}


func GetUsuario(w http.ResponseWriter, r *http.Request) {
	log.Output(1, "GET ID_usuario = " + r.PathValue("id"))

	db := banco.Banco()
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {panic(err)}

	var usuario Usuario
	db.First(&usuario, id)

	usuario.Senha = ""

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usuario)
}

func PostUsuario(w http.ResponseWriter, r *http.Request) {
	log.Output(1, "POST Usuario")

	var usuario Usuario;

	json.NewDecoder(r.Body).Decode(&usuario)

	fmt.Println(usuario)

	banco.Banco().Create(&usuario)
}

func DeleteUsuario(w http.ResponseWriter, r *http.Request) {
	log.Output(1, "DELETE ID_usuario = " + r.PathValue("id"))

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {panic(err)}

	banco.Banco().Delete(&Usuario{}, id)
}

func GetUsuarioTodos(w http.ResponseWriter, r *http.Request) {
	log.Output(1, "GET TODOS Usuario")

	db := banco.Banco()

	var usuarios []Usuario
	db.Find(&usuarios)

	for i := range usuarios {
		usuarios[i].Senha = "";
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usuarios)
}

func PatchUsuario(w http.ResponseWriter, r *http.Request) {
	log.Output(1, "PATCH ID_usuario = " + r.PathValue("id"))

	db := banco.Banco()

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {panic(err)}

	var usuario Usuario
	db.First(&usuario, id)

	json.NewDecoder(r.Body).Decode(&usuario)

	db.Save(&usuario)
}

