package trabalhos

import (
	"fmt"

	"github.com/Kiritogtsa/pi_in_go/internal/domains/entries"
)

type Update interface {
	update(*entries.Trabalhos) error
}
type Insert interface {
	insert(*entries.Trabalhos) error
}
type Persit interface {
	Persit(*entries.Trabalhos) error
	Update
	Insert
}

func (t *trabalhos) insert(tr *entries.Trabalhos) error {
	query := "insert into trabalhos(NOME,DESCRICAO) values(?,?)"
	stmt, err := t.conn.Conndb.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(tr.Nome, tr.Descricao)
	if err != nil {
		return err
	}
	return nil
}

func (t *trabalhos) update(tr *entries.Trabalhos) error {
	query := "update trabalhos set NOME=?,DESCRICAO=?"
	stmt, err := t.conn.Conndb.Prepare(query)
	if err != nil {
		return err
	}
	rows, err := stmt.Exec(tr.Nome, tr.Descricao)
	if err != nil {
		return err
	}
	fmt.Println(rows)
	return nil
}

func (t *trabalhos) Persit(tr *entries.Trabalhos) error {
	if tr.ID == 0 {
		return t.insert(tr)
	} else {
		return t.update(tr)
	}
}
