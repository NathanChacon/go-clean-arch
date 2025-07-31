package userRepository

import (
	"github.com/jmoiron/sqlx"

	userEntity "jobs.api.com/internal/domain/entities/user"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(payload userEntity.User) error {

}

func (r *UserRepository) GetById(id string) (userEntity.User, error) {

}
