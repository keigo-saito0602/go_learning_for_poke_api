package di

import (
	"github.com/keigo-saito0602/go_learning_for_poke_api/interface/handler"
	"github.com/keigo-saito0602/go_learning_for_poke_api/validation"
)

type Handlers struct {
	User *handler.UserHandler
	Memo *handler.MemoHandler
}

func NewHandlers(usecases *Usecases, validators *validation.Validators) *handler.Handlers {
	return &handler.Handlers{
		User: handler.NewUserHandler(usecases.User, validators.User),
		Memo: handler.NewMemoHandler(usecases.Memo, validators.Memo),
	}
}
