package user

import "github.com/bagastri07/gin-boilerplate-service/internal/model/contract"

type UseCase struct {
	userRepository contract.UserRepository
}

func New() *UseCase {
	return new(UseCase)
}

func (u *UseCase) WithUserRepository(repo contract.UserRepository) *UseCase {
	u.userRepository = repo
	return u
}
