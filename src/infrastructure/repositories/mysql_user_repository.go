package repositories

import (
	"Recu_ArqSoftware/src/domain/entities"
	"gorm.io/gorm"
)

type MySQLUserRepository struct {
	db *gorm.DB
}

func NewMySQLUserRepository(db *gorm.DB) *MySQLUserRepository {
	return &MySQLUserRepository{db: db}
}

func (r *MySQLUserRepository) Save(user *entities.User) error {
	return r.db.Create(user).Error
}
