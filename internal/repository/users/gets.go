package users

import (
	"log"

	"github.com/Kiritogtsa/pi_in_go/internal/domains/entries"
)

type Selects interface {
	Getall() ([]entries.User, error)
	GetByemail(string) (*entries.User, error)
}

func (u *userre) Getall() ([]entries.User, error) {
	conn := u.conn.Conndb
	query := "select * from users"
	result, err := conn.Query(query)
	if err != nil {
		return nil, err
	}
	var users []entries.User
	for result.Next() {
		user := &entries.User{}
		err := result.Scan(&user.ID, &user.Name, &user.Email, &user.Cpf, &user.DataAd, &user.DataNas, &user.Senha, &user.Grupo, &user.Telefone, &user.TrID)
		if err != nil {
			log.Println("Erro ao escanear linha:", err)
			return nil, err
		}
		users = append(users, *user)
	}
	return users, nil
}

func (u *userre) GetByemail(string) (*entries.User, error) {
	return nil, nil
}
