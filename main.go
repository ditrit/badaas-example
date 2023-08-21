package main

import (
	"github.com/ditrit/badaas"
	"github.com/ditrit/badaas/controllers"
)

func main() {
	badaas.BaDaaS.AddModules(
		controllers.AuthControllerModule,
	).Start()
}
