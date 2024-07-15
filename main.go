package main

import (
	"log"
	"net/http"

	api "api-go/api"
)

func main() {
	http.HandleFunc("/extract-cnh", api.HandleExtractCNH)
	http.HandleFunc("/extract-rg", api.HandleExtractRG)
	http.HandleFunc("/extract-passport", api.HandleExtractPassport)

	port := ":8080"
	log.Printf("API rodando na porta %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
