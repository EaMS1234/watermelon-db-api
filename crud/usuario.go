package crud

import (
	"api/banco"
	"log"

	"encoding/base64"
	"encoding/json"
	"net/http"
	"strconv"
)

type Usuario struct{
	Nome_de_usuario string `json:"nome"`
	E_mail string `json:"email"`
	Senha string `json:"senha"`
	Foto_de_Perfil string `json:"foto"`
	ID_usuario int `json:"id" gorm:"primaryKey"`
}

func (Usuario) TableName() string {
	return "Usuario"
}


func GetUsuario(w http.ResponseWriter, r *http.Request) {
	log.Output(0, "GET ID_usuario = " + r.PathValue("id"))

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {panic(err)}

	var usuario Usuario
	banco.Banco().Select("ID_usuario", "Nome_de_usuario", "E_mail", "Foto_de_Perfil").First(&usuario, id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usuario)
}

func PostUsuario(w http.ResponseWriter, r *http.Request) {
	log.Output(0, "POST Usuario")

	var usuario Usuario;

	json.NewDecoder(r.Body).Decode(&usuario)

	banco.Banco().Create(&usuario)
}

func DeleteUsuario(w http.ResponseWriter, r *http.Request) {
	log.Output(0, "DELETE ID_usuario = " + r.PathValue("id"))

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {panic(err)}

	banco.Banco().Delete(&Usuario{}, id)
}

func GetUsuarioTodos(w http.ResponseWriter, r *http.Request) {
	log.Output(0, "GET TODOS Usuario")

	var usuarios []Usuario
	banco.Banco().Select("ID_usuario", "Nome_de_usuario", "E_mail").Find(&usuarios)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usuarios)
}

func PatchUsuario(w http.ResponseWriter, r *http.Request) {
	log.Output(0, "PATCH ID_usuario = " + r.PathValue("id"))

	db := banco.Banco()

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {panic(err)}

	var usuario Usuario
	db.First(&usuario, id)

	json.NewDecoder(r.Body).Decode(&usuario)

	db.Save(&usuario)
}

func GetUsuarioRelatorio(w http.ResponseWriter, r *http.Request) {
	log.Output(0, "GET Relatorios ID_usuario = " + r.PathValue("id"))

	db := banco.Banco()

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {panic(err)}

	var usuario Usuario
	db.First(&usuario, id)

	var relatorios []Relatorio
	db.Find(&relatorios, "ID_Autor = ?", id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(relatorios)
}

func GetUsuarioFoto (w http.ResponseWriter, r *http.Request) {
	log.Output(0, "GET Foto ID_usuario = " + r.PathValue("id"))

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {panic(err)}

	var usuario Usuario;
	banco.Banco().Select("Foto_de_Perfil").First(&usuario, id)

	imagem, err := base64.StdEncoding.DecodeString(usuario.Foto_de_Perfil)
	if err != nil {panic(err)}

	if usuario.Foto_de_Perfil == "" {
		// Não possui foto de perfil
		w.WriteHeader(http.StatusNotFound)
		return
	}

  w.Header().Set("Content-Type", "image/png")
	w.Write(imagem)
}

