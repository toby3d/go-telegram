package telegram

import json "github.com/pquerna/ffjson/ffjson"

// SendLocationParameters represents data for SendLocation method.
type SendLocationParameters struct {
	// Unique identifier for the target private chat
	ChatID int64 `json:"chat_id"`

	// Latitude of the location
	Latitude float32 `json:"latitude"`

	// Longitude of the location
	Longitude float32 `json:"longitude"`

	// Period in seconds for which the location will be updated (see Live
	// Locations), should be between 60 and 86400.
	LivePeriod int `json:"live_period,omitempty"`

	// If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`

	// Sends the message silently. Users will receive a notification with no
	// sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// A JSON-serialized object for an inline keyboard. If empty, one 'Pay total
	// price' button will be shown. If not empty, the first button must be a Pay
	// button.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// NewLocation creates SendLocationParameters only with required parameters.
func NewLocation(chatID int64, latitude, longitude float32) *SendLocationParameters {
	return &SendLocationParameters{
		ChatID:    chatID,
		Latitude:  latitude,
		Longitude: longitude,
	}
}

// SendLocation send point on the map. On success, the sent Message is returned.
func (bot *Bot) SendLocation(params *SendLocationParameters) (msg *Message, err error) {
	dst, err := json.Marshal(params)
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodSendLocation)
	if err != nil {
		return
	}

	msg = new(Message)
	err = json.Unmarshal(*resp.Result, msg)
	return
}
