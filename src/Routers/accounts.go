package Routers

import (
	"account-service/Respond"
	"account-service/Services"
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func AccountRouter() http.Handler {
	r := chi.NewRouter()
	r.With(MyMiddleware).Post("/", createAccount)
	r.With(MyMiddleware).Get("/", findAllAccount)
	return r
}

func MyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// create new context from `r` request context, and assign key `"user"`
		// to value of `"123"`
		ctx := context.WithValue(r.Context(), "user", "123")

		// call the next handler in the chain, passing the response writer and
		// the updated request object with the new context value.
		//
		// note: context.Context values are nested, so any previously set
		// values will be accessible as well, and the new `"user"` key
		// will be accessible from this point forward.

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type accountReq struct {
	AccountId string `json:"account_id" `
	Name      string `json:"name"`
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	p := accountReq{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		panic(err)
	}

	data := Services.AccountModel{
		Id:        p.AccountId,
		IdAccount: p.Name,
	}
	Respond.RespondWithJSON(w, http.StatusCreated, data)
}

func findAllAccount(w http.ResponseWriter, r *http.Request) {
	Respond.RespondWithJSON(w, http.StatusOK, Services.FindAllAccounts())
}
