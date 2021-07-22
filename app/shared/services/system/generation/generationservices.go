package generation

import (
	"crypto/rand"
	"encoding/hex"
)

type GenerationService struct {
}

func NewGenerationService() *GenerationService {
	return &GenerationService{}
}

func (service *GenerationService) GeneratePublicId() string {
	uuid := make([]byte, 16)
	n, err := rand.Read(uuid)
	if n != len(uuid) || err != nil {
		return ""
	}

	return hex.EncodeToString(uuid)
}
