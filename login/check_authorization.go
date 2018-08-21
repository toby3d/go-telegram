package login

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"net/url"
	"strconv"
)

// ErrUserNotDefined describes error of an unassigned structure of user
var ErrUserNotDefined = errors.New("user is not defined")

// CheckAuthorization verify the authentication and the integrity of the data
// received by comparing the received hash parameter with the hexadecimal
// representation of the HMAC-SHA-256 signature of the data-check-string with the
// SHA256 hash of the bot's token used as a secret key.
func (a *App) CheckAuthorization(user *User) (ok bool, err error) {
	if user == nil {
		err = ErrUserNotDefined
		return
	}

	dataCheck := make(url.Values)
	dataCheck.Add(KeyAuthDate, string(user.AuthDate))
	dataCheck.Add(KeyFirstName, user.FirstName)
	dataCheck.Add(KeyID, strconv.Itoa(user.ID))

	// Add optional values if exist
	if user.LastName != "" {
		dataCheck.Add(KeyLastName, user.LastName)
	}
	if user.PhotoURL != "" {
		dataCheck.Add(KeyPhotoURL, user.PhotoURL)
	}
	if user.Username != "" {
		dataCheck.Add(KeyUsername, user.Username)
	}

	secretKey := sha256.Sum256([]byte(a.SecretKey))
	hash := hmac.New(sha256.New, secretKey[0:])
	if _, err = hash.Write([]byte(dataCheck.Encode())); err != nil {
		return
	}

	ok = hex.EncodeToString(hash.Sum(nil)) == user.Hash
	return
}
