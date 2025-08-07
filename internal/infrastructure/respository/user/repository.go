package userRepository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	userEntity "jobs.api.com/internal/domain/entities/user"
	domainErrors "jobs.api.com/internal/domain/errors"
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

func isDuplicateErr(err error) bool {
	var mysqlErr *mysql.MySQLError
	return errors.As(err, &mysqlErr) && mysqlErr.Number == 1062
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(payload userEntity.User) error {
	query := `
        INSERT INTO users (uuid, email, password, name, location, bio)
        VALUES (?, ?, ?, ?, ?, ?)
    `
	_, err := r.db.Exec(query, payload.UUID, payload.Email, payload.Password, payload.Name, payload.Location, payload.Bio)
	if isDuplicateErr(err) {
		return domainErrors.ErrUserAlreadyRegistered
	}
	fmt.Print(payload)

	fmt.Print(err)
	return err
}

func (r *UserRepository) GetById(id string) (userEntity.User, error) {
	var user UserModel
	var newUserEntity userEntity.User
	err := r.db.Get(&user, `SELECT * FROM users WHERE uuid = ?`, id)

	if errors.Is(err, sql.ErrNoRows) {
		return newUserEntity, domainErrors.ErrUserNotFound
	}

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
