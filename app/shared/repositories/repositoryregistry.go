package repositories

import (
	"witpgh-jobapi-go/app/shared/database"
	"witpgh-jobapi-go/app/shared/repositories/accountmanagement"
)

type RepositoryRegistry struct {
}

func NewRepositoryRegistry() *RepositoryRegistry {
	return &RepositoryRegistry{}
}

func (registry *RepositoryRegistry) GetEmployerAccountRepository() *accountmanagement.AccountRepository {
	return accountmanagement.NewAccountRepository(database.WITPGH)
}
