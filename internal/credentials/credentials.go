package credentials

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sikalabsx/direct-log-to-telegram/internal/decrypt"

	"github.com/sikalabsx/direct-log-to-telegram/internal/handle_error"
)

const BOT_TOKEN_ENCRYPTED = "uhZ0TE7E3bJyE5sm8sJFPgl92jjC42hER5UqAZ22lRXb/Dz1y8yvV+4sFNtCKAhzGx4qrTbMPHk4eu+15dBhzHh59syQxnBEwrU="
const CHAT_ID_ENCRYPTED = "i22r21e0gwFG1FJc66uD1EWOdiwsjZoSfdVe8Bj3oPdWjW6tJM8d"

func GetCredentials() (string, int64, error) {
	password := getPassword()

	botToken, err := decrypt.Decrypt(BOT_TOKEN_ENCRYPTED, password)
	if err != nil {
		handle_error.HandleErrorFatalln(fmt.Errorf("failed to decrypt BOT_TOKEN, check password in environment variable DIRECT_LOG_TO_TELEGRAM_PASSWORD"))
	}

	chatIdStr, err := decrypt.Decrypt(CHAT_ID_ENCRYPTED, password)
	if err != nil {
		handle_error.HandleErrorFatalln(fmt.Errorf("failed to decrypt CHAT_ID, check password in environment variable DIRECT_LOG_TO_TELEGRAM_PASSWORD"))
	}

	chatId, err := strconv.Atoi(chatIdStr)
	if err != nil {
		handle_error.HandleErrorFatalln(fmt.Errorf("failed to convert CHAT_ID to int"))
	}
	return botToken, int64(chatId), nil
}

func getPassword() string {
	password := os.Getenv("DIRECT_LOG_TO_TELEGRAM_PASSWORD")
	if password == "" {
		handle_error.HandleErrorFatalln(fmt.Errorf("DIRECT_LOG_TO_TELEGRAM_PASSWORD environment variable is not set"))
	}

	// If password starts with /, treat it as an absolute file path
	if strings.HasPrefix(password, "/") {
		password = getPasswordFromFile(password)
	}

	return password
}

func getPasswordFromFile(filePath string) string {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		handle_error.HandleErrorFatalln(fmt.Errorf("failed to read password from file %s: %w", filePath, err))
	}

	// Remove all whitespaces and newlines
	password := strings.TrimSpace(string(fileContent))
	password = strings.ReplaceAll(password, "\n", "")
	password = strings.ReplaceAll(password, "\r", "")
	password = strings.ReplaceAll(password, " ", "")
	password = strings.ReplaceAll(password, "\t", "")
	return password
}
