package navigation

import (
	"rating-bot/internal/data/telegram_redis"
	"rating-bot/internal/telegram/domain"
	"rating-bot/internal/telegram/navigation/graph_builder"
	"rating-bot/internal/telegram/navigation/user_graphs"
	"rating-bot/internal/telegram/use_case"
	"rating-bot/pkg/collections"
)

var rolesHandlerMap = map[string][]map[string]collections.Node[graph_builder.GraphParams]{
	domain.USER: user_graphs.UserGraph,
}

type HandlerImpl struct {
	useCase use_case.TelegramUseCase
}

func New(
	telegramUseCase use_case.TelegramUseCase,
) HandlerApi {
	return &HandlerImpl{
		useCase: telegramUseCase,
	}
}

func (handler *HandlerImpl) HandleMessage(message *telegram_redis.Message) {
	var userData = handler.useCase.RedisUseCase().GetUser(message)
	if userData == nil {
		return
	}

	graphBuilder := graph_builder.NewGraphParamsBuilder(
		userData,
		message,
		handler.useCase,
	)

	if result, ok := rolesHandlerMap[userData.Role]; ok {
		graphBuilder.HandleRoleMap(result).AdjustMarkupButtons()
	}

	var graphParams = graphBuilder.Build()
	handler.SendMessageIfOutputIsEmpty(graphParams)
	handler.SaveGraphParams(graphParams)
}

func (handler *HandlerImpl) SendMessageIfOutputIsEmpty(params *graph_builder.GraphParams) {
	if handler == nil || params.UserData == nil {
		return
	}
	if params.LinearMessage == nil {
		handler.useCase.RedisUseCase().EnqueueOutputMessage(
			domain.START_MESSAGE,
			&[][]string{},
			params.UserData.ChatId,
		)
	}
}

func (handler *HandlerImpl) SaveGraphParams(params *graph_builder.GraphParams) {
	handler.useCase.RedisUseCase().SaveUserData(params.UserData)
	if params.LinearMessage != nil {
		handler.useCase.RedisUseCase().EnqueueOutputMessage(
			params.LinearMessage.Text,
			&params.LinearMessage.MarkupButtons,
			params.UserData.ChatId,
		)
	}
	if params.RatingMessage != nil {
		handler.useCase.ExecRequestForRating(params.RatingMessage)
		params.RatingMessage = nil
	}
}
