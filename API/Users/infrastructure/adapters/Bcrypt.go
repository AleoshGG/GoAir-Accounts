package adapters

import "golang.org/x/crypto/bcrypt"

type Bcrypt struct {}

func NewBcrypt() *Bcrypt {
	return &Bcrypt{}
}

func (hs *Bcrypt) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (hs *Bcrypt) ValidatePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}