package services

import (
	"witpgh-jobapi-go/app/shared/services/system"
)

type ServiceRegistry struct {
}

func NewServiceRegistry() *ServiceRegistry {
	return &ServiceRegistry{}
}

func (service *ServiceRegistry) GetSystemServiceRegistry() *system.SystemServiceRegistry {
	return system.NewSystemServiceRegistry()
}
