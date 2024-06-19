package example

type Repository interface {
	Get(cid string) (string, error)
	// TODO: Add your repository method
}

func NewRepository() Repository {
	return &RepositoryImpl{}
}

type RepositoryImpl struct {
	// TODO: Add your repository implementation
}

func (r *RepositoryImpl) Get(cid string) (string, error) {
	// TODO: Add your repository method
	return "", nil
}