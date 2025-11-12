package crud

import (
	"api/auth"
	"api/banco"
	"log"
	"strings"

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
	log.Output(1, r.RemoteAddr + " GET ID_usuario = " + r.PathValue("id"))

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest) // não foi possível converter em inteiro
		return
	}

	var usuario Usuario
	if banco.Banco().Select("ID_usuario", "Nome_de_usuario", "E_mail", "Foto_de_Perfil").First(&usuario, id).Error != nil {
		http.Error(w, "ID inexistente", http.StatusNotFound) // ID não existe
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usuario)
}

func PostUsuario(w http.ResponseWriter, r *http.Request) {
	log.Output(1, r.RemoteAddr + " POST Usuario")

	var usuario Usuario;

	if json.NewDecoder(r.Body).Decode(&usuario) != nil {
		http.Error(w, "Campo inválido", http.StatusBadRequest)
		return
	}

	if err := banco.Banco().Create(&usuario).Error; err != nil {
		if strings.Contains(err.Error(), "1062") {
			http.Error(w, "E-Mail já existe", http.StatusConflict) // único campo unique é o E-Mail
			return
		}
	}
}

func DeleteUsuario(w http.ResponseWriter, r *http.Request) {
	log.Output(1, r.RemoteAddr + " DELETE ID_usuario = " + r.PathValue("id"))

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest) // não foi possível converter em inteiro
		return
	}

	id_sessao, err := auth.Validar(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if id_sessao != id {
		http.Error(w, "Usuário inválido", http.StatusUnauthorized)
		return
	}

	if banco.Banco().Delete(&Usuario{}, id).Error != nil {
		http.Error(w, "ID inexistente", http.StatusNotFound) // não foi possível encontrar no banco de dados
		return
	}
}

func GetUsuarioTodos(w http.ResponseWriter, r *http.Request) {
	log.Output(1, r.RemoteAddr + " GET TODOS Usuario")

	var usuarios []Usuario
	banco.Banco().Select("ID_usuario", "Nome_de_usuario", "E_mail").Find(&usuarios)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usuarios)
}

func PatchUsuario(w http.ResponseWriter, r *http.Request) {
	log.Output(1, r.RemoteAddr + " PATCH ID_usuario = " + r.PathValue("id"))

	db := banco.Banco()

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest) // não foi possível converter em inteiro
		return
	}

	id_sessao, err := auth.Validar(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if id_sessao != id {
		http.Error(w, "Usuário inválido", http.StatusUnauthorized)
		return
	}

	var usuario Usuario
	db.First(&usuario, id)

	if json.NewDecoder(r.Body).Decode(&usuario) != nil {
		http.Error(w, "Campo inválido", http.StatusBadRequest)
		return
	}

	db.Save(&usuario)
}

func GetUsuarioRelatorio(w http.ResponseWriter, r *http.Request) {
	log.Output(1, r.RemoteAddr + " GET Relatorios ID_usuario = " + r.PathValue("id"))

	db := banco.Banco()

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest) // não foi possível converter em inteiro
		return
	}

	var usuario Usuario
	if db.First(&usuario, id).Error != nil {
		http.Error(w, "ID inexistente", http.StatusNotFound)
		return
	}

	var relatorios []Relatorio
	if db.Find(&relatorios, "ID_Autor = ?", id).Error != nil {
		http.Error(w, "Não há relatórios", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(relatorios)
}

func GetUsuarioFoto (w http.ResponseWriter, r *http.Request) {
	log.Output(1, r.RemoteAddr + " GET Foto ID_usuario = " + r.PathValue("id"))

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest) // não foi possível converter em inteiro
		return
	}

	var usuario Usuario;
	banco.Banco().Select("Foto_de_Perfil").First(&usuario, id)

	imagem, err := base64.StdEncoding.DecodeString(usuario.Foto_de_Perfil)
	if err != nil {
		http.Error(w, "Não foi possível obter a imagem", http.StatusInternalServerError)
		return
	}

	if usuario.Foto_de_Perfil == "" {
		// Não possui foto de perfil
		http.Error(w, "Não foi possível encontrar a foto de perfil", http.StatusNotFound)
		return
	}

  w.Header().Set("Content-Type", "image/png")
	w.Write(imagem)
}

