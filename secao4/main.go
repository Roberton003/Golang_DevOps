package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"time"
	
)

type Server struct {
	ServerName    string
	ServerURL     string
	tempoExecucao float64
	status        int
	dataFalha	 string
}

func criarListaServidores(serverList *os.File) []Server {
	csvReader := csv.NewReader(serverList)
	data, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var servidores []Server
	for i, line := range data {
		if i > 0 {
			servidor := Server{
				ServerName: line[0],
				ServerURL:  line[1],
			}
			servidores = append(servidores, servidor)
		}
	}
	return servidores
}

func checkServer(servidores []Server) []Server {
	var downServers []Server
	// Itera sobre a lista de servidores
	for _, servidor := range servidores {
		    // Faz uma requisição GET para o servidor
		agora := time.Now()
		    // Se houver um erro na requisição, marca o servidor como inativo e adiciona à lista de servidores inativos
		get, err := http.Get(servidor.ServerURL)
		if err != nil {
			fmt.Println("Server %s is down[%s]\n", servidor.ServerName, err.Error())
			servidor.status = 0
			servidor.dataFalha = agora.Format("02/01/2006 15:04:05")
			downServers = append(downServers,servidor)
			continue
		}
		    // Se o status da resposta não for 200, marca o servidor como inativo e adiciona à lista de servidores inativos
		servidor.status = get.StatusCode
		if servidor.status != 200 {
			servidor.dataFalha = agora.Format("02/01/2006 15:04:05")
			downServers = append(downServers,servidor)
		}
		    // Calcula o tempo de execução da requisição e exibe o status, tempo de carga e URL do servidor
		servidor.tempoExecucao = time.Since(agora).Seconds()
		fmt.Printf("Status: [%d] Tempo de carga: [%f] URL: [%s]\n", servidor.status, servidor.tempoExecucao, servidor.ServerURL)
		
	}
	// Retorna a lista de servidores inativos
	return downServers
}
// Abre os arquivos de lista de servidores e de tempo de inatividade
func openFiles(serverListFile string, downtimefile string) (*os.File, *os.File) {
	
	serverList, err := os.OpenFile(serverListFile, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	    // Abre o arquivo de tempo de inatividade para escrita, criando o arquivo se ele não existir
	downtimeList, err := os.OpenFile(downtimefile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	    // Retorna os arquivos abertos
	return serverList, downtimeList
}
// Gera um arquivo CSV com a lista de servidores inativos
func generateDowntime(downtimefile *os.File, downservers []Server) {
	    // Cria um novo escritor CSV para o arquivo de tempo de inatividade
	csvWriter := csv.NewWriter(downtimefile)
	    // Itera sobre a lista de servidores inativos
	for _, servidor := range downservers {
		        // Cria uma nova linha para o arquivo CSV com as informações do servidor inativo
		line := []string{servidor.ServerName, servidor.ServerURL, servidor.dataFalha, fmt.Sprintf("%f", servidor.tempoExecucao), fmt.Sprintf("%d", servidor.status)}
		        // Escreve a linha no arquivo CSV
		csvWriter.Write(line)
	}
	        // Descarrega quaisquer dados em buffer para o arquivo CSV
csvWriter.Flush()
}
// Função principal do programa	
func main() {
	    // Abre os arquivos de lista de servidores e de tempo de inatividade
	serverList, downtimeList := openFiles(os.Args[1], os.Args[2])
	    // Fecha os arquivos quando a função main() terminar
	defer serverList.Close()
	defer downtimeList.Close()
	    // Verifica quais servidores estão inativos
	servidores := criarListaServidores(serverList)	
	    // Gera um arquivo CSV com a lista de servidores inativos
	downServers := checkServer(servidores)
	generateDowntime(downtimeList, downServers)
}

