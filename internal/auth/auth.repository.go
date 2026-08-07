package auth

import (
	"context"
	"database/sql"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) RepositoryContract {
	return &Repository{DB: db}
}

func (r *Repository) FindByEmail(ctx context.Context, email string) (AuthEntity, error) {
	var user AuthEntity
	query := `SELECT id,username, role, version FROM users WHERE email = $1`
	row := r.DB.QueryRowContext(ctx, query, email)
	if err := row.Scan(&user.ID, &user.Username, &user.Role, &user.Version); err != nil {
		return AuthEntity{}, err
	}
	return user, nil
}

func (r *Repository) Login(ctx context.Context, email string) (AuthLogin, error) {
	var user AuthLogin
	query := `SELECT id,email, password FROM users WHERE email = $1`
	row := r.DB.QueryRowContext(ctx, query, email)
	if err := row.Scan(&user.ID, &user.Email, &user.Password); err != nil {
		return AuthLogin{}, err
	}
	return user, nil
}

func (r *Repository) Create(ctx context.Context, entity RegisterPayload) (AuthEntity, error) {
	return AuthEntity{}, nil
}

func (r *Repository) GetVersion(ctx context.Context, userId string) (int, error) {
	return 0, nil
}
