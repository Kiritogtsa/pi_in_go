package handles

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/Kiritogtsa/pi_in_go/internal/regras"
)

var middler = regras.Newlvlsauthorizaction()

func Serverinit(port string) {
	time.Sleep(1000 * time.Millisecond * 10)
	for {
		r := chi.NewRouter()
		r.Group(func(r chi.Router) {
			r.Use(middler.Gerente)
			r.Route("/secret", privatesrotes)
		})
		r.Group(func(r chi.Router) {
			r.Use(middler.AthoUser)
			r.Route("/private", privatesrotes)
		})
		r.Group(func(r chi.Router) {
			r.Use(middler.AthoUser)
			r.Route("/", inithome)
		})
		r.Group(func(r chi.Router) {
			r.Route("/public", Publicassrotas)
		})
		http.ListenAndServe(port, r)
	}
}

func privatesrotes(r chi.Router) {
	r.Get("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(string("oi")))
	})
	r.Get("/user/ativados", func(w http.ResponseWriter, r *http.Request) {})
	r.Get("/user/desativados", func(w http.ResponseWriter, r *http.Request) {})
	r.Get("/user/all", func(w http.ResponseWriter, r *http.Request) {})
	r.Get("/user/{id}", func(w http.ResponseWriter, r *http.Request) {})
	r.Put("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(string("oi")))
	})
	r.Post("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(string("oi")))
	})
	r.Delete("/deletaruser", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(string("oi")))
	})
	r.Post("/trabalhos", func(w http.ResponseWriter, r *http.Request) {})
	r.Put("/", func(w http.ResponseWriter, r *http.Request) {})
}

func Publicassrotas(r chi.Router) {
	r.Post("/login", middler.Login)
	r.Get("/trabalhos", middler.Gettrabalhos)
}

func inithome(r chi.Router) {
	r.Get("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(string("oi")))
	})
}
