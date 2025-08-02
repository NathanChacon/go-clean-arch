package userRepository

import (
	"github.com/jmoiron/sqlx"

	userEntity "jobs.api.com/internal/domain/entities/user"
)

type UserModel struct {
	UUID     string `json:"uuid"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Bio      string `json:"bio"`
}

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(payload userEntity.User) error {
	query := `
        INSERT INTO users (uuid, email, password, name, location, bio)
        VALUES (?, ?, ?, ?, ?, ?, ?)
    `
	_, err := r.db.Exec(query, payload.UUID, payload.Email, payload.Password, payload.Name, payload.Location, payload.Bio)
	return err
}

func (r *UserRepository) GetById(id string) (userEntity.User, error) {
	var user UserModel
	var newUserEntity userEntity.User
	err := r.db.Get(&user, `SELECT * FROM users WHERE uuid = ?`, id)

	if err != nil {
		return newUserEntity, err
	}

	newUserEntity = userEntity.User{
		UUID:     user.UUID,
		Name:     user.Name,
		Password: user.Password,
		Email:    user.Email,
		Bio:      user.Bio,
		Location: user.Location,
	}

	return newUserEntity, err

}
