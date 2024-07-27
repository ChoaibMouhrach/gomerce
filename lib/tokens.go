package lib

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/url"
)

func GenerateToken(token *string, n int) error {
	bytes := make([]byte, n)

	if _, err := rand.Read(bytes); err != nil {
		return err
	}

	*token = hex.EncodeToString(bytes)

	return nil
}

func SendMagicLink(email string, token string) error {
	url, err := url.Parse(Env.APP_MAGIC_LINK)

	if err != nil {
		return err
	}

	q := url.Query()
	q.Add("token", token)
	url.RawQuery = q.Encode()

	err = Mail(
		"auth@yerapos.com",
		[]string{email},
		"Authentication",
		fmt.Sprintf("<a href='%s' >Sign In</a>", url.String()),
	)

	if err != nil {
		return err
	}

	return nil
}
