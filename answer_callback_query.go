package telegram

import json "github.com/pquerna/ffjson/ffjson"

// AnswerCallbackQueryParameters represents data for AnswerCallbackQuery method.
type AnswerCallbackQueryParameters struct {
	// Unique identifier for the query to be answered
	CallbackQueryID string `json:"callback_query_id"`

	// Text of the notification. If not specified, nothing will be shown to the
	// user, 0-200 characters
	Text string `json:"text,omitempty"`

	// URL that will be opened by the user's client. If you have created a Game
	// and accepted the conditions via @Botfather, specify the URL that opens
	// your game – note that this will only work if the query comes from a
	// callback_game button.
	//
	// Otherwise, you may use links like t.me/your_bot?start=XXXX that open your
	// bot with a parameter.
	URL string `json:"url,omitempty"`

	// If true, an alert will be shown by the client instead of a notification at
	// the top of the chat screen. Defaults to false.
	ShowAlert bool `json:"show_alert,omitempty"`

	// The maximum amount of time in seconds that the result of the callback
	// query may be cached client-side. Telegram apps will support caching
	// starting in version 3.14. Defaults to 0.
	CacheTime int `json:"cache_time,omitempty"`
}

// NewAnswerCallbackQuery creates AnswerCallbackQueryParameters only with
// required parameters.
func NewAnswerCallbackQuery(callbackQueryID string) *AnswerCallbackQueryParameters {
	return &AnswerCallbackQueryParameters{CallbackQueryID: callbackQueryID}
}

// AnswerCallbackQuery send answers to callback queries sent from inline
// keyboards. The answer will be displayed to the user as a notification at the
// top of the chat screen or as an alert. On success, True is returned.
//
// Alternatively, the user can be redirected to the specified Game URL. For this
// option to work, you must first create a game for your bot via @Botfather and
// accept the terms. Otherwise, you may use links like t.me/your_bot?start=XXXX
// that open your bot with a parameter.
func (bot *Bot) AnswerCallbackQuery(params *AnswerCallbackQueryParameters) (ok bool, err error) {
	dst, err := json.Marshal(params)
	if err != nil {
		return
	}

	resp, err := bot.request(dst, MethodAnswerCallbackQuery)
	if err != nil {
		return
	}

	err = json.Unmarshal(*resp.Result, &ok)
	return
}
