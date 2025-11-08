package main

import (
	"api/crud"

	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Usuário
	mux.HandleFunc("GET /usuario/{id}/", crud.GetUsuario)
	mux.HandleFunc("POST /usuario/", crud.PostUsuario)

	// Localização
	mux.HandleFunc("GET /localizacao/{id}/", crud.GetLocalizacao)

	// Corpo d'água
	mux.HandleFunc("GET /corpo/{id}/", crud.GetUsuario)
	mux.HandleFunc("POST /corpo/", crud.GetUsuario)

	log.Output(0, "Servindo na porta 8080")
	http.ListenAndServe(":8080", mux)
}

