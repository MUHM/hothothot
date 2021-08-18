package otp

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/png"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func GenerateKey(issuer string, accountName string) (*otp.Key, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: accountName,
	})
	return key, err
}

func KeyToImage(key *otp.Key) (string, error) {
	var buf bytes.Buffer
	img, err := key.Image(200, 200)
	png.Encode(&buf, img)
	return fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(buf.Bytes())), err
}

func Validate(passcode string, secret string) bool {
	valid := totp.Validate(passcode, secret)
	return valid
}
