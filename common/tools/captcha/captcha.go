package captcha

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/dchest/captcha"
	"github.com/google/uuid"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/syncx"
)

var (
	exclusiveCalls = syncx.NewSharedCalls()
	stats          = cache.NewStat("captcha")
	errorCaptcha   = errors.New("captcha: no rows in result set")

	cachePrefix = "captcha"
	cacheExpiry = 300 * time.Second
)

type (
	Captcha interface {
		RandomDigits(length int) (id, b64s string, err error)
		VerifyString(captchaId string, digits string) bool
	}

	defaultCaptcha struct {
		CacheRedis cache.Cache
	}
)

func NewCaptcha(c cache.CacheConf) Captcha {
	return &defaultCaptcha{
		CacheRedis: cache.New(c, exclusiveCalls, stats, errorCaptcha, func(o *cache.Options) { o.Expiry = cacheExpiry }),
	}
}

func (c *defaultCaptcha) RandomDigits(length int) (id, b64s string, err error) {
	digits := captcha.RandomDigits(length)
	id = uuid.New().String()
	c.CacheRedis.SetWithExpire(fmt.Sprintf("%s:%s", cachePrefix, id), digits, cacheExpiry)
	img := captcha.NewImage(id, digits, 136, 53)

	writer := bytes.Buffer{}
	img.WriteTo(&writer)

	return id, fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(writer.Bytes())), nil
}

func (c *defaultCaptcha) VerifyString(captchaId string, digits string) bool {
	var reald []byte
	key := fmt.Sprintf("%s:%s", cachePrefix, captchaId)
	c.CacheRedis.Get(key, &reald)
	c.CacheRedis.Del(key)
	ns := make([]byte, len(digits))
	for i := range ns {
		d := digits[i]
		switch {
		case '0' <= d && d <= '9':
			ns[i] = d - '0'
		case d == ' ' || d == ',':
			// ignore
		default:
			return false
		}
	}
	return bytes.Equal(ns, reald)
}
