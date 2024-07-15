package handles

import (
	"net/http"

	"github.com/Kiritogtsa/pi_in_go/utils"
)

type Handle struct {
	util utils.Utilsi
}

func Newhandles() *Handle {
	var util utils.Utilsi = utils.Newutil()
	return &Handle{util: util}
}

func (u *Handle) Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(string("teste")))
}

func (u *Handle) Gettrabalhos(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(string("oi")))
}

func (u *Handle) Getallusers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(string("oi")))
}

func (u *Handle) Postuser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(string("oi")))
}

func (u *Handle) Postrabalho(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(string("oi")))
}

func (u *Handle) Putuser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(string("oi")))
}

func (u *Handle) Puttrabalho(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(string("oi")))
}

func (u *Handle) Getuser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(string("oi")))
}

func (u *Handle) Deleteuser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(string("oi")))
}

func (u *Handle) Perfil(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(string("oi")))
}
