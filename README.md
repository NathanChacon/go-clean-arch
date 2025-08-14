# Go Job Board API
This project is a backend API built with Go that demonstrates modern backend practices, emphasizing clean architecture and scalable design. It provides a robust foundation for building maintainable and efficient server-side applications.

## Key Features
Clean Architecture: Clear separation of concerns with distinct layers for domain entities, use cases, and infrastructure, promoting maintainability and testability.

Authentication: Secure login system implemented using JSON Web Tokens (JWT) for stateless authentication.

Caching: Redis integration to cache frequently accessed routes, improving performance and reducing database load.

Database: Persistent data storage using MySQL, handling user data and job postings.

Core Functionality: Users can register, log in, view available jobs, and publish new job listings.

This project showcases practical usage of backend technologies and architectural principles, helping me build strong foundations for scalable backend systems.

## Environment Variables
To run this project locally, create a .env file in the root directory and configure the following environment variables:

```env
PORT=<your-server-port>

DB_URL=<your-mysql-connection-string>

JWT_SECRET=<your-jwt-secret-key>

REDIS_URL=<your-redis-connection-string>

PORT — The port number on which the server will listen (e.g., 8080)

DB_URL — MySQL connection string (e.g., user:password@tcp(localhost:3306)/dbname).

JWT_SECRET — Secret key used to sign and verify JWT tokens.

REDIS_URL — Redis connection URL (e.g., redis://localhost:6379).

Make sure to replace the placeholders with your actual credentials.
