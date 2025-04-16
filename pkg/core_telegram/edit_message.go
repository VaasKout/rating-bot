package core_telegram

type CallbackOperations interface{}

type EditMessage struct {
	ChatId    int64  `json:"chat_id"`
	MessageId int64  `json:"message_id"`
	Text      string `json:"text"`
	CallbackOperations
}

type DeleteKeyboard struct {
	ChatId    int64           `json:"chat_id"`
	MessageId int64           `json:"message_id"`
	Keyboard  [][]interface{} `json:"reply_markup"`
	CallbackOperations
}
