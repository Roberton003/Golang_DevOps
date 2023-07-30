package slack

import "os"

func SendSlack(textAlerta string) {
	token := os.Getenv("SLACK_AUTH_TOKEN")
	if token == "" {
		panic("SLACK_AUTH_TOKEN não encontrado")
	}
	channelID := os.Getenv("SLACK_CHANNEL_ID")
	if channelID == "" {
		panic("SLACK_CHANNEL_ID não encontrado")
	}

	client := slack.New(token, slack.Optiondebug(true))
	attachment := slack.Attachment{
		color: "danger",
		Pretext: "Alerta de servidor down",
		Text: textAlerta,
	}
	_, timestamp, err := client.PostMessage(channelID, slack.MsgOptionAttachments(attachment))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Mensagem enviada com sucesso %s as %s", channelID, timestamp)