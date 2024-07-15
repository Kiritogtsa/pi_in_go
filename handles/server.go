package handles

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/Kiritogtsa/pi_in_go/handles/handles"
	"github.com/Kiritogtsa/pi_in_go/internal/regras"
)

var (
	middler = regras.Newlvlsauthorizaction()
	handels = handles.Newhandles()
)

func Serverinit(port string) {
	for {
		r := chi.NewRouter()
		r.Group(func(r chi.Router) {
			r.Use(middler.Gerente)
			r.Route("/secret", admimrotes)
		})
		r.Group(func(r chi.Router) {
			r.Use(middler.AthoUser)
			r.Route("/private", privatesrotes)
		})
		r.Group(func(r chi.Router) {
			r.Route("/public", Publicassrotas)
		})
		http.ListenAndServe(port, r)
	}
}

// define as rotas privadas
func admimrotes(r chi.Router) {
	r.Get("/user/all", handels.Getallusers)
	r.Get("/user/{id}", handels.Getuser)
	r.Delete("/deletaruser", handels.Deleteuser)
	r.Post("/trabalhos", handels.Postrabalho)
	r.Put("/", func(w http.ResponseWriter, r *http.Request) {})
}

// define as rotas publicas
func Publicassrotas(r chi.Router) {
	r.Post("/login", handels.Login)
	r.Get("/trabalhos", handels.Gettrabalhos)
}

// rotas privadas para o usuario comum
func privatesrotes(r chi.Router) {
	r.Post("/perfil", handels.Perfil)
	r.Get("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(string("oi")))
	})
}
