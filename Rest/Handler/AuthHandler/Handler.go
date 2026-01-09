package authhandler

type Handler struct {
	service *authService
}

func NewHandler(service *authService) *Handler {
	return &Handler{
		service: service,
	}
}
