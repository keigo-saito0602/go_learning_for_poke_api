package di

import (
	"github.com/keigo-saito0602/go_learning_for_poke_api/usecase"
	"gorm.io/gorm"
)

type Usecases struct {
	User usecase.UserUsecase
	Memo usecase.MemoUsecase
}

func NewUsecases(db *gorm.DB, repository *Repositories) *Usecases {
	return &Usecases{
		User: usecase.NewUserUsecase(db, repository.User),
		Memo: usecase.NewMemoUsecase(db, repository.Memo),
	}
}
