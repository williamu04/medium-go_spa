package user_application

import (
	"github.com/williamu04/medium-clone/domain/repository"
	"github.com/williamu04/medium-clone/pkg"
)

type UserUseCase struct {
	Registration *RegistrationUseCase
	Login        *LoginUserUseCase
	Retrieve     *RetrieveUserUseCase
	RetrieveAll  *RetrieveAllUsersUseCase
	Update       *UpdateUserUseCase
	Delete       *DeleteUserUseCase
}

func NewUserUseCase(repository repository.UserRepository, passwordHasher *pkg.Hasher, tokenGenerator *pkg.JWTGen) *UserUseCase {
	return &UserUseCase{
		Registration: NewRegistrationUseCase(repository, *passwordHasher, *tokenGenerator),
		Login:        NewLoginUserUseCase(repository, *passwordHasher, *tokenGenerator),
		Retrieve:     NewRetrieveUserUseCase(repository),
		RetrieveAll:  NewRetrieveAllUsersUseCase(repository),
		Update:       NewUpdateUserUseCase(repository),
		Delete:       NewDeleteUserUseCase(repository),
	}
}
