package accountmanagement_test

import (
	"witpgh-jobapi-go/app/shared/database"
	"witpgh-jobapi-go/app/shared/repositories"
	"witpgh-jobapi-go/app/shared/services/system/generation"

	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestAddEmployer(t *testing.T) {

	os.Clearenv()
	err := godotenv.Load("/Users/airblair/projects/go/src/witpgh-jobapi-go/doc.env")

	if err != nil {
		log.Print(err)
	} else {
		database.ConnectWITJobBoard()

		assert := assert.New(t)

		var genService = generation.NewGenerationService()
		var employerKey = genService.GeneratePublicId()
		var publicId = genService.GeneratePublicId()
		var email = genService.GeneratePublicId() + "@" + genService.GeneratePublicId() + ".com"

		var myRepository = repositories.NewRepositoryRegistry().GetEmployerAccountRepository()
		p, err := myRepository.AddNewEmployer(publicId, employerKey, email, "pasword123", "Test", "Client")

		assert.Nil(err)
		assert.True(p.Id > 0, "Employer Id must be greater than 1")
		assert.Equal(employerKey, p.EmployerKey, "Employer Keys should be equal")
		assert.Equal(publicId, p.PublicId, "Public Ids should be equal")
		assert.Equal(email, p.Email, "Email Addresses should be equal")

		myRepository.DeleteEmployerById(p.Id)
	}
}
