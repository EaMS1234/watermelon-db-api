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
	ID_Corpo_d_agua int `json:"id" gorm:"primaryKey"`
}

func (Corpo) TableName() string {
	return "Corpo_d_agua"
}


func GetCorpo(w http.ResponseWriter, r *http.Request) {
	log.Output(1, "GET ID_Corpo_d_agua = " + r.PathValue("id"))

	db := banco.Banco()
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {panic(err)}

	var corpo Corpo
	db.First(&corpo, "ID_Corpo_d_agua = ?", id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(corpo)
}

func PostCorpo(w http.ResponseWriter, r *http.Request) {
	log.Output(1, "POST Corpo_d_agua")

	var corpo Corpo;

	json.NewDecoder(r.Body).Decode(&corpo)

	banco.Banco().Create(&corpo)
}

func DeleteCorpo(w http.ResponseWriter, r *http.Request) {
	log.Output(1, "DELETE ID_Corpo_d_agua = " + r.PathValue("id"))

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {panic(err)}

	banco.Banco().Delete(&Corpo{}, id)
}

func GetCorpoTodos(w http.ResponseWriter, r *http.Request) {
	log.Output(1, "GET TODOS Corpo_d_agua")

	var corpos []Corpo
	banco.Banco().Find(&corpos)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(corpos)
}

func PatchCorpo(w http.ResponseWriter, r *http.Request) {
	log.Output(1, "PATCH ID_Corpo_d_agua = " + r.PathValue("id"))

	db := banco.Banco()

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {panic(err)}

	var corpo Corpo
	db.First(&corpo, id)

	json.NewDecoder(r.Body).Decode(&corpo)

	db.Save(&corpo)
}

