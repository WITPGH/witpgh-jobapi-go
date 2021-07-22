package jobmanagement

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type JobRepository struct {
	db *sql.DB
}

func NewJobRepository(db *sql.DB) *JobRepository {
	return &JobRepository{db}
}

func (repository *JobRepository) AddNewEmployerJob(publicId string, employerId int, employerCompanyId int, isClosed int, title string, companyHQ string, jobTypeId int, regionalRestrictionId int, description string, closeDate string) (*EmployerJob, error) {
	var jobId int
	stmt, err := repository.db.Prepare(`INSERT into employer_jobs ("public_id", "employer_id", "employer_company_id", "is_closed", "job_title", "company_hq", "job_type_id", "regional_restriction_id", "job_description", "close_date") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	stmtErr := stmt.QueryRow(publicId, employerId, employerCompanyId, isClosed, title, companyHQ, jobTypeId, regionalRestrictionId, description, closeDate).Scan(&jobId)

	if stmtErr != nil {
		return nil, err
	}

	return repository.GetJobById(jobId)
}

func (repository *JobRepository) GetJobById(jobId int) (*EmployerJob, error) {
	var result EmployerJob
	err := repository.db.QueryRow("select id, public_id, employer_id, employer_company_id, is_closed, job_title, company_hq, job_type_id, regional_restriction_id, job_description, close_date from employer_jobs where id = $1 limit 1", jobId).Scan(&result.Id, &result.PublicId, &result.EmployerId, &result.EmployerCompanyId, &result.IsClosed, &result.Title, &result.CompanyHQ, &result.JobTypeId, &result.RegionalRestrictionId, &result.Description, &result.CloseDate)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			log.Println(err)
		}
	}

	return &result, err
}

func (repository *JobRepository) GetJobByPublicId(publicId string) (*EmployerJob, error) {
	var result EmployerJob
	err := repository.db.QueryRow("select id, public_id, employer_id, employer_company_id, is_closed, job_title, company_hq, job_type_id, regional_restriction_id, job_description, close_date from employer_jobs where public_id = $1 limit 1", publicId).Scan(&result.Id, &result.PublicId, &result.EmployerId, &result.EmployerCompanyId, &result.IsClosed, &result.Title, &result.CompanyHQ, &result.JobTypeId, &result.RegionalRestrictionId, &result.Description, &result.CloseDate)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			log.Println(err)
		}
	}

	return &result, err
}

func (repository *JobRepository) GetJobsByEmployerId(employerId int) ([]EmployerJob, error) {
	var items []EmployerJob
	rows, err := repository.db.Query("select id, public_id, employer_id, employer_company_id, is_closed, job_title, company_hq, job_type_id, regional_restriction_id, job_description, close_date from employer_jobs where employer_id = $1", employerId)

	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			var result EmployerJob
			rows.Scan(&result.Id, &result.PublicId, &result.EmployerId, &result.EmployerCompanyId, &result.IsClosed, &result.Title, &result.CompanyHQ, &result.JobTypeId, &result.RegionalRestrictionId, &result.Description, &result.CloseDate)
			items = append(items, result)
		}
	}

	return items, err
}

func (repository *JobRepository) GetEmployerJobsByRegion(employerId int, regionalRestrictionId int) ([]EmployerJob, error) {
	var items []EmployerJob
	rows, err := repository.db.Query("select id, public_id, employer_id, employer_company_id, is_closed, job_title, company_hq, job_type_id, regional_restriction_id, job_description, close_date from employer_jobs where employer_id = $1 and regional_restriction_id = $2", employerId, regionalRestrictionId)

	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			var result EmployerJob
			rows.Scan(&result.Id, &result.PublicId, &result.EmployerId, &result.EmployerCompanyId, &result.IsClosed, &result.Title, &result.CompanyHQ, &result.JobTypeId, &result.RegionalRestrictionId, &result.Description, &result.CloseDate)
			items = append(items, result)
		}
	}

	return items, err
}

func (repository *JobRepository) GetEmployerRemoteJobs(employerId int) ([]EmployerJob, error) {
	var items []EmployerJob
	rows, err := repository.db.Query("select id, public_id, employer_id, employer_company_id, is_closed, job_title, company_hq, job_type_id, regional_restriction_id, job_description, close_date from employer_jobs where employer_id = $1 and regional_restriction_id = $2", employerId, 1)

	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			var result EmployerJob
			rows.Scan(&result.Id, &result.PublicId, &result.EmployerId, &result.EmployerCompanyId, &result.IsClosed, &result.Title, &result.CompanyHQ, &result.JobTypeId, &result.RegionalRestrictionId, &result.Description, &result.CloseDate)
			items = append(items, result)
		}
	}

	return items, err
}

func (repository *JobRepository) GetOpenEmployer(employerId int) ([]EmployerJob, error) {
	var items []EmployerJob
	rows, err := repository.db.Query("select id, public_id, employer_id, employer_company_id, is_closed, job_title, company_hq, job_type_id, regional_restriction_id, job_description, close_date from employer_jobs where employer_id = $1 and is_closed = $2", employerId, 0)

	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			var result EmployerJob
			rows.Scan(&result.Id, &result.PublicId, &result.EmployerId, &result.EmployerCompanyId, &result.IsClosed, &result.Title, &result.CompanyHQ, &result.JobTypeId, &result.RegionalRestrictionId, &result.Description, &result.CloseDate)
			items = append(items, result)
		}
	}

	return items, err
}

func (repository *JobRepository) GetOpenEmployerRemoteJobs(employerId int) ([]EmployerJob, error) {
	var items []EmployerJob
	rows, err := repository.db.Query("select id, public_id, employer_id, employer_company_id, is_closed, job_title, company_hq, job_type_id, regional_restriction_id, job_description, close_date from employer_jobs where employer_id = $1 and is_closed = $2 and regional_restriction_id = $3", employerId, 0, 1)

	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			var result EmployerJob
			rows.Scan(&result.Id, &result.PublicId, &result.EmployerId, &result.EmployerCompanyId, &result.IsClosed, &result.Title, &result.CompanyHQ, &result.JobTypeId, &result.RegionalRestrictionId, &result.Description, &result.CloseDate)
			items = append(items, result)
		}
	}

	return items, err
}

func (repository *JobRepository) GetOpenEmployerJobsByRegion(employerId int, regionalRestrictionId int) ([]EmployerJob, error) {
	var items []EmployerJob
	rows, err := repository.db.Query("select id, public_id, employer_id, employer_company_id, is_closed, job_title, company_hq, job_type_id, regional_restriction_id, job_description, close_date from employer_jobs where employer_id = $1 and is_closed = $2 and regional_restriction_id = $3", employerId, 0, regionalRestrictionId)

	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			var result EmployerJob
			rows.Scan(&result.Id, &result.PublicId, &result.EmployerId, &result.EmployerCompanyId, &result.IsClosed, &result.Title, &result.CompanyHQ, &result.JobTypeId, &result.RegionalRestrictionId, &result.Description, &result.CloseDate)
			items = append(items, result)
		}
	}

	return items, err
}

func (repository *JobRepository) GetAllClosedEmployerJobs(employerId int) ([]EmployerJob, error) {
	var items []EmployerJob
	rows, err := repository.db.Query("select id, public_id, employer_id, employer_company_id, is_closed, job_title, company_hq, job_type_id, regional_restriction_id, job_description, close_date from employer_jobs where employer_id = $1 and is_closed = $2", employerId, 1)

	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			var result EmployerJob
			rows.Scan(&result.Id, &result.PublicId, &result.EmployerId, &result.EmployerCompanyId, &result.IsClosed, &result.Title, &result.CompanyHQ, &result.JobTypeId, &result.RegionalRestrictionId, &result.Description, &result.CloseDate)
			items = append(items, result)
		}
	}

	return items, err
}

func (repository *JobRepository) CloseEmployerJob(employerId, jobId int) (bool, error) {
	stmt, err := repository.db.Prepare("UPDATE employer_jobs set is_closed = 1, update_date = $1 where id = $2 and employer_id = $3")

	if err != nil {
		log.Println(err)
		return false, err
	}
	rtf, scanError := stmt.Exec(time.Now(), jobId, employerId)

	if scanError != nil {
		log.Println(scanError)
		return false, scanError
	}

	recordCount, updateErr := rtf.RowsAffected()

	if recordCount > 0 {
		return true, nil
	} else {
		return false, updateErr
	}
}
