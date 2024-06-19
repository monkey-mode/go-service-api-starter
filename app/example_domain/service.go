package example

type Service interface {
	Process(cid string) (string, error)
	// TODO: Add your service method
}

func NewService(repository Repository) Service {
	return &ServiceImpl{Repository: repository}
}

type ServiceImpl struct {
	Repository Repository
	// TODO: Add your service fields
}

func (s *ServiceImpl) Process(cid string) (string, error) {
	// TODO: Add your service method
	return s.Repository.Get(cid)
}
