package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	jobHandlers "jobs.api.com/internal/infrastructure/http/job"
	userHandler "jobs.api.com/internal/infrastructure/http/user"
	jobRepository "jobs.api.com/internal/infrastructure/respository/job"
	userRepository "jobs.api.com/internal/infrastructure/respository/user"
	jobUsecase "jobs.api.com/internal/usecases/job"
	userUseCase "jobs.api.com/internal/usecases/user"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
		return
	}

	dsn := os.Getenv("DB_URL")
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
		return
	}
	defer db.Close()

	userRepo := userRepository.NewUserRepository(db)
	userUseCases := userUseCase.NewUserUseCase(userRepo)
	userHandler := userHandler.NewUserHandler(userUseCases)

	jobRepo := jobRepository.NewJobMySQLRepository(db)
	jobUseCases := jobUsecase.NewJobUseCase(jobRepo, userRepo)
	jobHandler := jobHandlers.NewJobHandler(jobUseCases)

	router := mux.NewRouter()

	router.HandleFunc("/user/{id}", userHandler.GetUserById).Methods("GET")
	router.HandleFunc("/jobs", jobHandler.PostJob).Methods("POST")
	router.HandleFunc("/jobs/{id}", jobHandler.GetJobByID).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
