package telegram

// ExportChatInviteLinkParameters represents data for ExportChatInviteLink method.
type ExportChatInviteLinkParameters struct {
	// Unique identifier for the target chat
	ChatID int64 `json:"chat_id"`
}

// ExportChatInviteLink export an invite link to a supergroup or a channel. The
// bot must be an administrator in the chat for this to work and must have the
// appropriate admin rights. Returns exported invite link as String on success.
func (bot *Bot) ExportChatInviteLink(chatID int64) (string, error) {
	dst, err := parser.Marshal(&ExportChatInviteLinkParameters{ChatID: chatID})
	if err != nil {
		return "", err
	}

	resp, err := bot.request(dst, MethodExportChatInviteLink)
	if err != nil {
		return "", err
	}

	var inviteLink string
	err = parser.Unmarshal(resp.Result, &inviteLink)
	return inviteLink, err
}
