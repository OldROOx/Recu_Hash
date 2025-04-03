package usecases

import (
	"Recu_ArqSoftware/src/domain/entities"
	"Recu_ArqSoftware/src/domain/repositories"
	"Recu_ArqSoftware/src/domain/services"
)

type CreateUserUseCase struct {
	userRepo        repositories.UserRepository
	passwordService services.PasswordService
}

func NewCreateUserUseCase(userRepo repositories.UserRepository, passwordService services.PasswordService) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepo:        userRepo,
		passwordService: passwordService,
	}
}

func (uc *CreateUserUseCase) Execute(username, password string) (*entities.User, error) {
	hashedPassword, err := uc.passwordService.Hash(password)
	if err != nil {
		return nil, err
	}

	user := &entities.User{
		Username: username,
		Password: hashedPassword,
	}

	if err := uc.userRepo.Save(user); err != nil {
		return nil, err
	}

	return user, nil
}
