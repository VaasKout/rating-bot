package graph_builder

import (
	"rating-bot/internal/data/telegram_redis"
	"rating-bot/internal/telegram/domain"
	"rating-bot/internal/telegram/use_case"
	"rating-bot/pkg/collections"
)

type GraphParamsBuilder interface {
	HandleRoleMap(roleMap []map[string]collections.Node[GraphParams]) GraphParamsBuilder
	AdjustMarkupButtons() GraphParamsBuilder
	Build() *GraphParams
}

func NewGraphParamsBuilder(
	userData *telegram_redis.UserData,
	inputMessage *telegram_redis.Message,
	useCase use_case.TelegramUseCase,
) GraphParamsBuilder {
	return &GraphParams{
		UserData:     userData,
		InputMessage: inputMessage,
		UseCase:      useCase,
	}
}

func (params *GraphParams) HandleRoleMap(roleMap []map[string]collections.Node[GraphParams]) GraphParamsBuilder {
	if params == nil {
		return params
	}
	for _, item := range roleMap {
		if result, ok := item[params.InputMessage.Text]; ok {
			params.handleActionResult(&result)
			return params
		}
		if result, ok := item[params.UserData.CurrentState]; ok {
			params.handleActionResult(&result)
			return params
		}
	}
	params.LinearMessage = nil
	return params
}

func (params *GraphParams) AdjustMarkupButtons() GraphParamsBuilder {
	if params.UserData.CurrentState == domain.START_STATE && params.LinearMessage != nil {
		params.LinearMessage.MarkupButtons = [][]string{}
	}
	return params
}

func (params *GraphParams) Build() *GraphParams {
	return params
}

func (params *GraphParams) handleActionResult(node *collections.Node[GraphParams]) {
	if params == nil || params.UserData == nil {
		return
	}

	direction := node.Action(params)
	updatedState := ""
	switch direction {
	case collections.NEXT:
		updatedState = node.Next
	case collections.PREVIOUS:
		updatedState = node.Previous
	case collections.CURRENT:
		updatedState = node.Current
	default:
		updatedState = direction
	}

	if updatedState == "" {
		updatedState = domain.START_STATE
	}

	params.UserData.UpdateUserState(updatedState)
}
