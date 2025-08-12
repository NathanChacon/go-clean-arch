package jobHandlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	domainErrors "jobs.api.com/internal/domain/errors"
	"jobs.api.com/internal/infrastructure/middlewares"
	jobUsecase "jobs.api.com/internal/usecases/job"

	"github.com/go-redis/redis/v8"
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
	redis   *redis.Client
}

func NewJobHandler(usecase jobUsecase.UseCase, redis *redis.Client) *JobHandler {
	return &JobHandler{usecase: usecase, redis: redis}
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

func (handler *JobHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	const cacheKey = "jobs:all"

	cached, err := handler.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(cached))
		return
	}

	jobs, err := handler.usecase.GetAll()
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(jobs)

	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}

	err = handler.redis.Set(ctx, cacheKey, data, 5*time.Minute).Err()
	if err != nil {
		log.Printf("Redis SET error: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
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
