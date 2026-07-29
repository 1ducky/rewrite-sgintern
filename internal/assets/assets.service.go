package assets

type Service struct {
	StorageRepo RepositoryContract
}

func NewService(StorageRepo RepositoryContract) Service {
	return Service{StorageRepo: StorageRepo}
}
