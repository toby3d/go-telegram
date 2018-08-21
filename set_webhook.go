package telegram

import (
	"strconv"
	"strings"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// SetWebhookParameters represents data for SetWebhook method.
type SetWebhookParameters struct {
	// HTTPS url to send updates to. Use an empty string to remove webhook
	// integration
	URL string `json:"url"`

	// Upload your public key certificate so that the root certificate in use can
	// be checked. See our self-signed guide for details.
	Certificate InputFile `json:"certificate,omitempty"`

	// Maximum allowed number of simultaneous HTTPS connections to the webhook
	// for update delivery, 1-100. Defaults to 40. Use lower values to limit the
	// load on your bot‘s server, and higher values to increase your bot’s
	// throughput.
	MaxConnections int `json:"max_connections,omitempty"`

	// List the types of updates you want your bot to receive. For example,
	// specify [“message”, “edited_channel_post”, “callback_query”] to only
	// receive updates of these types. See Update for a complete list of
	// available update types. Specify an empty list to receive all updates
	// regardless of type (default). If not specified, the previous setting will
	// be used.
	//
	// Please note that this parameter doesn't affect updates created before the
	// call to the setWebhook, so unwanted updates may be received for a short
	// period of time.
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

// NewWebhook creates new SetWebhookParameters only with required parameters.
func NewWebhook(url string, file interface{}) *SetWebhookParameters {
	return &SetWebhookParameters{
		URL:         url,
		Certificate: file,
	}
}

// SetWebhook specify a url and receive incoming updates via an outgoing webhook.
// Whenever there is an update for the bot, we will send an HTTPS POST request to
// the specified url, containing a JSON-serialized Update. In case of an
// unsuccessful request, we will give up after a reasonable amount of attempts.
// Returns true.
//
// If you'd like to make sure that the Webhook request comes from Telegram, we
// recommend using a secret path in the URL, e.g. https://www.example.com/<token>.
// Since nobody else knows your bot‘s token, you can be pretty sure it’s us.
func (bot *Bot) SetWebhook(params *SetWebhookParameters) (ok bool, err error) {
	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.Add("url", params.URL)

	if len(params.AllowedUpdates) > 0 {
		args.Add("allowed_updates", strings.Join(params.AllowedUpdates, ","))
	}
	if params.MaxConnections > 0 &&
		params.MaxConnections <= 100 {
		args.Add("max_connections", strconv.Itoa(params.MaxConnections))
	}

	var resp *Response
	if params.Certificate != nil {
		resp, err = bot.Upload(MethodSetWebhook, "certificate", "cert.pem", params.Certificate, args)
	} else {
		var dst []byte
		dst, err = json.Marshal(params)
		if err != nil {
			return
		}

		resp, err = bot.request(dst, MethodSetWebhook)
	}
	if err != nil {
		return
	}

	err = json.Unmarshal(*resp.Result, &ok)
	return
}
