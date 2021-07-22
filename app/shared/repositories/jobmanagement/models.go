package jobmanagement

import (
	"time"
)

type EmployerJob struct {
	Id                    int       `json:"id"`
	PublicId              string    `json:"public_id"`
	EmployerId            int       `json:"employer_id"`
	EmployerCompanyId     int       `json:"employer_company_id"`
	IsClosed              int       `json:"is_closed"`
	Title                 string    `json:"title"`
	CompanyHQ             string    `json:"company_hq"`
	JobTypeId             int       `json:"job_type_id"`
	RegionalRestrictionId int       `json:"regional_restriction_id"`
	Description           string    `json:"description"`
	CloseDate             time.Time `json:"close_date"`
}
