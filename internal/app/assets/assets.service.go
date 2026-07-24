package assets

type Service struct {
	StorageRepo ReposioturyContract
}

func NewService(StorageRepo ReposioturyContract) Service {
	return Service{StorageRepo: StorageRepo}
}
