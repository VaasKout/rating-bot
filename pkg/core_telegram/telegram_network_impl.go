package core_telegram

import (
	"encoding/json"
	"fmt"
	"log"
	"rating-bot/configs"
	"rating-bot/pkg/core/network"
	"rating-bot/pkg/logger"
)

type TelegramNetworkImpl struct {
	networkApi network.NetworkApi
	botKey     string
	logger     *logger.Logger
}

func New(logger *logger.Logger, cfg *configs.Config) *TelegramNetworkImpl {
	var networkApi = network.New(logger)
	return &TelegramNetworkImpl{
		networkApi: networkApi,
		botKey:     cfg.BotProps.Token,
	}
}

func (telegramNetwork *TelegramNetworkImpl) ProcessCallback(operation CallbackOperations) *network.HttpResponse {
	switch operation.(type) {
	case *EditMessage:
		return telegramNetwork.editMessage(operation)
	case *DeleteKeyboard:
		return telegramNetwork.deleteKeyboard(operation)
	}

	return nil
}

func (telegramNetwork *TelegramNetworkImpl) editMessage(editMessageRequest CallbackOperations) *network.HttpResponse {
	var request = telegramNetwork.getHttpRequest(editMessageRequest, EditMessageTextMethod)
	return telegramNetwork.networkApi.MakePostRequest(request)
}

func (telegramNetwork *TelegramNetworkImpl) deleteKeyboard(keyboard CallbackOperations) *network.HttpResponse {
	var request = telegramNetwork.getHttpRequest(
		keyboard,
		DeleteKeyboardMethod,
	)
	return telegramNetwork.networkApi.MakePostRequest(request)
}

func (telegramNetwork *TelegramNetworkImpl) SendMessage(message *Message) *network.HttpResponse {
	switch {
	case message.PhotoId != "":
		return telegramNetwork.sendPhoto(message)
	case message.VideoId != "":
		return telegramNetwork.sendVideo(message)
	case message.VoiceId != "":
		return telegramNetwork.sendVoice(message)
	case message.StickerId != "":
		return telegramNetwork.sendSticker(message)
	case message.AudioId != "":
		return telegramNetwork.sendAudio(message)
	case message.DocumentId != "":
		return telegramNetwork.sendDocument(message)
	default:
		return telegramNetwork.processRequest(telegramNetwork.getHttpRequest(message, SendMessageMethod))
	}
}

func (telegramNetwork *TelegramNetworkImpl) sendPhoto(message *Message) *network.HttpResponse {
	return telegramNetwork.processRequest(telegramNetwork.getHttpRequest(message, SendPhotoMethod))
}

func (telegramNetwork *TelegramNetworkImpl) sendDocument(message *Message) *network.HttpResponse {
	return telegramNetwork.processRequest(telegramNetwork.getHttpRequest(message, SendDocumentMethod))
}

func (telegramNetwork *TelegramNetworkImpl) sendAudio(message *Message) *network.HttpResponse {
	return telegramNetwork.processRequest(telegramNetwork.getHttpRequest(message, SendAudioMethod))
}

func (telegramNetwork *TelegramNetworkImpl) sendVideo(message *Message) *network.HttpResponse {
	return telegramNetwork.processRequest(telegramNetwork.getHttpRequest(message, SendVideoMethod))
}

func (telegramNetwork *TelegramNetworkImpl) sendSticker(message *Message) *network.HttpResponse {
	return telegramNetwork.processRequest(telegramNetwork.getHttpRequest(message, SendStickerMethod))
}

func (telegramNetwork *TelegramNetworkImpl) sendVoice(message *Message) *network.HttpResponse {
	return telegramNetwork.processRequest(telegramNetwork.getHttpRequest(message, SendVoiceMethod))
}

func (telegramNetwork *TelegramNetworkImpl) GetUpdate(offset int64) (*TelegramResponse, error) {
	var url = fmt.Sprintf(
		"https://api.telegram.org/bot%s/getUpdates?offset=%d",
		telegramNetwork.botKey,
		offset,
	)
	result := telegramNetwork.networkApi.MakeGetRequest(url)
	if result.StatusCode != 200 {
		log.Println(result.Error.Error())
		return &TelegramResponse{}, result.Error
	}
	var telegramResponse TelegramResponse
	err := json.Unmarshal(result.Body, &telegramResponse)
	if err != nil {
		log.Println(err.Error())
	}
	return &telegramResponse, err

}

func (telegramNetwork *TelegramNetworkImpl) getHttpRequest(body interface{}, method string) *network.HttpRequest {

	var url = fmt.Sprintf(
		"https://api.telegram.org/bot%s/%s",
		telegramNetwork.botKey,
		method,
	)

	jsonMessage, err := json.Marshal(&body)
	if err != nil {
		fmt.Println(err)
		telegramNetwork.logger.Log.Error(fmt.Sprintf("Cannot marshal body: %v", body))
	}

	return &network.HttpRequest{
		Url:  url,
		Body: jsonMessage,
	}
}

func (telegramNetwork *TelegramNetworkImpl) processRequest(request *network.HttpRequest) *network.HttpResponse {
	var response = telegramNetwork.networkApi.MakePostRequest(request)
	telegramNetwork.validateResponse(response)
	return response
}

func (telegramNetwork *TelegramNetworkImpl) validateResponse(response *network.HttpResponse) {
	var errorDescription ErrorMessage

	if response.StatusCode == 200 {
		return
	}

	err := json.Unmarshal(response.Body, &errorDescription)
	if err != nil {
		log.Println(err)
		return
	}

	switch errorDescription.Description {
	case VoiceMessagesForbiddenError:
		response.Body = []byte(VoiceMessagesForbiddenMessage)
	}
}
