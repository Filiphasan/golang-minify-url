package helpers

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
	"time"
)

const (
	Charset       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	CharsetLength = 62
)

type TokenBuilder struct {
	epoch    int64
	addChars int
}

func NewTokenBuilder() *TokenBuilder {
	return &TokenBuilder{
		epoch:    0,
		addChars: 0,
	}
}

func (t *TokenBuilder) SetEpoch(epochDate string) *TokenBuilder {
	parseDate, _ := time.Parse("2006-01-02", epochDate)
	unixMs := parseDate.UnixNano() / 1000000
	t.epoch = unixMs
	return t
}

func (t *TokenBuilder) SetAddChars(addChars int) *TokenBuilder {
	t.addChars = addChars
	return t
}

func (t *TokenBuilder) Build() string {
	epochToken := t.generateTokenFomEpoch()
	addCharsToken := t.generateTokenFromAddChars()
	return fmt.Sprintf("%s%s", epochToken, addCharsToken)
}

func (t *TokenBuilder) generateTokenFomEpoch() string {
	var currentTime = time.Now().UnixNano() / 1000000
	var diff = currentTime - t.epoch
	token := strings.Builder{}

	for diff > CharsetLength {
		token.WriteString(string(Charset[diff%CharsetLength]))
		diff = diff / CharsetLength
	}

	return token.String()
}

func (t *TokenBuilder) generateTokenFromAddChars() string {
	token := strings.Builder{}
	n := big.NewInt(int64(CharsetLength + 1))

	for i := 0; i < t.addChars; i++ {
		rndIndex, _ := rand.Int(rand.Reader, n)
		token.WriteString(string(Charset[int(rndIndex.Int64())]))
	}

	return token.String()
}
