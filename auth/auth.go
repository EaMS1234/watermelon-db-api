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

type sessao struct {
	ID_usuario int
	Validade time.Time
}

var sessoes_global map[string]sessao = make(map[string]sessao)


func Validar (w http.ResponseWriter, r *http.Request) bool {
	cookie_token, err := r.Cookie("token")
	if err != nil {
		log.Output(1, "Não possui token")
		return false
	}

	cookie_usuario, err := r.Cookie("usuario")
	if err != nil {
		log.Output(1, "Não possui ID")
		return false
	}

	token := cookie_token.Value
	id, err := strconv.Atoi(cookie_usuario.Value)
	if err != nil {
		log.Output(1, "ID não é um número")
		return false
	}

	if sessoes_global[token].ID_usuario != id {
		log.Output(1, "ID inválido")
		return false
	}

	if sessoes_global[token].Validade.Before(time.Now()) {
		log.Output(1, "Token expirou")
		return false
	}

	log.Output(1, "Autenticado com sucesso!")
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


		// Cria os cookies
		cookie_token := http.Cookie {
			Name: "token",
			Value: token,
			Path: "/",
		}

		cookie_usuario := http.Cookie {
			Name: "usuario",
			Value: strconv.Itoa(usuario.ID_usuario),
			Path: "/",
		}


		if form.Manter {
			// Vence em um mês a partir de agora
			sessoes_global[token] = sessao{usuario.ID_usuario, time.Now().AddDate(0, 1, 0)}
			cookie_token.Expires = sessoes_global[token].Validade
			cookie_usuario.Expires = sessoes_global[token].Validade
		} else {
			sessoes_global[token] = sessao{usuario.ID_usuario, time.Now()}
		}


		log.Output(1, "Sessão: " + token + " ID_usuario = " + strconv.Itoa(sessoes_global[token].ID_usuario))


		// Salva os cookies
		http.SetCookie(w, &cookie_token)
		http.SetCookie(w, &cookie_usuario)

		w.WriteHeader(http.StatusOK)
		return

	} else {
		http.Error(w, "Senha incorreta", http.StatusBadRequest)
		return
	}
}

func DeleteAuth (w http.ResponseWriter, r *http.Request) {

}

