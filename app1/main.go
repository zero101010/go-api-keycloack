package main

import (
	"encoding/json"
	"log"
	"net/http"

	oidc "github.com/coreos/go-oidc"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

// TODO criar em envs
// Dados do keycloack
var (
	clientID     = "app"
	clientSecret = "c6c5a314-8968-4ac9-aac9-db3c35228273"
)

// func index(w http.ResponseWriter, r *http.Request) {
// 	state := "example"
// 	http.Redirect(w, r, config.AuthCodeURL(state), http.StatusFound)
// }

func main() {
	ctx := context.Background()
	// provider responsável pela autenticação do keycloack
	provider, err := oidc.NewProvider(ctx, "http://localhost:8080/auth/realms/demo")

	if err != nil {
		log.Fatal(err)
	}

	// configurações que serão setadas para o keycloack usando o oauth2
	config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  "http://localhost:8081/auth/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email", "roles"},
	}
	// Chave única para decode
	state := "teste"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		http.Redirect(w, r, config.AuthCodeURL(state), http.StatusFound)
	})

	http.HandleFunc("/auth/callback", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("state") != state {
			http.Error(w, "State doesn't match", http.StatusBadRequest)
			return
		}
		// Troca de authorizationToken para accessToken dentro do nosso servidor
		oauth2Token, err := config.Exchange(ctx, r.URL.Query().Get("code"))
		if err != nil {
			http.Error(w, "Problem to get accessToken", http.StatusInternalServerError)
			return
		}
		rawIDToken, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			http.Error(w, "Problem to get accessToken", http.StatusInternalServerError)
			return
		}

		// Cria o objeto de resposta para dar o MARSHAL
		res := struct {
			OAuth2Token *oauth2.Token
			IDToken     string
		}{
			oauth2Token, rawIDToken,
		}
		data, err := json.MarshalIndent(res, "", "  ")
		w.Write(data)

	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
