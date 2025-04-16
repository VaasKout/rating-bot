package navigation

import (
	"rating-bot/internal/data/telegram_redis"
)

type HandlerApi interface {
	HandleMessage(message *telegram_redis.Message)
}
