package encryption

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type EncryptionService struct {
}

func NewEncryptionService() *EncryptionService {
	return &EncryptionService{}
}

func (service *EncryptionService) HashString(password string) string {
	key, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return ""
	}

	p := string(key)

	return p
}

// MatchString returns true if the hash matches the password
func (service *EncryptionService) MatchString(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err == nil {
		return true
	} else {
		return false
	}
}
