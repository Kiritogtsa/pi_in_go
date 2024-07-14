package regras

import (
	"net/http"

	"github.com/gorilla/sessions"
)

type middlerlvlsauthoriazactionit interface {
	AthoUser(http.Handler) http.Handler
	Gerente(http.Handler) http.Handler
	Gerenteoraxiliar(http.Handler) http.Handler
	Login(http.ResponseWriter, *http.Request)
	Gettrabalhos(http.ResponseWriter, *http.Request)
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

func (a *middlerlvlsauthoriazaction) Login(w http.ResponseWriter, r *http.Request) {
}

func (a *middlerlvlsauthoriazaction) Gerenteoraxiliar(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
}

func (a *middlerlvlsauthoriazaction) Gerente(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
}

func (a *middlerlvlsauthoriazaction) Gettrabalhos(w http.ResponseWriter, r *http.Request) {
}
