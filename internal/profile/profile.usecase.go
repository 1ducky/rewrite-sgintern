package profile

import (
	"RewriteProject/internal/db/mapper"
	"context"
	"strings"
)

type Usecase struct {
	repo RepoContract
}

func NewUsecase(repo RepoContract) UsecaseContract {
	return &Usecase{repo: repo}
}

func (uc *Usecase) CreateUser(ctx context.Context, req CreateRequest) error {
	var CreateInput CreateData
	if req.ID == "" {
		return ErrInvalidUsername
	}
	CreateInput.ID = req.ID
	if req.Username == "" {
		return ErrInvalidUsername
	}
	CreateInput.Username = req.Username
	CreateInput.Tag = CreateInput.Username[0:4] + strings.Split(req.Username, " ")[0]
	err := uc.repo.CreateUser(ctx, CreateInput)
	if err != nil {
		transalte := mapper.MapMySQLError(err)
		if transalte == mapper.ErrDuplicate {
			return transalte
		}
	}
	return nil
}

func (uc *Usecase) GetUserByID(ctx context.Context, id string) (Profile, error) {
	return Profile{}, nil
}
func (uc *Usecase) UpdateProfile(ctx context.Context, req UpdateRequest) error {
	return nil
}

func (uc *Usecase) DeleteUser(ctx context.Context, id string) error {
	return nil
}

func (uc *Usecase) UploadAvatar(ctx context.Context, id string) error {
	return nil
}

func (uc *Usecase) UploadCover(ctx context.Context, id string) error {
	return nil
}
