package di

import (
	"github.com/keigo-saito0602/go_learning_for_poke_api/infrastructure/db"
	"github.com/keigo-saito0602/go_learning_for_poke_api/interface/handler"
	"github.com/keigo-saito0602/go_learning_for_poke_api/validation"
	"gorm.io/gorm"
)

type Container struct {
	DB       *gorm.DB
	Repos    *Repositories
	Usecases *Usecases
	Handlers *handler.Handlers
}

func NewContainer() *Container {
	db := db.NewMySQLDB()
	repos := NewRepositories(db)
	usecases := NewUsecases(db, repos)
	validators := validation.NewValidators(repos.User)
	handlers := NewHandlers(usecases, validators)

	return &Container{
		DB:       db,
		Repos:    repos,
		Usecases: usecases,
		Handlers: handlers,
	}
}
