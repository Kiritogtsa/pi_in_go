package regras

import (
	"net/http"

	"github.com/gorilla/sessions"
)

type middlerlvlsauthoriazactionit interface {
	AthoUser(http.Handler) http.Handler
	Gerente(http.Handler) http.Handler
	Gerenteoraxiliar(http.Handler) http.Handler
}

type middlerlvlsauthoriazaction struct {
	Store sessions.Store
}

func Newlvlsauthorizaction() middlerlvlsauthoriazactionit {
	return &middlerlvlsauthoriazaction{}
}

func (a *middlerlvlsauthoriazaction) AthoUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func (a *middlerlvlsauthoriazaction) Gerenteoraxiliar(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
}

func (a *middlerlvlsauthoriazaction) Gerente(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
}
