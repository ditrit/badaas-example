package main

import (
	"net/http"

	"github.com/Masterminds/semver/v3"
	"github.com/ditrit/badaas"
	badaasControllers "github.com/ditrit/badaas/controllers"
	badaasModels "github.com/ditrit/badaas/persistence/models"
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

		fx.Provide(NewAPIVersion),
		// add routes provided by badaas
		badaasControllers.InfoControllerModule,
		// badaasControllers.AuthControllerModule,
		badaasControllers.EAVControllerModule,

		// start example routes
		fx.Provide(NewHelloController),
		fx.Invoke(AddExampleRoutes),

		// start example eav data
		// fx.Invoke(CreateEAVCRUDObjects),

		// start example data
		badaasControllers.GetCRUDModule[badaasModels.Product](),
		badaasControllers.GetCRUDModule[badaasModels.Sale](),
		fx.Provide(NewEntityMapping),
		badaasControllers.CRUDControllerModule,
		// fx.Invoke(CreateCRUDObjects),

		// create httpServer
		fx.Provide(NewHTTPServer),
		// Finally: we invoke the newly created server
		fx.Invoke(func(*http.Server) { /* we need this function to be empty*/ }),
	).Run()
}

func NewAPIVersion() *semver.Version {
	return semver.MustParse("0.0.0-unreleased")
}

type EntityMappingParams struct {
	fx.In

	// TODO ver como sacar este models.
	// TODO esto hasta se podria hacer automaticamente
	ProductsCRUDController badaasControllers.CRUDController `name:"models.ProductCRUDController"`
	SalesCRUDController    badaasControllers.CRUDController `name:"models.SaleCRUDController"`
}

func NewEntityMapping(params EntityMappingParams) map[string]badaasControllers.CRUDController {
	return map[string]badaasControllers.CRUDController{
		"product": params.ProductsCRUDController,
		"sale":    params.SalesCRUDController,
	}
}
