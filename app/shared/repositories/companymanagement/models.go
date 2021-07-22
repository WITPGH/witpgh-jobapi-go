package companymanagement

type Company struct {
	Id                 int     `json:"id"`
	PublicId           string  `json:"public_id"`
	EmployerId         int     `json:"employer_id"`
	Status             int     `json:"status"`
	CompanyName        string  `json:"name"`
	CompanyStatement   string  `json:"statement"`
	CompanyLogo        string  `json:"logo"`
	HasLogo            int     `json:"has_logo"`
	CompanyURL         string  `json:"url"`
	CompanyDescription string  `json:"description"`
	AddressLine1       string  `json:"address_line_1"`
	AddressLine2       string  `json:"address_line_2"`
	City               string  `json:"city"`
	State              string  `json:"state"`
	Zipcode            string  `json:"zipcode"`
	Country            string  `json:"country"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
}
