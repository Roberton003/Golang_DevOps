package telegram

import ( // Importação das bibliotecas
	"os"
	"strconv"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type Message struct {	// Struct que representa a mensagem
	Text string `json:"text"`
	GroupID int64 `json:"group_id"`
}

func SendTelegram(botApi string, message Message)
	bot, err := tgbotapi.NewBotAPI(botApi)
	if err != nil {
		panic(err)
	}
	message := Message{}
	message.Text = mensagem
	groupId := os.Getenv("TELEGRAM_GROUP_ID")
	if groupId == "" {
		panic("TELEGRAM_GROUP_ID não configurado")
	}
	message.GroupID, err = strconv.ParseInt(groupId, 10, 64)
	if err != nil {
		panic(err)
	}
	bot.Debug = true
	alertText := tgbotapi.NewMessage(message.GroupID, message.Text)
	bot.Send(alertText)
} // Fim da função SendTelegram


