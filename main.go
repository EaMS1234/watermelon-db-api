package main

import (
	"api/auth"
	"api/crud"

	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /usuario", crud.GetUsuarioTodos)                         // Todos os usuários
	mux.HandleFunc("GET /usuario/{id}", crud.GetUsuario)                         // Usuário por ID
	mux.HandleFunc("POST /usuario", crud.PostUsuario)                            // Adicionar usuário
	mux.HandleFunc("DELETE /usuario/{id}", crud.DeleteUsuario)                   // Remove um usuário por ID
	mux.HandleFunc("PATCH /usuario/{id}", crud.PatchUsuario)                     // Altera um usuário por ID

	mux.HandleFunc("GET /usuario/{id}/foto", crud.GetUsuarioFoto)                // Foto de perfil de um usuário por ID

	mux.HandleFunc("GET /relatorio", crud.GetRelatorioTodos)                     // Todos os relatórios
	mux.HandleFunc("GET /relatorio/{id}", crud.GetRelatorio)                     // Relatório por ID
	mux.HandleFunc("POST /relatorio", crud.PostRelatorio)												 // Adicionar relatório
	mux.HandleFunc("DELETE /relatorio/{id}", crud.DeleteRelatorio)               // Remove um relatório por ID
	mux.HandleFunc("PATCH /relatorio/{id}", crud.PatchRelatorio)                 // Altera um relatório por ID

	mux.HandleFunc("GET /usuario/{id}/relatorio", crud.GetUsuarioRelatorio)      // Todos os relatórios de um usuário

	mux.HandleFunc("GET /local", crud.GetLocalizacaoTodos)                       // Todas as localizações
	mux.HandleFunc("GET /local/{id}", crud.GetLocalizacao)                       // Localização por ID

	mux.HandleFunc("GET /corpo", crud.GetCorpoTodos)                             // Todos os Corpos d'Água
	mux.HandleFunc("GET /corpo/{id}", crud.GetCorpo)                             // Corpos d'Água por ID
	mux.HandleFunc("POST /corpo", crud.PostCorpo)                                // Adicionar Corpo d'Água
	mux.HandleFunc("DELETE /corpo/{id}", crud.DeleteCorpo)                       // Remove um Corpo d'Água por ID
	mux.HandleFunc("PATCH /corpo/{id}", crud.PatchCorpo)                         // Altera um Corpo d'Água por ID

	mux.HandleFunc("GET /local/{id}/corpo", crud.GetLocalCorpo)                  // Todos os corpos d'água de um local
	mux.HandleFunc("GET /corpo/{id}/local", crud.GetCorpoLocal)                  // Todos os locais de um Corpo d'Água
	
	mux.HandleFunc("POST /corpo/{id}/local", crud.PostCorpoLocal)                // Atribui um local a um corpo d'água
	mux.HandleFunc("DELETE /corpo/{id}/local/{id_local}", crud.DeleteCorpoLocal) // Remove um local de um corpo d'água

	mux.HandleFunc("GET /auth", auth.GetAuth)                                    // Status atual da autenticação
	mux.HandleFunc("POST /auth", auth.PostAuth)                                  // Autenticar
	mux.HandleFunc("DELETE /auth", auth.DeleteAuth)                              // Deslogar

	log.Output(0, "Servindo na porta 8080")
	http.ListenAndServe(":8080", mux)
}

