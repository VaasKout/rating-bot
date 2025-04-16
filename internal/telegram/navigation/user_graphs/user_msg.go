package user_graphs

import (
	"rating-bot/internal/data/telegram_redis"
	"rating-bot/internal/telegram/domain"
)

var loadingMessage = telegram_redis.InitOutputMessage(
	domain.USER_GET_RATING_LOADING_MESSAGE,
	&[][]string{},
)
