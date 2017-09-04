package telegram

import (
	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// GetFile get basic info about a file and prepare it for downloading. For the moment, bots can download files of up to 20MB in size. On success, a File object is returned. The file can then be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>, where <file_path> is taken from the response. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile again.
//
// Note: This function may not preserve the original file name and MIME type. You should save the file's MIME type and name (if available) when the File object is received.
func (bot *Bot) GetFile(file string) (*File, error) {
	var args http.Args
	args.Add("file_id", file) // File identifier to get info about

	resp, err := bot.get("getFile", &args)
	if err != nil {
		return nil, err
	}

	var data File
	err = json.Unmarshal(*resp.Result, &data)
	return &data, err
}