package employer

import (
	"encoding/json"
	"log"
	"net/http"
	"witpgh-jobapi-go/app/shared/repositories"
	"witpgh-jobapi-go/app/shared/response"
	"witpgh-jobapi-go/app/shared/services"
)

func CreateEmployerAccount(w http.ResponseWriter, r *http.Request) {

	var request new_employer_request
	decodeErr := json.NewDecoder(r.Body).Decode(&request)

	if decodeErr != nil {
		response.SendJSONMessage(w, http.StatusInternalServerError, response.FriendlyError)
	}

	var serviceRegistry = services.NewServiceRegistry().GetSystemServiceRegistry()
	var publicId = serviceRegistry.GetGenerationService().GeneratePublicId()
	var employerKey = serviceRegistry.GetGenerationService().GeneratePublicId()
	var tempPassword = serviceRegistry.GetGenerationService().GeneratePublicId()
	var hashedPassword = serviceRegistry.GetEncryptionService().HashString(tempPassword)

	var employerRepository = repositories.NewRepositoryRegistry().GetEmployerAccountRepository()
	newEmployer, employerErr := employerRepository.AddNewEmployer(publicId, employerKey, request.Email, hashedPassword, request.Firstname, request.Lastname)

	if employerErr != nil {
		log.Println(employerErr)
		response.SendJSONMessage(w, http.StatusInternalServerError, response.FriendlyError)
	} else {
		//email temporary password
		response.SendJSON(w, newEmployer)
	}

}
