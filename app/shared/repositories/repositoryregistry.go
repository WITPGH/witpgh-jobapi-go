package repositories

import (
	"witpgh-jobapi-go/app/shared/database"
	"witpgh-jobapi-go/app/shared/repositories/accountmanagement"
	"witpgh-jobapi-go/app/shared/repositories/companymanagement"
	"witpgh-jobapi-go/app/shared/repositories/jobmanagement"
)

type RepositoryRegistry struct {
}

func NewRepositoryRegistry() *RepositoryRegistry {
	return &RepositoryRegistry{}
}

func (registry *RepositoryRegistry) GetEmployerAccountRepository() *accountmanagement.AccountRepository {
	return accountmanagement.NewAccountRepository(database.WITPGH)
}

func (registry *RepositoryRegistry) GetEmployerCompanyRepository() *companymanagement.CompanyRepository {
	return companymanagement.NewCompanyRepository(database.WITPGH)
}

func (registry *RepositoryRegistry) GetEmployerJobRepository() *jobmanagement.JobRepository {
	return jobmanagement.NewJobRepository(database.WITPGH)
}
