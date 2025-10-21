package direct_log_to_telegram

import (
	"github.com/sikalabs/slu/utils/telegram_utils"
	"github.com/sikalabsx/direct-log-to-telegram/internal/credentials"
	"github.com/sikalabsx/direct-log-to-telegram/internal/handle_error"
)

func Log(message string) error {
	botToken, chatId, err := credentials.GetCredentials()
	handle_error.HandleErrorFatalln(err)
	return telegram_utils.TelegramSendMessage(botToken, chatId, message)
}
