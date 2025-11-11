package auth

import (
	"api/banco"
	"api/crud"

	"crypto/rand"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

var sessoes_global map[string]int = make(map[string]int)

func Validar (w http.ResponseWriter, r *http.Request) bool {
	cookie_token, err := r.Cookie("token")
	if err != nil {
		log.Output(1, "Não possui token")
		return false
	}

	token := cookie_token.Value

	if sessoes_global[token] == 0 {
		log.Output(1, "Token inválido")
		return false
	}

	log.Output(1, token + " autenticado com sucesso")
	return true
}


func GetAuth (w http.ResponseWriter, r *http.Request) {
	log.Output(1, "GET Auth")

	if Validar(w, r) {
		w.WriteHeader(http.StatusOK)		
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func PostAuth (w http.ResponseWriter, r *http.Request) {
	log.Output(1, "POST Auth")

	var form struct{
		Email string `json:"email"`
		Senha string `json:"senha"`
		Manter bool `json:"manter"`
	}

	json.NewDecoder(r.Body).Decode(&form)

	var usuario crud.Usuario
	banco.Banco().Select("E_mail", "Senha", "ID_usuario").First(&usuario, "E_mail = ?", form.Email)

	if usuario.Senha == form.Senha {
		// Login com sucesso


		// Cria o token da sessão
		var token string

		// Gera um token novo até que ele seja único
		for {
			token = rand.Text()

			if _, res := sessoes_global[token]; !res {
				break
			}
		}

		// Salva a sessão localmente
		sessoes_global[token] = usuario.ID_usuario

		// Cria os cookies
		cookie_token := http.Cookie {
			Name: "token",
			Value: token,
			Path: "/",
			HttpOnly: true,
		}

		if form.Manter {
			validade := time.Now().AddDate(0, 3, 0)

			// Vence em um mês a partir de agora
			cookie_token.Expires = validade
		}

		log.Output(1, "Sessão: " + token + " ID_usuario = " + strconv.Itoa(sessoes_global[token]))

		// Salva os cookies
		http.SetCookie(w, &cookie_token)

		w.WriteHeader(http.StatusOK)
		return

	} else {
		http.Error(w, "Senha incorreta", http.StatusBadRequest)
		return
	}
}

func DeleteAuth (w http.ResponseWriter, r *http.Request) {

}

