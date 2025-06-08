package validation

import "github.com/keigo-saito0602/go_learning_for_poke_api/infrastructure/repository"

type Validators struct {
	User *UserValidator
	Memo *MemoValidator
}

func NewValidators(userRepo repository.UserRepository) *Validators {
	return &Validators{
		User: NewUserValidator(userRepo),
		Memo: NewMemoValidator(NewUserValidator(userRepo)),
	}
}
