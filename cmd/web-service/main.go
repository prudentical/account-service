package main

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"

	"account-service/internal/api"
	"account-service/internal/app"
	"account-service/internal/configuration"
	"account-service/internal/database"
	"account-service/internal/discovery"
	"account-service/internal/persistence"
	"account-service/internal/service"
	"account-service/internal/util"
)

func main() {
	fx.New(
		fx.Provide(configuration.NewConfig),
		fx.Provide(app.NewLogger),
		fx.Provide(app.NewFxLogger),
		fx.Provide(app.ProvideEcho),
		fx.Provide(app.NewAppSetupManager),
		fx.Provide(discovery.NewServiceDiscovery),
		fx.Provide(util.NewValidator),
		fx.Provide(api.NewHTTPErrorHandler),
		fx.Provide(database.NewDatabaseConnection),
		fx.Provide(persistence.NewAccountDAO),
		fx.Provide(service.NewAccountService),
		asHandler(api.NewAccountHandler),
		asHandler(api.NewHealthCheck),
		fx.Provide(fx.Annotate(
			app.NewRESTApp,
			fx.ParamTags(`group:"handlers"`),
		)),
		fx.WithLogger(func(log app.FxLogger) fxevent.Logger {
			return &log
		}),
		fx.Invoke(app.ManageLifeCycle),
	).Run()
}

func asHandler(handler interface{}) fx.Option {
	return fx.Provide(fx.Annotate(
		handler,
		fx.As(new(api.Handler)),
		fx.ResultTags(`group:"handlers"`),
	))
}
