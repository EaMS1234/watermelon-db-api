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
	ID_localizacao int `json:"id"`
}

func (Localizacao) TableName() string {
	return "Localizacao"
}


func GetLocalizacao(w http.ResponseWriter, r *http.Request) {
	log.Output(1, "GET ID_localizacao = %v" + r.PathValue("id"))

	db := banco.Banco()
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {panic(err)}

	var local Localizacao
	db.First(&local, "ID_localizacao = ?", id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(local)
}

