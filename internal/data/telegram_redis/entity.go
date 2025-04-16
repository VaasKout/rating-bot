package telegram_redis

import (
	"encoding/json"
	"fmt"
)

type UserData struct {
	UserName     string `json:"user_name"`
	ChatId       int64  `json:"chat_id"`
	Role         string `json:"role"`
	CurrentState string `json:"current_state"`
}

type Message struct {
	ChatId        int64      `json:"chat_id"`
	UserName      string     `json:"username"`
	MarkupButtons [][]string `json:"markup_buttons"`
	Text          string     `json:"text"`
}

func (userData *UserData) UpdateUserState(newState string) {
	if userData != nil {
		userData.CurrentState = newState
	}
}

func InitOutputMessage(text string, buttons *[][]string) *Message {
	var markupButtons = make([][]string, 0)
	if buttons != nil {
		markupButtons = *buttons
	}
	return &Message{
		Text:          text,
		MarkupButtons: markupButtons,
	}
}

func MapUserDataToJson(userData *UserData) string {
	result, err := json.Marshal(userData)
	if err != nil {
		fmt.Println("MapUserDataToJson")
		fmt.Println(err)
		return ""
	}
	return string(result)
}

func MapJsonToUserData(body string) *UserData {
	var userData UserData
	if body == "" {
		return &userData
	}
	err := json.Unmarshal([]byte(body), &userData)
	if err != nil {
		fmt.Println("MapJsonToUserData")
		fmt.Println(err)
		return nil
	}
	return &userData
}

func MapMessageToJson(message *Message) string {
	result, err := json.Marshal(message)
	if err != nil {
		fmt.Println("MapMessageToJson")
		fmt.Println(err)
		return ""
	}
	return string(result)
}

func MapJsonToMessage(body string) *Message {
	var message Message
	err := json.Unmarshal([]byte(body), &message)
	if err != nil {
		fmt.Println("MapJsonToMessage")
		fmt.Println(err)
		return nil
	}
	return &message
}
