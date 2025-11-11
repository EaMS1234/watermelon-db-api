package crud

import (
	"api/banco"
	"log"

	"encoding/json"
	"net/http"
	"strconv"
)

type Localizacao struct{
	Cidade string `json:"cidade"`
	Estado string `json:"estado"`
	ID_localizacao int `json:"id" gorm:"primaryKey"`
}

func (Localizacao) TableName() string {
	return "Localizacao"
}


func GetLocalizacao(w http.ResponseWriter, r *http.Request) {
	log.Output(1, r.RemoteAddr + " GET ID_localizacao = " + r.PathValue("id"))

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		var local []Localizacao
		if banco.Banco().Where("Estado = ?", r.PathValue("id")).Find(&local).Error != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest) // não foi possível converter em inteiro
			return
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(local)
			return
		}
	}

	var local Localizacao
	if banco.Banco().First(&local, "ID_localizacao = ?", id).Error != nil {
		http.Error(w, "ID inexistente", http.StatusNotFound) // não foi possível encontrar
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(local)
}

func GetLocalizacaoTodos(w http.ResponseWriter, r *http.Request) {
	log.Output(1, r.RemoteAddr + " GET TODOS Localizacao")

	db := banco.Banco()

	var locais []Localizacao
	db.Find(&locais)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(locais)
}

