package telegram

type (
	// GetChatParameters represents data for GetChat method.
	GetChatParameters struct {
		// Unique identifier for the target chat
		ChatID int64 `json:"chat_id"`
	}

	// GetChatAdministratorsParameters represents data for GetChatAdministrators
	// method.
	GetChatAdministratorsParameters struct {
		// Unique identifier for the target chat
		ChatID int64 `json:"chat_id"`
	}

	// GetChatMemberParameters represents data for GetChatMember method.
	GetChatMemberParameters struct {
		// Unique identifier for the target chat
		ChatID int64 `json:"chat_id"`

		// Unique identifier of the target user
		UserID int `json:"user_id"`
	}

	// GetChatMembersCountParameters represents data for GetChatMembersCount method.
	GetChatMembersCountParameters struct {
		// Unique identifier for the target chat
		ChatID int64 `json:"chat_id"`
	}

	// GetFileParameters represents data for GetFile method.
	GetFileParameters struct {
		// File identifier to get info about
		FileID string `json:"file_id"`
	}

	// GetUpdatesParameters represents data for GetUpdates method.
	GetUpdatesParameters struct {
		// Identifier of the first update to be returned. Must be greater by one than the highest among the
		// identifiers of previously received updates. By default, updates starting with the earliest unconfirmed
		// update are returned. An update is considered confirmed as soon as getUpdates is called with an offset
		// higher than its update_id. The negative offset can be specified to retrieve updates starting from -offset
		// update from the end of the updates queue. All previous updates will forgotten.
		Offset int `json:"offset,omitempty"`

		// Limits the number of updates to be retrieved. Values between 1—100 are accepted. Defaults to 100.
		Limit int `json:"limit,omitempty"`

		// Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling. Should be positive, short
		// polling should be used for testing purposes only.
		Timeout int `json:"timeout,omitempty"`

		// List the types of updates you want your bot to receive. For example, specify ["message",
		// "edited_channel_post", "callback_query"] to only receive updates of these types. See Update for a complete
		// list of available update types. Specify an empty list to receive all updates regardless of type (default).
		// If not specified, the previous setting will be used.
		//
		// Please note that this parameter doesn't affect updates created before the call to the getUpdates, so
		// unwanted updates may be received for a short period of time.
		AllowedUpdates []string `json:"allowed_updates,omitempty"`
	}

	// GetUserProfilePhotosParameters represents data for GetUserProfilePhotos method.
	GetUserProfilePhotosParameters struct {
		// Unique identifier of the target user
		UserID int `json:"user_id"`

		// Sequential number of the first photo to be returned. By default, all
		// photos are returned.
		Offset int `json:"offset,omitempty"`

		// Limits the number of photos to be retrieved. Values between 1—100 are
		// accepted. Defaults to 100.
		Limit int `json:"limit,omitempty"`
	}

	// GetGameHighScoresParameters represents data for GetGameHighScores method.
	GetGameHighScoresParameters struct {
		// Target user id
		UserID int `json:"user_id"`

		// Required if inline_message_id is not specified. Identifier of the sent
		// message
		MessageID int `json:"message_id,omitempty"`

		// Required if inline_message_id is not specified. Unique identifier for the
		// target chat
		ChatID int64 `json:"chat_id,omitempty"`

		// Required if chat_id and message_id are not specified. Identifier of the
		// inline message
		InlineMessageID string `json:"inline_message_id,omitempty"`
	}

	// GetStickerSetParameters represents data for GetStickerSet method.
	GetStickerSetParameters struct {
		// Name of the sticker set
		Name string `json:"name"`
	}
)

// NewGameHighScores creates GetGameHighScoresParameters only with required parameters.
func NewGameHighScores(userID int) *GetGameHighScoresParameters {
	return &GetGameHighScoresParameters{
		UserID: userID,
	}
}

// GetChat get up to date information about the chat (current name of the user
// for one-on-one conversations, current username of a user, group or channel,
// etc.). Returns a Chat object on success.
func (bot *Bot) GetChat(chatID int64) (*Chat, error) {
	dst, err := parser.Marshal(&GetChatParameters{ChatID: chatID})
	if err != nil {
		return nil, err
	}

	resp, err := bot.request(dst, MethodGetChat)
	if err != nil {
		return nil, err
	}

	var chat Chat
	err = parser.Unmarshal(resp.Result, &chat)
	return &chat, err
}

// GetChatAdministrators get a list of administrators in a chat. On success,
// returns an Array of ChatMember objects that contains information about all
// chat administrators except other bots. If the chat is a group or a supergroup
// and no administrators were appointed, only the creator will be returned.
func (bot *Bot) GetChatAdministrators(chatID int64) ([]ChatMember, error) {
	dst, err := parser.Marshal(&GetChatAdministratorsParameters{ChatID: chatID})
	if err != nil {
		return nil, err
	}

	resp, err := bot.request(dst, MethodGetChatAdministrators)
	if err != nil {
		return nil, err
	}

	var chatMembers []ChatMember
	err = parser.Unmarshal(resp.Result, &chatMembers)
	return chatMembers, err
}

// GetChatMember get information about a member of a chat. Returns a ChatMember
// object on success.
func (bot *Bot) GetChatMember(chatID int64, userID int) (*ChatMember, error) {
	dst, err := parser.Marshal(&GetChatMemberParameters{
		ChatID: chatID,
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}

	resp, err := bot.request(dst, MethodGetChatMember)
	if err != nil {
		return nil, err
	}

	var chatMember ChatMember
	err = parser.Unmarshal(resp.Result, &chatMember)
	return &chatMember, err
}

// GetChatMembersCount get the number of members in a chat. Returns Int on
// success.
func (bot *Bot) GetChatMembersCount(chatID int64) (int, error) {
	dst, err := parser.Marshal(&GetChatMembersCountParameters{ChatID: chatID})
	if err != nil {
		return 0, err
	}

	resp, err := bot.request(dst, MethodGetChatMembersCount)
	if err != nil {
		return 0, err
	}

	var count int
	err = parser.Unmarshal(resp.Result, &count)
	return count, err
}

// GetFile get basic info about a file and prepare it for downloading. For the
// moment, bots can download files of up to 20MB in size. On success, a File
// object is returned. The file can then be downloaded via the link
// https://api.telegram.org/file/bot<token>/<file_path>, where <file_path> is
// taken from the response. It is guaranteed that the link will be valid for at
// least 1 hour. When the link expires, a new one can be requested by calling
// getFile again.
//
// Note: This function may not preserve the original file name and MIME type. You
// should save the file's MIME type and name (if available) when the File object
// is received.
func (bot *Bot) GetFile(fileID string) (*File, error) {
	dst, err := parser.Marshal(&GetFileParameters{FileID: fileID})
	if err != nil {
		return nil, err
	}

	resp, err := bot.request(dst, MethodGetFile)
	if err != nil {
		return nil, err
	}

	var file File
	err = parser.Unmarshal(resp.Result, &file)
	return &file, err
}

// GetMe testing your bot's auth token. Requires no parameters. Returns basic
// information about the bot in form of a User object.
func (bot *Bot) GetMe() (*User, error) {
	resp, err := bot.request(nil, MethodGetMe)
	if err != nil {
		return nil, err
	}

	var me User
	err = parser.Unmarshal(resp.Result, &me)
	return &me, err
}

// GetUpdates receive incoming updates using long polling. An Array of Update objects is returned.
func (bot *Bot) GetUpdates(params *GetUpdatesParameters) ([]Update, error) {
	if params == nil {
		params = &GetUpdatesParameters{Limit: 100}
	}

	src, err := parser.Marshal(params)
	if err != nil {
		return nil, err
	}

	resp, err := bot.request(src, MethodGetUpdates)
	if err != nil {
		return nil, err
	}

	updates := make([]Update, params.Limit)
	err = parser.Unmarshal(resp.Result, &updates)
	return updates, err
}

// GetUserProfilePhotos get a list of profile pictures for a user. Returns a UserProfilePhotos object.
func (bot *Bot) GetUserProfilePhotos(params *GetUserProfilePhotosParameters) (*UserProfilePhotos, error) {
	dst, err := parser.Marshal(params)
	if err != nil {
		return nil, err
	}

	resp, err := bot.request(dst, MethodGetUserProfilePhotos)
	if err != nil {
		return nil, err
	}

	var photos UserProfilePhotos
	err = parser.Unmarshal(resp.Result, &photos)
	return &photos, err
}

// GetWebhookInfo get current webhook status. Requires no parameters. On success,
// returns a WebhookInfo object. If the bot is using getUpdates, will return an
// object with the url field empty.
func (bot *Bot) GetWebhookInfo() (*WebhookInfo, error) {
	resp, err := bot.request(nil, MethodGetWebhookInfo)
	if err != nil {
		return nil, err
	}

	var info WebhookInfo
	err = parser.Unmarshal(resp.Result, &info)
	return &info, err
}

// GetGameHighScores get data for high score tables. Will return the score of the
// specified user and several of his neighbors in a game. On success, returns an
// Array of GameHighScore objects.
func (bot *Bot) GetGameHighScores(params *GetGameHighScoresParameters) ([]GameHighScore, error) {
	dst, err := parser.Marshal(params)
	if err != nil {
		return nil, err
	}

	resp, err := bot.request(dst, MethodGetGameHighScores)
	if err != nil {
		return nil, err
	}

	var scores []GameHighScore
	err = parser.Unmarshal(resp.Result, &scores)
	return scores, err
}

// GetStickerSet get a sticker set. On success, a StickerSet object is returned.
func (bot *Bot) GetStickerSet(name string) (*StickerSet, error) {
	dst, err := parser.Marshal(&GetStickerSetParameters{Name: name})
	if err != nil {
		return nil, err
	}

	resp, err := bot.request(dst, MethodGetStickerSet)
	if err != nil {
		return nil, err
	}

	var set StickerSet
	err = parser.Unmarshal(resp.Result, &set)
	return &set, err
}