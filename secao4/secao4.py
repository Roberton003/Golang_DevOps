import csv
import http.client
import os
import sys
import time

class Server:
    def __init__(self, server_name, server_url, status=0, data_falha="", tempo_execucao=0):
        self.server_name = server_name
        self.server_url = server_url
        self.status = status
        self.data_falha = data_falha
        self.tempo_execucao = tempo_execucao

def open_files(server_list_file, downtime_file):
    # Abre o arquivo de lista de servidores para leitura
    server_list = open(server_list_file, "r")
    # Abre o arquivo de tempo de inatividade para escrita, criando o arquivo se ele não existir
    downtime_list = open(downtime_file, "a+")
    # Retorna os arquivos abertos
    return server_list, downtime_list

def generate_downtime(downtime_file, down_servers):
    # Cria um novo escritor CSV para o arquivo de tempo de inatividade
    csv_writer = csv.writer(downtime_file)
    # Itera sobre a lista de servidores inativos
    for servidor in down_servers:
        # Cria uma nova linha para o arquivo CSV com as informações do servidor inativo
        line = [servidor.server_name, servidor.server_url, servidor.data_falha, str(servidor.tempo_execucao), str(servidor.status)]
        # Escreve a linha no arquivo CSV
        csv_writer.writerow(line)
    # Descarrega quaisquer dados em buffer para o arquivo CSV
    downtime_file.flush()

def check_server(servidores):
    down_servers = []
    for servidor in servidores:
        # Obtém a hora atual
        agora = time.time()
        # Faz uma requisição GET para o servidor
        conn = http.client.HTTPConnection(servidor.server_url)
        conn.request("GET", "/")
        res = conn.getresponse()
        # Se houver um erro na requisição, marca o servidor como inativo e adiciona à lista de servidores inativos
        if res.status != 200:
            print("Server %s is down[%s]\n" % (servidor.server_name, res.reason))
            servidor.status = 0
            servidor.data_falha = time.strftime("%d/%m/%Y %H:%M:%S", time.localtime(agora))
            down_servers.append(servidor)
            continue
        # Se o status da resposta não for 200, marca o servidor como inativo e adiciona à lista de servidores inativos
        servidor.status = res.status
        if servidor.status != 200:
            servidor.data_falha = time.strftime("%d/%m/%Y %H:%M:%S", time.localtime(agora))
            down_servers.append(servidor)
        # Calcula o tempo de execução da requisição e exibe o status, tempo de carga e URL do servidor
        servidor.tempo_execucao = time.time() - agora
        print("Status: [%d] Tempo de carga: [%f] URL: [%s]\n" % (servidor.status, servidor.tempo_execucao, servidor.server_url))
    # Retorna a lista de servidores inativos
    return down_servers

def main():
    # Abre os arquivos de lista de servidores e de tempo de inatividade
    server_list, downtime_list = open_files(sys.argv[1], sys.argv[2])
    # Fecha os arquivos quando a função main() terminar
    server_list.close()
    downtime_list.close()
    # Cria uma lista de servidores a partir do arquivo de lista de servidores
    servidores = []
    for line in server_list:
        server_name, server_url = line.strip().split(",")
        servidores.append(Server(server_name, server_url))
    # Verifica quais servidores estão inativos
    down_servers = check_server(servidores)
    # Gera um arquivo CSV com a lista de servidores inativos
    generate_downtime(downtime_list, down_servers)

if __name__ == "__main__":
    main()