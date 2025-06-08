package di

import (
	"github.com/keigo-saito0602/go_learning_for_poke_api/infrastructure/repository"
	"gorm.io/gorm"
)

type Repositories struct {
	User repository.UserRepository
	Memo repository.MemoRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		User: repository.NewUserRepository(db),
		Memo: repository.NewMemoRepository(db),
	}
}
