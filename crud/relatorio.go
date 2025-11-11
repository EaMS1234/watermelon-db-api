package crud

import (
	"api/banco"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Relatorio struct{
	ID_relatorio int `json:"id" gorm:"primaryKey"`
	ID_Autor int `json:"autor"`
	ID_Corpo_d_agua int `json:"corpo"`
	Tipo_de_relatorio string `json:"tipo"`
	Data string `json:"data"`
	Descricao string `json:"descricao"`
	Temperatura float64 `json:"temperatura"`
	Cor_Aparente string `json:"cor"`
	Acidez float64 `json:"acidez"`
	Oxigenio_Dissolvido float64 `json:"oxigenio"`
	Demanda_Bioquimica_de_Oxigenio float64 `json:"demanda"`
	Nitrogenio_Total float64 `json:"nitrogenio"`
	Fosforo_Total float64 `json:"fosforo"`
	Metais_Pesados float64 `json:"metais"`
	Cloro_Residual float64 `json:"cloro"`
	Composto_Organico_Volatil float64 `json:"composto"`
	Coliformes float64 `json:"coliformes"`
	Avaliacao_Biologica float64 `json:"avaliacao"`
	Solidos_Totais_Dissolvidos float64 `json:"solidos"`
	Solidos_em_Suspensao float64 `json:"suspensao"`
	Odor string `json:"odor"`
	Sabor string `json:"sabor"`
}

func (Relatorio) TableName() string {
	return "Relatorio"
}

func GetRelatorio(w http.ResponseWriter, r *http.Request) {
	log.Output(0, "GET ID_relatorio = " + r.PathValue("id"))

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {panic(err)}

	var relatorio Relatorio
	banco.Banco().First(&relatorio, id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(relatorio)
}

func PostRelatorio(w http.ResponseWriter, r *http.Request) {
	log.Output(0, "POST Relatorio")

	var relatorio Relatorio

	json.NewDecoder(r.Body).Decode(&relatorio)

	banco.Banco().Create(&relatorio)
}

func DeleteRelatorio(w http.ResponseWriter, r *http.Request) {
	log.Output(0, "DELETE ID_relatorio = " + r.PathValue("id"))

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {panic(err)}

	banco.Banco().Delete(&Relatorio{}, id)	
}

func GetRelatorioTodos(w http.ResponseWriter, r *http.Request) {
	log.Output(0, "GET TODOS Relatorio")

	var relatorios []Relatorio
	banco.Banco().Find(&relatorios)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(relatorios)
}

func PatchRelatorio(w http.ResponseWriter, r *http.Request) {
	log.Output(0, "PATCH ID_relatorio = " + r.PathValue("id"))

	db := banco.Banco()

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {panic(err)}

	var relatorio Relatorio
	db.First(&relatorio, id)

	json.NewDecoder(r.Body).Decode(&relatorio)
	
	db.Save(&relatorio)
}

