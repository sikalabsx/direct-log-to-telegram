# direct-log-to-telegram

A simple Go library for sending log messages directly to Telegram. Perfect for monitoring applications, scripts, and services with instant notifications.

## Features

- Simple one-function API: `Log(message string) error`
- Encrypted credential storage (bot token and chat ID are encrypted in code)
- Password from environment variable or file
- Zero external configuration files needed
- Built on top of proven telegram utilities

## Installation

```bash
go get github.com/sikalabsx/direct-log-to-telegram
```

## Usage

### Basic Example

```go
package main

import (
    "log"
    "github.com/sikalabsx/direct-log-to-telegram/pkg/direct_log_to_telegram"
)

func main() {
    err := direct_log_to_telegram.Log("Hello from my application!")
    if err != nil {
        log.Fatalf("Failed to send message: %v", err)
    }
}
```

### Environment Configuration

The library supports multiple ways to provide the decryption password:

**Option 1: Direct password**
```bash
export DIRECT_LOG_TO_TELEGRAM_PASSWORD="your-password-here"
```

**Option 2: Password from file**
```bash
export DIRECT_LOG_TO_TELEGRAM_PASSWORD="/path/to/password/file"
```

**Option 3: Default file location (no environment variable needed)**

If `DIRECT_LOG_TO_TELEGRAM_PASSWORD` is not set, the library will automatically try to read the password from:
```
/etc/direct-log-to-telegram/DIRECT_LOG_TO_TELEGRAM_PASSWORD
```

When using a file path (option 2 or 3), the library will read the password from the file, automatically trimming whitespace and newlines.

### Running the Example

```bash
cd examples/simple
export DIRECT_LOG_TO_TELEGRAM_PASSWORD="your-password-here"
go run main.go
```

## How It Works

1. The library stores encrypted Telegram bot token and chat ID in the code
2. At runtime, it reads the decryption password from `DIRECT_LOG_TO_TELEGRAM_PASSWORD`
3. Credentials are decrypted and used to send messages via Telegram Bot API
4. Messages are sent to the pre-configured chat

## Use Cases

- Application health monitoring
- Error notifications
- Deployment notifications
- Script completion alerts
- System monitoring
- CI/CD pipeline notifications

## Security Notes

- Bot token and chat ID are encrypted and stored in the code
- Password must be provided via environment variable
- Password can be read from a file for better security in production
- No plain-text credentials in configuration files

## License

MIT

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
