package telegram

import json "github.com/pquerna/ffjson/ffjson"

// GetChatAdministratorsParameters represents data for GetChatAdministrators
// method.
type GetChatAdministratorsParameters struct {
	// Unique identifier for the target chat
	ChatID int64 `json:"chat_id"`
}

// GetChatAdministrators get a list of administrators in a chat. On success,
// returns an Array of ChatMember objects that contains information about all
// chat administrators except other bots. If the chat is a group or a supergroup
// and no administrators were appointed, only the creator will be returned.
func (bot *Bot) GetChatAdministrators(chatID int64) (members []ChatMember, err error) {
	dst, err := json.Marshal(&GetChatAdministratorsParameters{ChatID: chatID})
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodGetChatAdministrators)
	if err != nil {
		return
	}

	err = json.Unmarshal(*resp.Result, &members)
	return
}
