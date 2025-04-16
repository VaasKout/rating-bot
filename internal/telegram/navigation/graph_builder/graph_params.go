package graph_builder

import (
	"rating-bot/internal/data/telegram_redis"
	"rating-bot/internal/telegram/domain"
	"rating-bot/internal/telegram/use_case"
	"rating-bot/pkg/collections"
)

type GraphParams struct {
	UserData      *telegram_redis.UserData
	InputMessage  *telegram_redis.Message
	LinearMessage *telegram_redis.Message
	RatingMessage *telegram_redis.Message
	UseCase       use_case.TelegramUseCase
}

func (params *GraphParams) BackPressed() bool {
	return params != nil && params.InputMessage != nil && params.InputMessage.Text == domain.GENERIC_BACK_BUTTON
}

func (params *GraphParams) GoBack(msg *telegram_redis.Message) string {
	params.LinearMessage = msg
	return collections.PREVIOUS
}
