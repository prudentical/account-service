package app

import (
	"account-service/configuration"
	"account-service/internal/discovery"
	"context"
	"fmt"
	"log/slog"

	"go.uber.org/fx"
)

type AppSetupManager interface {
	Setup() error
	Shutdown() error
}

type appSetupManagerImpl struct {
	app       RESTApp
	discovery discovery.ServiceDiscovery
}

func NewAppSetupManager(app RESTApp, discovery discovery.ServiceDiscovery) AppSetupManager {
	return appSetupManagerImpl{app, discovery}
}

func (a appSetupManagerImpl) Setup() error {
	err := a.app.setup()
	if err != nil {
		return err
	}

	err = a.discovery.Register()
	if err != nil {
		return err
	}

	return nil
}

func (a appSetupManagerImpl) Shutdown() error {
	return nil
}

func ManageLifeCycle(lc fx.Lifecycle, config configuration.Config, log *slog.Logger, app RESTApp, manager AppSetupManager) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("Starting the server")
			err := manager.Setup()
			if err != nil {
				return err
			}
			go app.server().Start(fmt.Sprintf(":%v", config.Server.Port))
			return err
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Shuting down the server")
			err := manager.Shutdown()
			if err != nil {
				return err
			}
			return app.server().Shutdown(ctx)
		},
	})
}
