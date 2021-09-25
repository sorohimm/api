package infrastructure

import (
	"api/internal/config"
	"api/internal/controllers"
	"api/internal/interfaces"
	"api/internal/repos"
	"api/internal/services"
	"context"

	"go.uber.org/zap"
)

type IInjector interface {
	InjectBookController() controllers.BookControllers
}

var env *environment

type environment struct {
	logger   *zap.SugaredLogger
	dbClient interfaces.IDBHandler
}

func (e *environment) InjectBookController() controllers.BookControllers {
	return controllers.BookControllers{
		Log: e.logger,
		BookService: &services.BookService{
			DBHandler:  e.dbClient,
			DBBookRepo: &repos.DBBookRepo{},
		},
	}
}

func Injector(logger *zap.SugaredLogger, ctx context.Context, cfg *config.Config) (IInjector, error) {
	dbClient, err := initPostgresClient(cfg, ctx)
	if err != nil {
		return nil, err
	}

	env = &environment{
		logger:   logger,
		dbClient: dbClient,
	}

	return env, nil
}
