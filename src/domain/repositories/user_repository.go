package repositories

import (
	"Recu_ArqSoftware/src/domain/entities"
)

type UserRepository interface {
	Save(user *entities.User) error
}
