package user_graphs

import (
	"rating-bot/internal/telegram/navigation/graph_builder"
	"rating-bot/pkg/collections"
)

var userGetRatingStartAction = func(params *graph_builder.GraphParams) string {
	params.LinearMessage = loadingMessage
	params.RatingMessage = params.InputMessage
	return collections.NEXT
}
