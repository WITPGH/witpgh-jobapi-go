package accountmanagement

type Employer struct {
	Id                int    `json:"id"`
	PublicId          string `json:"public_id"`
	EmployerKey       string `json:"employer_key"`
	Status            int    `json:"status"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	MustResetPassword int    `json:"must_reset_password"`
	Firstname         string `json:"firstname"`
	Lastname          string `json:"lastname"`
	SendMobileNotices int    `json:"send_mobile_notices"`
}
