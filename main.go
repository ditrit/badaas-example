package main

import (
	"net/http"

	"github.com/ditrit/badaas"
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
	badaas.ConfigCommandParameters(rootCfg)

	rootCfg.Execute()
}

// Run the http server for badaas
func runHTTPServer(cmd *cobra.Command, args []string) {
	fx.New(
		// Modules
		badaas.BadaasModule,

		// logger for fx
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),

		// add routes provided by badaas
		fx.Invoke(router.AddInfoRoutes),
		fx.Invoke(router.AddLoginRoutes),
		fx.Invoke(router.AddCRUDRoutes),

		// start example routes and data
		fx.Provide(NewHelloController),
		fx.Invoke(AddExampleRoutes),
		fx.Invoke(StartExample),

		// create httpServer
		fx.Provide(NewHTTPServer),
		// Finally: we invoke the newly created server
		fx.Invoke(func(*http.Server) { /* we need this function to be empty*/ }),
	).Run()
}
