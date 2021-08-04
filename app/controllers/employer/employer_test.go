package employer_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"witpgh-jobapi-go/app/route"
	"witpgh-jobapi-go/app/shared/database"
	"witpgh-jobapi-go/app/shared/repositories"
	"witpgh-jobapi-go/app/shared/services/system/generation"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var (
	testServer *httptest.Server
	reader     io.Reader
)

type new_employer_request struct {
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

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

func init() {

	os.Clearenv()

	err := godotenv.Load("/Users/airblair/projects/go/src/witpgh-jobapi-go/doc.env")

	if err != nil {
		log.Print(err)
	}

	database.ConnectWITJobBoard()

	testServer = httptest.NewServer(route.LoadRoutes())

}

func TestAddEmployer(t *testing.T) {
	assert := assert.New(t)

	var genService = generation.NewGenerationService()

	var employer_request new_employer_request

	employer_request.Firstname = genService.GeneratePublicId()
	employer_request.Lastname = genService.GeneratePublicId()
	employer_request.Email = genService.GeneratePublicId() + "@" + genService.GeneratePublicId() + ".com"

	var newEmployer Employer

	jsonRequest, err := json.Marshal(employer_request)

	if err != nil {
		log.Print("json marshalling error")
		log.Print(err)
	}

	var creationUrl = fmt.Sprintf("%s/employers/account/create", testServer.URL)

	request, err := http.NewRequest("POST", creationUrl, bytes.NewBuffer(jsonRequest))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)

	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		t.Error("Request is Unauthorized")
	}

	if response.StatusCode == http.StatusOK {
		err := json.Unmarshal(body, &newEmployer)

		if err != nil {
			t.Error(err)
		} else {
			assert.True(newEmployer.Id > 0, "Employer Id must be greater than 1")
			assert.NotEmpty(newEmployer.EmployerKey, "Employer Key should not be empty")
			assert.NotEmpty(newEmployer.PublicId, "Public Id should not be empty")
			assert.Equal(employer_request.Email, newEmployer.Email, "Email Addresses should be equal")
			assert.Equal(employer_request.Firstname, newEmployer.Firstname, "Firstnames should be equal")
			assert.Equal(employer_request.Lastname, newEmployer.Lastname, "Lastnames should be equal")

			repositories.NewRepositoryRegistry().GetEmployerAccountRepository().DeleteEmployerById(newEmployer.Id)
		}
	}
}
