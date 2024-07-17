package handles

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/Kiritogtsa/pi_in_go/internal/domains/entries"
	"github.com/Kiritogtsa/pi_in_go/internal/repository/trabalhos"
	"github.com/Kiritogtsa/pi_in_go/internal/repository/users"
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
	userdao := users.Nuserre()
	var userpost entries.User
	err := json.NewDecoder(r.Body).Decode(&userpost)
	if err != nil {
		http.Error(w, "nao foi possivel obter o json", http.StatusInternalServerError)
		return
	}
	user, err := userdao.GetByemail(userpost.Email)
	if err != nil {
		http.Error(w, "nao foi possivel obter o json", http.StatusInternalServerError)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(string(user.Senha)), []byte(string(userpost.Senha)))
	if err != nil {
		http.Error(w, "senha invalida", http.StatusNonAuthoritativeInfo)
		return
	}
	w.Write([]byte(string("teste")))
}

func (u *Handle) Gettrabalhos(w http.ResponseWriter, r *http.Request) {
	trabalhosdao := trabalhos.Newtradb()
	trabalhos, err := trabalhosdao.Getall()
	if err != nil {
		http.Error(w, "not entries trabalhos", http.StatusInternalServerError)
		return
	}
	json, err := json.Marshal(trabalhos)
	if err != nil {
		http.Error(w, "nao foi possivel virar json", http.StatusInternalServerError)
		return
	}
	w.Write(json)
}

func (u *Handle) Getallusers(w http.ResponseWriter, r *http.Request) {
	usersdao := users.Nuserre()
	users, err := usersdao.Getall()
	if err != nil {
		http.Error(w, "not entries users", http.StatusInternalServerError)
		return
	}
	json, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "nao foi possivel virar json", http.StatusInternalServerError)
		return
	}
	w.Write(json)
}

func (u *Handle) Postuser(w http.ResponseWriter, r *http.Request) {
	var user entries.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(" -- ", err)
		http.Error(w, "json invalido", http.StatusInternalServerError)
		return
	}
	userdao := users.Nuserre()
	err = userdao.Persit(&user)
	if err != nil {
		fmt.Println(" -- ", err)
		http.Error(w, "deu algo erro ao criar o usuario", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(string("oi")))
}

func (u *Handle) Postrabalho(w http.ResponseWriter, r *http.Request) {
	var trabalho entries.Trabalhos
	err := json.NewDecoder(r.Body).Decode(&trabalho)
	if err != nil {
		fmt.Println(" -- ", err)
		http.Error(w, "json invalido", http.StatusInternalServerError)
		return
	}
	trabalhosdao := trabalhos.Newtradb()
	err = trabalhosdao.Persit(&trabalho)
	if err != nil {
		fmt.Println(" -- ", err)
		http.Error(w, "deu algo erro ao criar o usuario", http.StatusInternalServerError)
		return
	}
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
