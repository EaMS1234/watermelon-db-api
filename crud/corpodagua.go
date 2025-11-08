package crud

import (
	"api/banco"
	"log"

	"encoding/json"
	"net/http"
	"strconv"
)

type Corpo struct{
	Nome string `json:"nome"`
	Tipo string `json:"tipo"`
	ID_Corpo_d_agua int `json:"id"`
	Locais []Localizacao `json:"locais" gorm:"-"`
}

func (Corpo) TableName() string {
	return "Corpo_d_agua"
}


func GetCorpo(w http.ResponseWriter, r *http.Request) {
	log.Output(1, "GET ID_Corpo_d_agua = %v" + r.PathValue("id"))

	db := banco.Banco()
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {panic(err)}

	var corpo Corpo
	db.First(&corpo, "ID_Corpo_d_agua = ?", id)

	// Recebe as localizações
	var corpos_localizacao []Corpo_Localizacao
	db.Where("ID_Corpo_d_agua = ?", id).Find(&corpos_localizacao)

	for _, v := range corpos_localizacao {
		var temp Localizacao
		db.First(&temp, "ID_Localizacao = ?", v.ID_Localizacao)
		corpo.Locais = append(corpo.Locais, temp)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(corpo)
}

func PostCorpo(w http.ResponseWriter, r *http.Request) {
	log.Output(1, "POST Corpo_d_agua")

	db := banco.Banco()

	var corpo Corpo;

	corpo.Nome = r.FormValue("nome")
	corpo.Tipo = r.FormValue("tipo")

	db.Create(&corpo)
}

func GetCorpoTodos(w http.ResponseWriter, r *http.Request) {
	db := banco.Banco()

	var corpos []Corpo
	db.Find(&corpos)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(corpos)
}

