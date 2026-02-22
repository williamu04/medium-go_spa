package userapplication

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

func NewUserUseCase(repository repository.UserRepository, passwordHasher *pkg.Hasher, tokenGenerator *pkg.JWTGen, sluger *pkg.Sluger) *UserUseCase {
	return &UserUseCase{
		Registration: NewRegistrationUseCase(repository, *passwordHasher, *tokenGenerator, *sluger),
		Login:        NewLoginUserUseCase(repository, *passwordHasher, *tokenGenerator),
		Retrieve:     NewRetrieveUserUseCase(repository),
		RetrieveAll:  NewRetrieveAllUsersUseCase(repository),
		Update:       NewUpdateUserUseCase(repository, *sluger),
		Delete:       NewDeleteUserUseCase(repository),
	}
}
