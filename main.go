package main

import (
	"net/http"

	"github.com/Masterminds/semver/v3"
	"github.com/ditrit/badaas"
	"github.com/ditrit/badaas-example/controllers"
	"github.com/ditrit/badaas/badorm"
	"github.com/ditrit/badaas/configuration"
	"github.com/ditrit/badaas/router"
	"github.com/ditrit/verdeter"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

var rootCfg = verdeter.BuildVerdeterCommand(verdeter.VerdeterConfig{
	Use:   "badaas-example",
	Short: "Example of BadAss",
	Long:  "A HTTP server build over BadAas that uses its Login and Object Storage features",
	Run:   runHTTPServer,
})

func main() {
	err := configuration.NewCommandInitializer().Init(rootCfg)
	if err != nil {
		panic(err)
	}

	rootCfg.Execute()
}

// Run the http server for badaas
func runHTTPServer(cmd *cobra.Command, args []string) {
	fx.New(
		fx.Provide(GetModels),
		badaas.BadaasModule,

		// logger for fx
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),

		fx.Provide(NewAPIVersion),
		// add routes provided by badaas
		router.InfoRouteModule,
		router.AuthRoutesModule,

		// start example routes
		fx.Provide(controllers.NewHelloController),
		fx.Invoke(AddExampleRoutes),

		// create httpServer
		fx.Provide(NewHTTPServer),
		// Finally: we invoke the newly created server
		fx.Invoke(func(*http.Server) { /* we need this function to be empty*/ }),
	).Run()
}

func NewAPIVersion() *semver.Version {
	return semver.MustParse("0.0.0-unreleased")
}

func GetModels() badorm.GetModelsResult {
	return badorm.GetModelsResult{
		Models: []any{},
	}
}
