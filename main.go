package main

import (
	"desafios/password"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	// Iniciar servidor web na porta 8080
	fmt.Println("Servidor rodando em http://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func handler(w http.ResponseWriter, _ *http.Request) {
	senha := password.GeneratePassword(60) //função handler
	//acrescentando o html básico para criar uma "interface" de teste simples

	//Código HTML5
	_, err := fmt.Fprintf(w, `
			<html>
				<head><title>Gerador de senha</title></head>
				<body>
					<h1>Senha Gerada:</h1>
					<p><strong>%s</strong></p>
					<form method="get">
						<button type="submit">Gerar senha!</button>
					</form>
				</body>
			</html>
		`, senha)
	if err != nil {
		return
	}
}
