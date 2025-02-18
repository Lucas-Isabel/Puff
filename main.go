package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "https://raw.githubusercontent.com/Lucas-Isabel/Puff/main/useControl.txt"

	// Faz a requisição HTTP GET
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Erro ao fazer a requisição:", err)
		return
	}
	defer resp.Body.Close()

	// Lê o conteúdo do corpo da resposta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler o corpo da resposta:", err)
		return
	}

	// Converte o conteúdo para string e remove espaços em branco
	content := strings.TrimSpace(string(body))

	// Verifica se o conteúdo é igual a "true"
	if content == "true" {
		fmt.Println("O conteúdo do arquivo é 'true'.")
	} else {
		fmt.Println("O conteúdo do arquivo não é 'true'.")
	}
}
