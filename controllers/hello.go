package controllers

import (
	"net/http"

	"github.com/ditrit/badaas/httperrors"
)

type HelloController interface {
	SayHello(http.ResponseWriter, *http.Request) (any, httperrors.HTTPError)
}

type helloControllerImpl struct{}

func NewHelloController() HelloController {
	return &helloControllerImpl{}
}

func (*helloControllerImpl) SayHello(response http.ResponseWriter, r *http.Request) (any, httperrors.HTTPError) {
	return "hello world", nil
}
