package profile

import (
	"errors"
	"strings"
)

type Usecase struct {
	repo RepoContract
}

func NewUsecase(repo RepoContract) UsecaseContract {
	return &Usecase{repo: repo}
}

func (uc *Usecase) CreateUser(req CreateRequest) error {
	var CreateInput CreateData
	if req.ID == "" {
		return errors.New("id is required")
	}
	CreateInput.ID = req.ID
	if req.Username == "" {
		return errors.New("username is required")
	}
	CreateInput.Username = req.Username
	CreateInput.Tag = CreateInput.Username[0:4] + strings.Split(req.Username, " ")[0]
	err := uc.repo.CreateUser(CreateInput)
	if err != nil {
		return err
	}
	return nil
}

func (uc *Usecase) UpdateProfile(req UpdateRequest) error {
	return nil
}

func (uc *Usecase) GetUser(id string) (Profile, error) {
	return Profile{}, nil
}

func (uc *Usecase) DeleteUser(id string) error {
	return nil
}

func (uc *Usecase) UploadAvatar(id string) error {
	return nil
}

func (uc *Usecase) UploadCover(id string) error {
	return nil
}
