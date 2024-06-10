package app

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"

	"account-service/internal/api"
	"account-service/internal/configuration"
	"account-service/internal/database"
	"account-service/internal/discovery"
	"account-service/internal/persistence"
	"account-service/internal/service"
	"account-service/internal/util"
)

type Application interface {
	Run()
}

func NewApplication() Application {
	return FxContainer{}
}

type FxContainer struct {
}

func (FxContainer) Run() {
	fx.New(
		fx.Provide(configuration.NewConfig),
		fx.Provide(NewLogger),
		fx.Provide(NewFxLogger),
		fx.Provide(ProvideEcho),
		fx.Provide(NewAppSetupManager),
		fx.Provide(discovery.NewServiceDiscovery),
		fx.Provide(util.NewValidator),
		fx.Provide(api.NewHTTPErrorHandler),
		fx.Provide(database.NewDatabaseConnection),
		fx.Provide(persistence.NewAccountDAO),
		fx.Provide(service.NewAccountService),
		asHandler(api.NewAccountHandler),
		asHandler(api.NewHealthCheck),
		fx.Provide(fx.Annotate(
			NewRESTApp,
			fx.ParamTags(`group:"handlers"`),
		)),
		fx.WithLogger(func(log FxLogger) fxevent.Logger {
			return &log
		}),
		fx.Invoke(ManageLifeCycle),
	).Run()
}

func asHandler(handler interface{}) fx.Option {
	return fx.Provide(fx.Annotate(
		handler,
		fx.As(new(api.Handler)),
		fx.ResultTags(`group:"handlers"`),
	))
}
