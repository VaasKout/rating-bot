package use_case

import (
	"rating-bot/internal/data/telegram_redis"
	"rating-bot/internal/telegram/use_case/parser_use_case"
	"rating-bot/internal/telegram/use_case/tg_redis_use_case"
)

type TelegramUseCase interface {
	RedisUseCase() tg_redis_use_case.TelegramRedisUseCaseApi
	ExecRequestForRating(msg *telegram_redis.Message)
	LayoutParserUseCase() parser_use_case.TelegramParserUseCaseApi
}
