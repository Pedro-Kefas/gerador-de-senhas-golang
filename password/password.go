package password

//Usando o pacote math/rand e time do go para conseguir gerar senhas de forma aleatória

import (
	"fmt"
	"math/rand"
	"time"
)

func GeneratePassword(tamanho int) string {
	caracteres := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()+?><:{}[]$€£¥₽¢₩₮₦₲œþðøåäö"
	rand.Seed(time.Now().UnixNano()) //Isso garante que a senha mude a cada execução

	senha := ""
	for i := 0; i < tamanho; i++ {
		indice := rand.Intn(len(caracteres)) // isso vai pegar um índice aleatório
		senha += string(caracteres[indice])  // adiciona o caractere que corresponde à senha
	}
	fmt.Println("Senha especial aleatória gerada:", senha)
	return senha
}
