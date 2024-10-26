package user

import "github.com/bagastri07/gin-boilerplate-service/internal/model/contract"

type Handler struct {
	userUseCase contract.UserUseCase
}

func New() *Handler {
	return new(Handler)
}

func (h *Handler) WithUserUseCase(uc contract.UserUseCase) *Handler {
	h.userUseCase = uc
	return h
}
