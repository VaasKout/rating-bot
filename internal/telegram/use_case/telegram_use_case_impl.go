package use_case

import (
	"rating-bot/configs"
	"rating-bot/internal/data/telegram_redis"
	"rating-bot/internal/layout_parser"
	"rating-bot/internal/telegram/use_case/parser_use_case"
	"rating-bot/internal/telegram/use_case/tg_redis_use_case"
	"rating-bot/pkg/core/redis"
	"rating-bot/pkg/core_telegram"
	"rating-bot/pkg/logger"
)

type TelegramUseCaseImpl struct {
	tgRedis         telegram_redis.TelegramRedisApi
	telegramNetwork core_telegram.TelegramNetworkApi
	layoutParser    layout_parser.LayoutParserApi
	redisUseCase    tg_redis_use_case.TelegramRedisUseCaseApi
	parserUseCase   parser_use_case.TelegramParserUseCaseApi
	ratingMsgChan   chan telegram_redis.Message
}

func (useCase *TelegramUseCaseImpl) ExecRequestForRating(msg *telegram_redis.Message) {
	if msg == nil || msg.Text == "" {
		return
	}
	useCase.ratingMsgChan <- *msg
}

func (useCase *TelegramUseCaseImpl) RedisUseCase() tg_redis_use_case.TelegramRedisUseCaseApi {
	if useCase.redisUseCase == nil {
		useCase.redisUseCase = tg_redis_use_case.NewRedisUseCase(useCase.tgRedis, useCase.telegramNetwork)
	}
	return useCase.redisUseCase
}

func (useCase *TelegramUseCaseImpl) LayoutParserUseCase() parser_use_case.TelegramParserUseCaseApi {
	if useCase.parserUseCase == nil {
		useCase.parserUseCase = parser_use_case.NewLayoutParserUseCase(useCase.layoutParser)
	}
	return useCase.parserUseCase
}

func (useCase *TelegramUseCaseImpl) initRatingChanHandler() {
	for {
		value := <-useCase.ratingMsgChan
		app := useCase.LayoutParserUseCase().GetAppInfo(value.Text)
		msg := resultAppInfoMessage(app)
		useCase.RedisUseCase().EnqueueOutputMessage(
			msg.Text,
			&msg.MarkupButtons,
			value.ChatId,
		)
	}
}

func New(
	log *logger.Logger,
	cfg *configs.Config,
	coreRds redis.RedisApi,
	layoutParser layout_parser.LayoutParserApi,
	ratingMsgChan chan telegram_redis.Message,
) TelegramUseCase {
	var tgNetwork = core_telegram.New(log, cfg)
	var tgRedis = telegram_redis.New(coreRds)

	useCase := &TelegramUseCaseImpl{
		tgRedis:         tgRedis,
		telegramNetwork: tgNetwork,
		layoutParser:    layoutParser,
		ratingMsgChan:   ratingMsgChan,
	}

	go useCase.initRatingChanHandler()
	return useCase
}
