package jobHandlers

import (
	"encoding/json"
	"net/http"

	jobEntity "jobs.api.com/internal/domain/entities/job"
	jobUsecase "jobs.api.com/internal/usecases/job"

	"github.com/gorilla/mux"
)

type JobHandler struct {
	usecase jobUsecase.UseCase
}

func NewJobHandler(usecase jobUsecase.UseCase) *JobHandler {
	return &JobHandler{usecase: usecase}
}

func (handler *JobHandler) PostJob(w http.ResponseWriter, r *http.Request) {
	var job jobEntity.Job

	if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	err := handler.usecase.PostJob(&job)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
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
