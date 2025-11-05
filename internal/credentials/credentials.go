package credentials

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sikalabs/sikalabs-crypt-go/pkg/sikalabs_crypt"
	"github.com/sikalabsx/direct-log-to-telegram/internal/error_utils"
)

const BOT_TOKEN_ENCRYPTED = "jp+90p1nRMMFmiTWbErBGBiaMOP77Q9Gs3zHCp4uRQBIOKxUvdIhVKVvk6XuEXnZb6TEcpvMQeU6mQPwR7aMWuPTKiUHIYQVarAdsYUkpaJhNUdgIT94+tK6"
const CHAT_ID_ENCRYPTED = "xfkGiSzXH+X+JtXxDtBoCzhgyG1kpXXwiCg+DaI9WFb+481RkHoYDsJzEhyWZ0l39OBf8XThMg=="

func GetCredentials() (string, int64, error) {
	password := getPassword()

	botToken, err := sikalabs_crypt.SikaLabsSymmetricDecryptV1(password, BOT_TOKEN_ENCRYPTED)
	if err != nil {
		error_utils.HandleErrorFatalln(fmt.Errorf("failed to decrypt BOT_TOKEN, check password in DIRECT_LOG_TO_TELEGRAM_PASSWORD environment variable or /etc/direct-log-to-telegram/DIRECT_LOG_TO_TELEGRAM_PASSWORD file"))
	}

	chatIdStr, err := sikalabs_crypt.SikaLabsSymmetricDecryptV1(password, CHAT_ID_ENCRYPTED)
	if err != nil {
		error_utils.HandleErrorFatalln(fmt.Errorf("failed to decrypt CHAT_ID, check password in DIRECT_LOG_TO_TELEGRAM_PASSWORD environment variable or /etc/direct-log-to-telegram/DIRECT_LOG_TO_TELEGRAM_PASSWORD file"))
	}

	chatId, err := strconv.Atoi(chatIdStr)
	if err != nil {
		error_utils.HandleErrorFatalln(fmt.Errorf("failed to convert CHAT_ID to int"))
	}
	return botToken, int64(chatId), nil
}

func getPassword() string {
	password := os.Getenv("DIRECT_LOG_TO_TELEGRAM_PASSWORD")
	if password == "" {
		// Try to read from default file path
		password = getPasswordFromFile("/etc/direct-log-to-telegram/DIRECT_LOG_TO_TELEGRAM_PASSWORD")
		return password
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
		error_utils.HandleErrorFatalln(fmt.Errorf("failed to read password from file %s: %w", filePath, err))
	}

	// Remove all whitespaces and newlines
	password := strings.TrimSpace(string(fileContent))
	password = strings.ReplaceAll(password, "\n", "")
	password = strings.ReplaceAll(password, "\r", "")
	password = strings.ReplaceAll(password, " ", "")
	password = strings.ReplaceAll(password, "\t", "")
	return password
}
