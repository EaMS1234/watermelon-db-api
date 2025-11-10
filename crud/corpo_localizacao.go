package crud

import (
	"api/banco"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Corpo_Localizacao struct{
	ID_Corpo_d_agua int `json:"corpo"`
	ID_Localizacao int `json:"localizacao"`
	ID_corpo_localizacao int `json:"id" gorm:"primaryKey"`
}

func (Corpo_Localizacao) TableName() string {
	return "Corpo_Localizacao"
}


func GetCorpoLocal (w http.ResponseWriter, r *http.Request) {
	log.Output(1, "GET Local ID_Corpo_d_agua = " + r.PathValue("id"))

	db := banco.Banco()

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {panic(err)}

	var corpo_local []Corpo_Localizacao
	db.Find(&corpo_local, "ID_Corpo_d_agua = ?", id)

	var locais []Localizacao
	for _, v := range corpo_local {
		var temp Localizacao
		db.First(&temp, v.ID_Localizacao)
		locais = append(locais, temp)
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(locais)
}

func GetLocalCorpo (w http.ResponseWriter, r *http.Request) {
	log.Output(1, "GET Local ID_Localizacao = " + r.PathValue("id"))

	db := banco.Banco()

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {panic(err)}

	var corpo_local []Corpo_Localizacao
	db.Find(&corpo_local, "ID_Localizacao = ?", id)

	var corpos []Corpo
	for _, v := range corpo_local {
		var temp Corpo
		db.First(&temp, v.ID_Corpo_d_agua)
		corpos = append(corpos, temp)
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(corpos)

}

