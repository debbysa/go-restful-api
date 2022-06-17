package middleware

import (
	"github.com/debbysa/go-restful-api/helper"
	"github.com/debbysa/go-restful-api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	// karena meneruskan request, jadi ada handler disini
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "RAHASIA" == request.Header.Get("X-API-Key") {
		// oke
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		// error
		writer.Header().Set("content-type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)
	}
}
