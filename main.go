package main

// implement user confirm email flow ?

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	authHandler "jobs.api.com/internal/infrastructure/http/authentication"
	jobHandlers "jobs.api.com/internal/infrastructure/http/job"
	userHandler "jobs.api.com/internal/infrastructure/http/user"
	"jobs.api.com/internal/infrastructure/middlewares"
	jobRepository "jobs.api.com/internal/infrastructure/respository/job"
	userRepository "jobs.api.com/internal/infrastructure/respository/user"
	"jobs.api.com/internal/infrastructure/utils/passwordHasher"
	uuidGenerator "jobs.api.com/internal/infrastructure/utils/uuid"
	authenticationUseCase "jobs.api.com/internal/usecases/authentication"
	jobUsecase "jobs.api.com/internal/usecases/job"
	userUseCase "jobs.api.com/internal/usecases/user"
)

func initializeDb() (*sqlx.DB, error) {
	dsn := os.Getenv("DB_URL")
	dsn = dsn + "?parseTime=true"
	db, err := sqlx.Connect("mysql", dsn)

	return db, err
}

func initializeRedis() *redis.Client {
	redisUrl := os.Getenv("REDIS_URL")
	redisClient := redis.NewClient(&redis.Options{
		Addr: redisUrl,
		DB:   0,
	})

	return redisClient
}

func main() {
	ctx := context.Background()
    if _, err := os.Stat(".env"); err == nil {
        if err := godotenv.Load(); err != nil {
            log.Println("Warning: could not load .env file:", err)
        }
    }
	db, err := initializeDb()

	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
		return
	}
	defer db.Close()

	redisClient := initializeRedis()

	if err := redisClient.Ping(ctx).Err(); err != nil {
		fmt.Println("Failed to connect to Redis: %v", err)
	}

	uuidGenerator := uuidGenerator.NewUuidGenerator()
	passwordHasher := passwordHasher.NewPasswordHasher()

	userRepo := userRepository.NewUserRepository(db)
	userUseCases := userUseCase.NewUserUseCase(userRepo, uuidGenerator, passwordHasher)
	userHandler := userHandler.NewUserHandler(userUseCases)

	jobRepo := jobRepository.NewJobRepository(db)
	jobUseCases := jobUsecase.NewJobUseCase(jobRepo, uuidGenerator)
	jobHandler := jobHandlers.NewJobHandler(jobUseCases, redisClient)

	authenticationUseCase := authenticationUseCase.NewAutheticationUseCase(userRepo, passwordHasher)
	authenticationHandler := authHandler.NewAuthHandler(authenticationUseCase)

	router := mux.NewRouter()

	router.HandleFunc("/create-account", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/login", authenticationHandler.Login).Methods("POST")

	protectedRoutes := router.PathPrefix("/").Subrouter()
	protectedRoutes.HandleFunc("/user/{id}", userHandler.GetUserById).Methods("GET")
	protectedRoutes.HandleFunc("/jobs", jobHandler.PostJob).Methods("POST")
	protectedRoutes.HandleFunc("/jobs", jobHandler.GetAll).Methods("GET")
	protectedRoutes.HandleFunc("/jobs/{id}", jobHandler.GetJobByID).Methods("GET")

	protectedRoutes.Use(middlewares.AuthMiddleware)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
