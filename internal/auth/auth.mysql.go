package auth

import (
	"context"
	"database/sql"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) CredentialRepositoryContract {
	return &Repository{DB: db}
}

// func (r *Repository) FindByEmail(ctx context.Context, email string) (AuthEntity, error) {
// 	var user AuthEntity
// 	query := `SELECT id,username, role, version FROM users WHERE email = $1`
// 	row := r.DB.QueryRowContext(ctx, query, email)
// 	if err := row.Scan(&user.ID, &user.Username, &user.Role, &user.Version); err != nil {
// 		return AuthEntity{}, err
// 	}
// 	return user, nil
// }

func (r *Repository) Authentication(ctx context.Context, email string) (AuthLogin, error) {
	var user AuthLogin
	query := `SELECT id,email, password FROM users WHERE email = $1`
	row := r.DB.QueryRowContext(ctx, query, email)
	if err := row.Scan(&user.ID, &user.Email, &user.Password); err != nil {
		return AuthLogin{}, err
	}
	return user, nil
}

func (r *Repository) Registration(ctx context.Context, entity RegisterPayload) error {
	query := `INSERT INTO credential (email,password,user_id) VALUES ($1,$2,$3)`
	_, err := r.DB.ExecContext(ctx, query, entity.Email, entity.Password, entity.UserID)
	if err != nil {
		return err
	}
	return nil
}
