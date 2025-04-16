package tg_redis_use_case

import (
	"rating-bot/internal/data/telegram_redis"
)

type TelegramRedisUseCaseApi interface {
	GetUser(inputMessage *telegram_redis.Message) *telegram_redis.UserData
	SaveUserData(userData *telegram_redis.UserData)
	EnqueueOutputMessage(text string, markupButtons *[][]string, chatId int64)
	PopMessage() *telegram_redis.Message
	GetNewMessages() *[]telegram_redis.Message
	SendTelegramMessage(message *telegram_redis.Message)
}
