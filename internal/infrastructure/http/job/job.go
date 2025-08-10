package jobHandlers

import (
	"encoding/json"
	"errors"
	"net/http"

	domainErrors "jobs.api.com/internal/domain/errors"
	"jobs.api.com/internal/infrastructure/middlewares"
	jobUsecase "jobs.api.com/internal/usecases/job"

	"github.com/gorilla/mux"
)

func isErrorInList(err error, list []error) bool {
	for _, e := range list {
		if errors.Is(err, e) {
			return true
		}
	}
	return false
}

type JobHandler struct {
	usecase jobUsecase.UseCase
}

func NewJobHandler(usecase jobUsecase.UseCase) *JobHandler {
	return &JobHandler{usecase: usecase}
}

func (handler *JobHandler) PostJob(writer http.ResponseWriter, request *http.Request) {
	authData := request.Context().Value("authData").(middlewares.AuthData)
	var badRequestErrors = []error{
		domainErrors.ErrJobCompanyIDRequired,
		domainErrors.ErrJobCreatedByRequired,
		domainErrors.ErrJobDescriptionRequired,
		domainErrors.ErrJobDescriptionTooShort,
		domainErrors.ErrJobLocationRequired,
		domainErrors.ErrJobMissingRequiredField,
	}

	var job jobUsecase.JobDTO

	if err := json.NewDecoder(request.Body).Decode(&job); err != nil {
		http.Error(writer, "Invalid payload", http.StatusBadRequest)
		return
	}

	job.CreatedBy = authData.Id

	err := handler.usecase.PostJob(job)

	if isErrorInList(err, badRequestErrors) {
		http.Error(writer, "Invalid payload", http.StatusBadRequest)
		return
	}

	if errors.Is(err, domainErrors.ErrUserNotFound) {
		http.Error(writer, "Failed to post job because user dont exists", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
}

func (handler *JobHandler) GetJobByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	job, err := handler.usecase.GetById(id)
	if err != nil {
		http.Error(w, "Job not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(job)
}
