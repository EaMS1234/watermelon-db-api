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
	CorposDAgua []Corpo `json:"corpos_d_agua" gorm:"-"`
}

func (Localizacao) TableName() string {
	return "Localizacao"
}


func GetLocalizacao(w http.ResponseWriter, r *http.Request) {
	log.Output(1, "GET ID_localizacao = " + r.PathValue("id"))

	db := banco.Banco()
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {panic(err)}

	var local Localizacao
	db.First(&local, "ID_localizacao = ?", id)

	// Gera uma lista com todos os corpos d'água presentes na localização
	var corpos_localizacao []Corpo_Localizacao
	db.Where("ID_Localizacao = ?", id).Find(&corpos_localizacao)

	for _, v := range corpos_localizacao {
		var temp Corpo
		db.First(&temp, "ID_Corpo_d_agua = ?", v.ID_Corpo_d_agua)
		local.CorposDAgua = append(local.CorposDAgua, temp)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(local)
}

func GetLocalizacaoTodos(w http.ResponseWriter, r *http.Request) {
	log.Output(1, "GET TODOS Localizacao")

	db := banco.Banco()

	var locais []Localizacao
	db.Find(&locais)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(locais)
}

