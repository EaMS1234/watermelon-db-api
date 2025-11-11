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
	log.Output(1, r.RemoteAddr + " GET ID_Corpo_d_agua = " + r.PathValue("id"))

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest) // não foi possível converter em inteiro
		return
	}

	var corpo Corpo
	if banco.Banco().First(&corpo, "ID_Corpo_d_agua = ?", id).Error != nil {
		http.Error(w, "ID inexistente", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(corpo)
}

func PostCorpo(w http.ResponseWriter, r *http.Request) {
	log.Output(1, r.RemoteAddr + " POST Corpo_d_agua")

	var corpo Corpo;

	if json.NewDecoder(r.Body).Decode(&corpo) != nil {	
		http.Error(w, "Campo Inválido", http.StatusBadRequest)
		return
	}

	banco.Banco().Create(&corpo) // Não precisa verificar por duplicatas
}

func DeleteCorpo(w http.ResponseWriter, r *http.Request) {
	log.Output(1, r.RemoteAddr + " DELETE ID_Corpo_d_agua = " + r.PathValue("id"))

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusNotFound)
		return
	}

	if banco.Banco().Delete(&Corpo{}, id).Error != nil {
		http.Error(w, "ID inexistente", http.StatusNotFound)
		return
	}
}

func GetCorpoTodos(w http.ResponseWriter, r *http.Request) {
	log.Output(1, r.RemoteAddr + " GET TODOS Corpo_d_agua")

	var corpos []Corpo
	banco.Banco().Find(&corpos)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(corpos)
}

func PatchCorpo(w http.ResponseWriter, r *http.Request) {
	log.Output(1, r.RemoteAddr + " PATCH ID_Corpo_d_agua = " + r.PathValue("id"))

	db := banco.Banco()

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusNotFound)
		return
	}

	var corpo Corpo
	db.First(&corpo, id)

	if json.NewDecoder(r.Body).Decode(&corpo) != nil {
		http.Error(w, "Campo Inválido", http.StatusBadRequest)
		return
	}

	db.Save(&corpo)
}

