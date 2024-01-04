module github.com/utkinn/telegram-news-in-group-remover

go 1.21.1

require (
	github.com/dlclark/regexp2 v1.10.0
	github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.6.0
	github.com/joho/godotenv v1.5.1
)

replace github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.6.0 => github.com/OvyFlash/telegram-bot-api/v5 v5.0.0-20231230151827-6d16deaa376e
