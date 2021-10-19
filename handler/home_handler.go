package handler

import (
	"github.com/rysmaadit/go-template/common/responder"
	"net/http"
)

func Home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := []byte("Welcome")
		responder.NewHttpResponse(r, w, http.StatusOK, res, nil)
	}
}
