package accountmanagement

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db}
}

func (repository *AccountRepository) AddNewEmployer(publicId string, employerKey string, email string, password string, firstname string, lastname string) (*Employer, error) {
	var employerId int
	stmt, err := repository.db.Prepare(`INSERT into employers ("public_id", "employer_key", "status", "email", "password", "must_reset_password", "firstname", "lastname", "send_mobile_notices") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	stmtErr := stmt.QueryRow(publicId, employerKey, 1, email, password, 1, firstname, lastname, 0).Scan(&employerId)

	if stmtErr != nil {
		return nil, err
	}

	return repository.GetEmployerById(employerId)
}

func (repository *AccountRepository) GetEmployerById(employerId int) (*Employer, error) {
	var result Employer
	err := repository.db.QueryRow("select id, public_id, employer_key, status, email, password, must_reset_password, firstname, lastname, send_mobile_notices from employers where id = $1 limit 1", employerId).Scan(&result.Id, &result.PublicId, &result.EmployerKey, &result.Status, &result.Email, &result.Password, &result.MustResetPassword, &result.Firstname, &result.Lastname, &result.SendMobileNotices)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			log.Println(err)
		}
	}

	return &result, err
}

func (repository *AccountRepository) UpdateEmployerEmailAddress(employerId int, emailAddress string) (*Employer, error) {
	stmt, err := repository.db.Prepare("UPDATE employers set email = $1, update_date = $2 where id = $3")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, scanError := stmt.Exec(emailAddress, time.Now(), employerId)

	if scanError != nil {
		log.Println(scanError)
		return nil, scanError
	}

	return repository.GetEmployerById(employerId)
}

func (repository *AccountRepository) UpdateEmployerPassword(employerId int, password string) (*Employer, error) {
	stmt, err := repository.db.Prepare("UPDATE employers set password = $1, update_date = $2 where id = $3")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, scanError := stmt.Exec(password, time.Now(), employerId)

	if scanError != nil {
		log.Println(scanError)
		return nil, scanError
	}

	return repository.GetEmployerById(employerId)
}

func (repository *AccountRepository) UpdateEmployerMobileNotice(employerId int, mobileNotice int) (*Employer, error) {
	stmt, err := repository.db.Prepare("UPDATE employers set send_mobile_notices = $1, update_date = $2 where id = $3")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, scanError := stmt.Exec(mobileNotice, time.Now(), employerId)

	if scanError != nil {
		log.Println(scanError)
		return nil, scanError
	}

	return repository.GetEmployerById(employerId)
}

func (repository *AccountRepository) GetAllEmployers() ([]Employer, error) {
	var items []Employer
	rows, err := repository.db.Query("select id, public_id, employer_key, status, email, password, must_reset_password, firstname, lastname, send_mobile_notices from employers")

	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			var result Employer
			rows.Scan(&result.Id, &result.PublicId, &result.EmployerKey, &result.Status, &result.Email, &result.Password, &result.MustResetPassword, &result.Firstname, &result.Lastname, &result.SendMobileNotices)
			items = append(items, result)
		}
	}

	return items, err
}
