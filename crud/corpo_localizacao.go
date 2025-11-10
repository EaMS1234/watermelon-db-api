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
	ID_Localizacao int `json:"local"`
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
	log.Output(1, "GET Corpo ID_Localizacao = " + r.PathValue("id"))

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

func PostCorpoLocal (w http.ResponseWriter, r *http.Request) {
	log.Output(1, "POST Local ID_Corpo_d_agua = " + r.PathValue("id"))

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {panic(err)}

	var corpo_local Corpo_Localizacao
	json.NewDecoder(r.Body).Decode(&corpo_local)

	corpo_local.ID_Corpo_d_agua = id

	banco.Banco().Create(&corpo_local)
}

func DeleteCorpoLocal (w http.ResponseWriter, r *http.Request) {
	log.Output(1, "DELETE Local ID_Localizacao " + r.PathValue("id_local") + " para o ID_Corpo_d_agua = " + r.PathValue("id"))

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {panic(err)}

	id_local, err := strconv.Atoi(r.PathValue("id_local"))
	if err != nil {panic(err)}

	banco.Banco().Delete(&Corpo_Localizacao{}, "ID_Corpo_d_agua = ? AND ID_Localizacao = ?", id, id_local)
}

