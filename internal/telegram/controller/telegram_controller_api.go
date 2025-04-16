package controller

type TelegramControllerApi interface {
	ListenUpdates()
	SendEnqueuedMessages()
}
