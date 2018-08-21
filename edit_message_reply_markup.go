package telegram

import json "github.com/pquerna/ffjson/ffjson"

// EditMessageReplyMarkupParameters represents data for EditMessageReplyMarkup method.
type EditMessageReplyMarkupParameters struct {
	// Required if inline_message_id is not specified. Unique identifier for the
	// target chat or username of the target channel (in the format
	// @channelusername)
	ChatID int64 `json:"chat_id,omitempty"`

	// Required if inline_message_id is not specified. Identifier of the sent
	// message
	MessageID int `json:"message_id,omitempty"`

	// Required if chat_id and message_id are not specified. Identifier of the
	// inline message
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// A JSON-serialized object for an inline keyboard.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// EditMessageReplyMarkup edit only the reply markup of messages sent by the bot
// or via the bot (for inline bots). On success, if edited message is sent by the
// bot, the edited Message is returned, otherwise True is returned.
func (bot *Bot) EditMessageReplyMarkup(params *EditMessageReplyMarkupParameters) (msg *Message, err error) {
	dst, err := json.Marshal(params)
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodEditMessageReplyMarkup)
	if err != nil {
		return
	}

	msg = new(Message)
	err = json.Unmarshal(*resp.Result, msg)
	return
}
