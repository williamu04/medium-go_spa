package pkg

import (
	"crypto/rand"
	"math/big"
)

type Randomer struct {
	logger *Logger
}

func NewRandomer() *Randomer {
	return &Randomer{}
}

const (
	letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	numbers = "0123456789"
)

func (r *Randomer) RandomWord(length int) string {
	if length <= 0 {
		return ""
	}

	b := make([]byte, length)

	for i := range length {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			r.logger.Errorf("Failed to generate random word: %v", err)
		}
		b[i] = letters[idx.Int64()]
	}
	return string(b)
}
