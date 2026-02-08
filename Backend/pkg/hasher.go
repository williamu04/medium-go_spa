package pkg

import "golang.org/x/crypto/bcrypt"

type Hasher struct {
	cost int
}

func NewHasher() *Hasher {
	return &Hasher{cost: bcrypt.DefaultCost}
}

func (h *Hasher) Hash(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	return string(hashed), err
}

func (h *Hasher) Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
