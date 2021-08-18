package captcha

import (
	"bytes"
	"encoding/base64"
	"fmt"

	"github.com/dchest/captcha"
	"github.com/google/uuid"
)

func RandomDigits(length int) (id, b64s string, err error) {
	digits := captcha.RandomDigits(length)
	id = uuid.New().String()
	// captcha.Store.Set(string(digits),string(digits))
	img := captcha.NewImage(id, digits, 136, 53)

	writer := bytes.Buffer{}
	img.WriteTo(&writer)

	return id, fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(writer.Bytes())), nil
}

func VerifyString(captchaId string, digits string) bool {
	return captcha.VerifyString(captchaId, digits)
}
