package root_graph

import (
	"rating-bot/internal/data/telegram_redis"
	"rating-bot/internal/telegram/domain"
	"rating-bot/internal/telegram/navigation/graph_builder"
	"rating-bot/pkg/collections"
)

var RootGraph = map[string]collections.Node[graph_builder.GraphParams]{
	startNode.Current:  startNode,
	cancelNode.Current: cancelNode,
}

var startNode = collections.Node[graph_builder.GraphParams]{
	Current: domain.GENERIC_START_BUTTON,
	Next:    domain.START_STATE,
	Action: func(params *graph_builder.GraphParams) string {
		params.LinearMessage = telegram_redis.InitOutputMessage(
			domain.START_MESSAGE,
			&[][]string{},
		)
		return collections.NEXT
	},
}

var cancelNode = collections.Node[graph_builder.GraphParams]{
	Current: domain.GENERIC_CANCEL_BUTTON,
	Next:    domain.START_STATE,
	Action: func(params *graph_builder.GraphParams) string {
		params.LinearMessage = telegram_redis.InitOutputMessage(
			domain.CANCEL_MESSAGE,
			&[][]string{},
		)
		return collections.NEXT
	},
}
