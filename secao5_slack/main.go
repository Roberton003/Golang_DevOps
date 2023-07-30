package main

import
 "alertamanager/slack"



func main() { 
	slack.SendSlack("Alerta de servidor down: Google\n Erro: Erro ao conectar com o servidor\n Horário: 24/07/2023 12:00:00")
}

