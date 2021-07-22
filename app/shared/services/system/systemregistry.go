package system

import (
	"witpgh-jobapi-go/app/shared/services/system/generation"
	"witpgh-jobapi-go/app/shared/services/system/security/encryption"
)

type SystemServiceRegistry struct {
}

func NewSystemServiceRegistry() *SystemServiceRegistry {
	return &SystemServiceRegistry{}
}

func (service *SystemServiceRegistry) GetGenerationService() *generation.GenerationService {
	return generation.NewGenerationService()
}

func (service *SystemServiceRegistry) GetEncryptionService() *encryption.EncryptionService {
	return encryption.NewEncryptionService()
}
