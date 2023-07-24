package main

import (
	"github.com/ditrit/badaas-example/controllers"
	"github.com/ditrit/badaas/router/middlewares"
	"github.com/gorilla/mux"
)

// Initialize example routes
func AddExampleRoutes(
	router *mux.Router,
	jsonController middlewares.JSONController,
	helloController controllers.HelloController,
) {
	router.HandleFunc(
		"/hello",
		jsonController.Wrap(helloController.SayHello),
	).Methods("GET")
}
