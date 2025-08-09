package main

// to do - improve domain error separating by use cases. Ex user errors should be in package userErrors etc
import (
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	authHandler "jobs.api.com/internal/infrastructure/http/authentication"
	jobHandlers "jobs.api.com/internal/infrastructure/http/job"
	userHandler "jobs.api.com/internal/infrastructure/http/user"
	jobRepository "jobs.api.com/internal/infrastructure/respository/job"
	userRepository "jobs.api.com/internal/infrastructure/respository/user"
	"jobs.api.com/internal/infrastructure/utils/passwordHasher"
	uuidGenerator "jobs.api.com/internal/infrastructure/utils/uuid"
	authenticationUseCase "jobs.api.com/internal/usecases/authentication"
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

	uuidGenerator := uuidGenerator.NewUuidGenerator()
	passwordHasher := passwordHasher.NewPasswordHasher()

	userRepo := userRepository.NewUserRepository(db)
	userUseCases := userUseCase.NewUserUseCase(userRepo, uuidGenerator, passwordHasher)
	userHandler := userHandler.NewUserHandler(userUseCases)

	jobRepo := jobRepository.NewJobMySQLRepository(db)
	jobUseCases := jobUsecase.NewJobUseCase(jobRepo, uuidGenerator)
	jobHandler := jobHandlers.NewJobHandler(jobUseCases)

	authenticationUseCase := authenticationUseCase.NewAutheticationUseCase(userRepo, passwordHasher)
	authenticationHandler := authHandler.NewAuthHandler(authenticationUseCase)

	router := mux.NewRouter()

	router.HandleFunc("/user/{id}", userHandler.GetUserById).Methods("GET")
	router.HandleFunc("/create-account", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/login", authenticationHandler.Login).Methods("POST")
	router.HandleFunc("/jobs", jobHandler.PostJob).Methods("POST")
	router.HandleFunc("/jobs/{id}", jobHandler.GetJobByID).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
