package example

func NewApp() Handler {
	repository := NewRepository()
	service := NewService(repository)

	return NewHandler(service)
}
