package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	auth "github.com/abbot/go-http-auth"
)
// Função para retornar a senha do usuário
// site gerador de senha crypt: https://unix4lyfe.org/crypt/
func secret(user, realm string) string {
	if user == "Roberton3333" {
		// Retorna a senha criptografada do usuário
        // A senha é "golang"
		return "$1$yGWx65Pd$5A0PEku/fJ5TWbjGmCM1U."
	}
	return ""
}
// Verifica se o número de argumentos é válido
func main() {
	if len(os.Args) != 3 {
		fmt.Println("Uso: go run main.go <diretorio> <porta>")
		os.Exit(1)
	}
	    // Lê o diretório e a porta a partir dos argumentos da linha de comando
	httpDir := os.Args[1]
	porta := os.Args[2]
	
    // Cria um autenticador básico usando a função secret para verificar as credenciais do usuário

	authenticator := auth.NewBasicAuthenticator("localhost", secret)
	    // Cria um manipulador de autenticação que envolve o manipulador de arquivos
    // Isso garante que o usuário precise fazer login antes de acessar os arquivos

	http.Handle("/", authenticator.Wrap(func (w http.ResponseWriter, r *auth.AuthenticatedRequest) {
		    // Cria um manipulador de arquivos para o diretório especificado
		http.FileServer(http.Dir(httpDir)).ServeHTTP(w, &r.Request)
	}))  
	    // Inicia o servidor na porta especificada

	fmt.Printf("Subindo servidor na porta %s ...", porta)
	log.Fatal(http.ListenAndServe(":"+porta, nil))
		
}
