package users

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/Kiritogtsa/pi_in_go/internal/domains/entries"
)

type Insert interface {
	insert(*entries.User) error
}
type Update interface {
	update(*entries.User) error
}
type persit interface {
	Persit(*entries.User) error
	Insert
	Update
}

func (u *userre) insert(user *entries.User) error {
	conn := u.conn.Conndb
	query := "insert into users(NOME,CPF,EMAIL,DATA_NASCIMENTO,TELEFONE,DATA_ADMISSAO,SEXO,SENHA,GRUPO,TR_ID) values(?,?,?,?,?,?,?,?,?,?,?)"
	stmt, err := conn.Prepare(query)
	if err != nil {
		return err
	}
	senha, err := bcrypt.GenerateFromPassword([]byte(user.Senha), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(user.Name, user.Cpf, user.Email, user.DataNas, user.Telefone, user.DataAd, user.Sexo, senha, user.Grupo, user.TrID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = int(id)
	return nil
}
func (u *userre) update(user *entries.User) error {
	conn := u.conn.Conndb
	query := "update users set NOME=?,EMAIL=?,CPF=?,TELEFONE=?,DATA_ADMISSAO=? where ID = ?"

	stmt, err := conn.Prepare(query)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(user.Name, user.Email, user.Cpf, user.Telefone, user.DataAd)
	if err != nil {
		return err
	}
	linas, err := result.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Println(" linhas mudadas", linas)
	return nil
}
func (u *userre) Persit(user *entries.User) error {
	if user.ID == 0 {
		return u.insert(user)
	} else {
		return u.update(user)
	}
}
