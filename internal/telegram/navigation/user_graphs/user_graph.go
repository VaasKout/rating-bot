package user_graphs

import (
	"rating-bot/internal/telegram/domain"
	"rating-bot/internal/telegram/navigation/graph_builder"
	"rating-bot/internal/telegram/navigation/root_graph"
	"rating-bot/pkg/collections"
)

var UserGraph = []map[string]collections.Node[graph_builder.GraphParams]{
	root_graph.RootGraph,
	userGetRatingGraph,
}

var userGetRatingGraph = map[string]collections.Node[graph_builder.GraphParams]{
	domain.START_STATE: userGetRatingStartNode,
}

// -------GET RATING------
var userGetRatingStartNode = collections.Node[graph_builder.GraphParams]{
	Current: domain.START_STATE,
	Next:    domain.START_STATE,
	Action:  userGetRatingStartAction,
}
