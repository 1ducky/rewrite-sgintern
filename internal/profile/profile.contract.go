package profile

type UsecaseContract interface {
	CreateUser(req CreateRequest) error
	UpdateProfile(req UpdateRequest) error
	GetUser(id string) (Profile, error)
	DeleteUser(id string) error

	UploadAvatar(id string) error
	UploadCover(id string) error
}

type RepoContract interface {
	CreateUser(req CreateData) error
	UpdateProfile(req UpdateRequest) error
	GetUser(id string) (Profile, error)
	DeleteUser(id string) error
}
