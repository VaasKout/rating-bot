package tg_redis_use_case

import (
	"fmt"
	"rating-bot/internal/data/telegram_redis"
	"rating-bot/internal/telegram/domain"
	"rating-bot/pkg/core_telegram"
	"time"
)

const PRIVATE_TYPE = "private"

type TelegramUseCaseRedisImpl struct {
	tgRedis            telegram_redis.TelegramRedisApi
	telegramNetworkApi core_telegram.TelegramNetworkApi
}

func NewRedisUseCase(
	tgRedis telegram_redis.TelegramRedisApi,
	tgApi core_telegram.TelegramNetworkApi,
) TelegramRedisUseCaseApi {
	return &TelegramUseCaseRedisImpl{
		tgRedis:            tgRedis,
		telegramNetworkApi: tgApi,
	}
}

func (repository *TelegramUseCaseRedisImpl) GetUser(inputMessage *telegram_redis.Message) *telegram_redis.UserData {
	fmt.Println(inputMessage)
	if inputMessage == nil || inputMessage.ChatId == 0 {
		return nil
	}
	role := repository.getRole(inputMessage.UserName)
	if len(role) > 0 {
		var user = repository.tgRedis.GetUserData(inputMessage.ChatId)
		if user != nil && len(user.UserName) > 0 && user.Role == role {
			return user
		}
		var newUser = &telegram_redis.UserData{
			ChatId:       inputMessage.ChatId,
			Role:         role,
			UserName:     inputMessage.UserName,
			CurrentState: domain.START_STATE,
		}
		repository.SaveUserData(newUser)
		return newUser
	}
	repository.tgRedis.ClearUserData(inputMessage.ChatId)
	return nil
}

func (repository *TelegramUseCaseRedisImpl) getRole(userName string) string {
	switch {
	case repository.tgRedis.IsUser(userName):
		return domain.USER
	default:
		return ""
	}
}

func (repository *TelegramUseCaseRedisImpl) SaveUserData(userData *telegram_redis.UserData) {
	repository.tgRedis.SaveUserData(userData)
}

func (repository *TelegramUseCaseRedisImpl) GetNewMessages() *[]telegram_redis.Message {

	messageList := make([]telegram_redis.Message, 0)
	offset := repository.tgRedis.GetOffset()
	telegramResponse, err := repository.telegramNetworkApi.GetUpdate(offset)

	if err != nil {
		fmt.Println("GET UPDATES ERROR")
		fmt.Println(err)
		return &messageList
	}

	for _, item := range telegramResponse.Result {
		if item.Message != nil && item.Message.Chat.Type == PRIVATE_TYPE {
			messageList = append(messageList, telegram_redis.Message{
				ChatId:   item.Message.Chat.Id,
				UserName: item.Message.From.UserName,
				Text:     item.Message.Text,
			})
		}
		repository.tgRedis.SaveOffset(item.UpdateId + 1)
	}

	return &messageList
}

func (repository *TelegramUseCaseRedisImpl) EnqueueOutputMessage(text string, markupButtons *[][]string, chatId int64) {
	if len(text) > 0 && chatId != 0 {
		msg := &telegram_redis.Message{
			ChatId:        chatId,
			MarkupButtons: *markupButtons,
			Text:          text,
		}
		repository.tgRedis.RPushMessage(msg)
	}
}

func (repository *TelegramUseCaseRedisImpl) PopMessage() *telegram_redis.Message {
	return repository.tgRedis.PopMessage()
}

func (repository *TelegramUseCaseRedisImpl) SendTelegramMessage(message *telegram_redis.Message) {
	if message == nil || len(message.Text) == 0 {
		return
	}
	var telegramMessage = core_telegram.NewMessageBuilder().
		ChatId(message.ChatId).
		ReplyKeyboard(&message.MarkupButtons).
		Text(message.Text, false).
		Build()

	result := repository.telegramNetworkApi.SendMessage(telegramMessage)
	switch result.StatusCode {
	case 429:
		repository.tgRedis.LPushMessage(message)
		time.Sleep(time.Second * 3)
	default:
		return
	}
}
