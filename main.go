package main

import (
	"fmt"
	"net/http"
	"time"
)

// Função para verificar o status de uma URL
func checkURL(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Erro ao verificar a URL %s: %v\n", url, err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false
	}

	return true
}

func main() {
	// URLs a serem verificadas
	urls := []string{
		"https://mariovalente.com.br",
		"https://gugaco.com.br",
	}

	for {
		for _, url := range urls {
			if !checkURL(url) {
				mensagem := fmt.Sprintf("A URL %s não está respondendo corretamente!", url)
				fmt.Sprintf(mensagem)

			}
		}

		// Aguardar um intervalo de tempo antes da próxima verificação (por exemplo, 10 minutos)
		time.Sleep(10 * time.Minute)
	}
}
