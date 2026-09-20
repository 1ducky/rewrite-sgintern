package profile

import "RewriteProject/internal/db"

type MySQLRepository struct {
	db db.DBTX
}

func NewMySQLRepository(db db.DBTX) RepoContract {
	return &MySQLRepository{db: db}
}

func (r *MySQLRepository) CreateUser(req CreateData) error {
	return nil
}

func (r *MySQLRepository) UpdateProfile(req UpdateRequest) error {
	return nil
}

func (r *MySQLRepository) GetUser(id string) (Profile, error) {
	return Profile{}, nil
}

func (r *MySQLRepository) DeleteUser(id string) error {
	return nil
}
