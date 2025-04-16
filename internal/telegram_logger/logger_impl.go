package telegram_logger

import (
	"fmt"
	"html"
	"rating-bot/configs"
	"rating-bot/pkg/core/redis"
	"rating-bot/pkg/core_telegram"
	"rating-bot/pkg/logger"
	"time"
)

const (
	RATING_BOT_LOG = "rating_bot_log"
)

type LoggerImpl struct {
	botLogChatId int64
	coreTelegram core_telegram.TelegramNetworkApi
	redis        redis.RedisApi
}

func New(
	cfg *configs.Config,
	logger *logger.Logger,
	coreRedis redis.RedisApi,
) LoggerApi {
	var coreTelegram = core_telegram.New(logger, cfg)
	return &LoggerImpl{
		botLogChatId: cfg.BotProps.BotLogChatId,
		coreTelegram: coreTelegram,
		redis:        coreRedis,
	}
}

func (tgLogger *LoggerImpl) SendLogs() {
	tgLogger.sendLogs(RATING_BOT_LOG, tgLogger.botLogChatId)
}

func (tgLogger *LoggerImpl) sendLogs(key string, chatId int64) {
	for {
		time.Sleep(time.Second)
		var logMsg = tgLogger.redis.LPop(key)
		if logMsg == "" {
			continue
		}
		message := core_telegram.NewMessageBuilder().ChatId(chatId).Text(logMsg, false).Build()
		result := tgLogger.coreTelegram.SendMessage(message)
		if result.StatusCode == 429 {
			err := tgLogger.redis.LPush(key, logMsg)
			if err != nil {
				fmt.Println(err)
			}
			time.Sleep(time.Second * 10)
		}
	}
}

func (tgLogger *LoggerImpl) EnqueueBotLogMessage(message string) {
	if message == "" {
		return
	}
	botMsg := GetBotAnswerMessage(message)
	err := tgLogger.redis.RPush(RATING_BOT_LOG, botMsg)
	if err != nil {
		fmt.Println(err)
	}
}

func (tgLogger *LoggerImpl) EnqueueUserBotLogMessage(message string, userId int64, userTag string) {
	if message == "" || userId == 0 || userTag == "" {
		return
	}
	escapedMessage := html.EscapeString(message)
	userMessage := GetUserMessage(escapedMessage, userId, userTag)
	err := tgLogger.redis.RPush(RATING_BOT_LOG, userMessage)
	if err != nil {
		fmt.Println(err)
	}
}
