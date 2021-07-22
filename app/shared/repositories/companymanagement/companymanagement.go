package companymanagement

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type CompanyRepository struct {
	db *sql.DB
}

func NewCompanyRepository(db *sql.DB) *CompanyRepository {
	return &CompanyRepository{db}
}

func (repository *CompanyRepository) AddNewCompany(publicId string, employerId int, name string, statement string, logo string, url string, description string, addressLine1 string, addressLine2 string, city string, state string, zipcode string, country string, latitude float64, longitude float64) (*Company, error) {
	var companyId int
	stmt, err := repository.db.Prepare(`INSERT into employer_companies ("public_id", "employer_id", "status", "company_name", "company_statement", "company_logo", "has_logo", "company_url", "company_description", "address_line1", "address_line2", "city", "state", "zipcode", "country", "latitude", "longitude") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17) RETURNING id`)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	stmtErr := stmt.QueryRow(publicId, employerId, 1, name, statement, logo, 1, url, description, addressLine1, addressLine2, city, state, zipcode, country, latitude, longitude).Scan(&companyId)

	if stmtErr != nil {
		return nil, err
	}

	return repository.GetCompanyById(companyId)
}

func (repository *CompanyRepository) GetCompanyById(companyId int) (*Company, error) {
	var result Company
	err := repository.db.QueryRow("select id, public_id, employer_id, status, company_name, company_statement, company_logo, has_logo, company_url, company_description, address_line1, address_line2, city, state, zipcode, country, latitude, longitude from employer_companies where id = $1 limit 1", companyId).Scan(&result.Id, &result.PublicId, &result.EmployerId, &result.Status, &result.CompanyName, &result.CompanyStatement, &result.CompanyLogo, &result.HasLogo, &result.CompanyURL, &result.CompanyDescription, &result.AddressLine1, &result.AddressLine2, &result.City, &result.State, &result.Zipcode, &result.Country, &result.Latitude, &result.Longitude)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			log.Println(err)
		}
	}

	return &result, err
}

func (repository *CompanyRepository) GetCompaniesByEmployerId(employerId int) ([]Company, error) {
	var items []Company
	rows, err := repository.db.Query("select id, public_id, employer_id, status, company_name, company_statement, company_logo, has_logo, company_url, company_description, address_line1, address_line2, city, state, zipcode, country, latitude, longitude from employer_companies where employer_id = $1", employerId)

	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			var result Company
			rows.Scan(&result.Id, &result.PublicId, &result.EmployerId, &result.Status, &result.CompanyName, &result.CompanyStatement, &result.CompanyLogo, &result.HasLogo, &result.CompanyURL, &result.CompanyDescription, &result.AddressLine1, &result.AddressLine2, &result.City, &result.State, &result.Zipcode, &result.Country, &result.Latitude, &result.Longitude)
			items = append(items, result)
		}
	}

	return items, err
}

func (repository *CompanyRepository) UpdateCompany(companyId int, employerId int, name string, statement string, logo string, url string, description string, addressLine1 string, addressLine2 string, city string, state string, zipcode string, country string, latitude float64, longitude float64) (*Company, error) {
	stmt, err := repository.db.Prepare("UPDATE employer_companies set company_name = $1, company_statement = $2, company_logo = $3, company_url = $4, company_description = $5, address_line1 = $6, address_line2 = $7, city = $8, state = $9, zipcode = $10, country = $11, latitude = $12, longitude = $13,  update_date = $14 where id = $15 and employer_id = $16")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, scanError := stmt.Exec(name, state, logo, url, description, addressLine1, addressLine2, city, state, zipcode, country, latitude, longitude, time.Now(), companyId, employerId)

	if scanError != nil {
		log.Println(scanError)
		return nil, scanError
	}

	return repository.GetCompanyById(companyId)
}

func (repository *CompanyRepository) ActivateCompany(employerId, companyId int) (bool, error) {
	stmt, err := repository.db.Prepare("UPDATE employer_companies set status = 1, update_date = $1 where id = $2 and employer_id = $3")

	if err != nil {
		log.Println(err)
		return false, err
	}
	rtf, scanError := stmt.Exec(time.Now(), companyId, employerId)

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

func (repository *CompanyRepository) InactivateCompany(employerId, companyId int) (bool, error) {
	stmt, err := repository.db.Prepare("UPDATE employer_companies set status = 0, update_date = $1 where id = $2 and employer_id = $3")

	if err != nil {
		log.Println(err)
		return false, err
	}
	rtf, scanError := stmt.Exec(time.Now(), companyId, employerId)

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
