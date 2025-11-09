package main

import (
	"api/crud"

	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /usuario/", crud.GetUsuarioTodos)      // Todos os usuários
	mux.HandleFunc("GET /usuario/{id}/", crud.GetUsuario)      // Usuário por ID
	mux.HandleFunc("POST /usuario", crud.PostUsuario)          // Adicionar usuário
	mux.HandleFunc("DELETE /usuario/{id}", crud.DeleteUsuario) // Remove um usuário por ID

	mux.HandleFunc("GET /local/", crud.GetLocalizacaoTodos)    // Todas as localizações
	mux.HandleFunc("GET /local/{id}/", crud.GetLocalizacao)    // Localização por ID

	mux.HandleFunc("GET /corpo/", crud.GetCorpoTodos)          // Todos os Corpos d'Água
	mux.HandleFunc("GET /corpo/{id}/", crud.GetCorpo)          // Corpos d'Água por ID
	mux.HandleFunc("POST /corpo", crud.PostCorpo)              // Adicionar Corpo d'Água
	mux.HandleFunc("DELETE /corpo/{id}", crud.DeleteCorpo)     // Remove um Corpo d'Água por ID

	log.Output(0, "Servindo na porta 8080")
	http.ListenAndServe(":8080", mux)
}

