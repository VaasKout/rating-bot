package core_telegram

import (
	"rating-bot/pkg/core/network"
)

type TelegramNetworkApi interface {
	GetUpdate(offset int64) (*TelegramResponse, error)
	SendMessage(message *Message) *network.HttpResponse
	ProcessCallback(operation CallbackOperations) *network.HttpResponse
}
