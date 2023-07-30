package telegram

import (
	
	"log"
	"os"
	"net/http"
	"encoding/json"
)

type Message struct {
	Text    string `json:"text"`
	GroupId int64  `json:"group_id"`
}
type errorMessage struct {
	Error    string `json:"text"`
}

func SendTelegram(w http.ResponseWriter, r *http.Request) {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("Token não configurado")
	}
	var errorMessage errorMessage
	message := Message{}
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		errorMessage.Error = ("Erro ao decodificar JSON")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

}

