package code

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/aaronland/go-string/random"
)

func NewSecretCode() (string, error) {

	return NewSecretCodeWithTime(time.Now())
}

func NewSecretCodeWithTime(t time.Time) (string, error) {

	opts := random.DefaultOptions()
	opts.AlphaNumeric = true
	opts.Chars = 64

	code, err := random.String(opts)

	if err != nil {
		return "", err
	}

	ts := t.Unix()

	raw_code := fmt.Sprintf("%d:%s", ts, code)
	sum_code := sha256.Sum256([]byte(raw_code))
	str_code := fmt.Sprintf("%x", sum_code)

	return str_code, nil
}
