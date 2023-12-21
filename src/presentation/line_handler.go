package presentation

import (
	"github.com/labstack/echo/v4"
	"github.com/line/line-bot-sdk-go/v8/linebot"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
	"log/slog"
	"merubo.lambda/src/presentation/client_model/request_model"
	"os"
)

type LineHandler struct {
	logger  *slog.Logger
	lineBot *messaging_api.MessagingApiAPI
}

func NewLineHandler(logger *slog.Logger, lineBot *messaging_api.MessagingApiAPI) (handler *LineHandler) {
	lineHandler := new(LineHandler)
	lineHandler.logger = logger
	lineHandler.lineBot = lineBot
	return lineHandler
}

func (LineHandler *LineHandler) SendContactMessage(e echo.Context) (err error) {
	LineHandler.logger.Info("SendContactMessage 開始")
	contactBody := &request_model.ContactRequestModel{}
	if err := e.Bind(contactBody); err != nil {
		LineHandler.logger.Error(err.Error(), "message", "ContactBodyのBindに失敗しました。")
		return err
	}

	replyMessage := createResponseMessage(contactBody)

	message, err := LineHandler.lineBot.PushMessage(
		&messaging_api.PushMessageRequest{
			// 自分のlineIdを入れる。
			To: os.Getenv("OWNER_LINE_ID"),
			Messages: []messaging_api.MessageInterface{
				messaging_api.TextMessage{
					Text: replyMessage,
				},
			},
		}, "",
	)
	if err != nil {
		LineHandler.logger.Error(err.Error(), "message", "メッセージの送信に失敗しました。")
		return
	}
	LineHandler.logger.Info("SendContactMessage 終了", "message", message)
	return
}

func createResponseMessage(contactBody *request_model.ContactRequestModel) (replyMessage string) {
	replyMessage = contactBody.Name + "様 より問い合わせがきました。"
	replyMessage += "\n\n --------------\n\n"
	replyMessage += contactBody.Content
	replyMessage += "\n\n -------------- \n\n"
	replyMessage += "メアド:" + contactBody.Email
	return replyMessage
}

// EventHandler いつかのため用
func (LineHandler *LineHandler) EventHandler(e echo.Context) (err error) {
	channelSecret := os.Getenv("CHANNEL_SECRET")

	cb, err := webhook.ParseRequest(channelSecret, e.Request())
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			e.Response().WriteHeader(400)
		} else {
			e.Response().WriteHeader(500)
		}
		return
	}

	LineHandler.logger.Info("Line Handler 開始")

	for _, event := range cb.Events {
		switch e := event.(type) {
		// メッセージイベントの場合
		case webhook.MessageEvent:
			switch message := e.Message.(type) {
			// メッセージが送信されてきた場合
			case webhook.TextMessageContent:
				LineHandler.logger.Info("メッセージがきた")
				if _, err = LineHandler.lineBot.PushMessage(
					&messaging_api.PushMessageRequest{
						// 自分のlineIdを入れる。
						To: "U9da059653faa1b08c307f051641517f3",
						Messages: []messaging_api.MessageInterface{
							messaging_api.TextMessage{
								Text: message.Text,
							},
						},
					}, "",
				); err != nil {
					LineHandler.logger.Error("エラーが発生しました", "errorMsg", err.Error(), "big area", "MessageEvent", "small area", "TextMessageContent")
				} else {
					LineHandler.logger.Info("メッセージを送信しました。")
				}
			}
		default:
			LineHandler.logger.Info("default")
		}
	}
	return
}
