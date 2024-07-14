package trabalhos

import "github.com/Kiritogtsa/pi_in_go/internal/domains/entries"

type Gets interface {
	Getall() ([]entries.Trabalhos, error)
}

func (t *trabalhos) Getall() ([]entries.Trabalhos, error) {
	query := "select * from trabalhos"
	rows, err := t.conn.Conndb.Query(query)
	if err != nil {
		return nil, err
	}
	trabalhos := make([]entries.Trabalhos, 1)

	for rows.Next() {
		var trabalho entries.Trabalhos
		err = rows.Scan(&trabalho.ID, &trabalho.Nome, &trabalho.Descricao)
		if err != nil {
			continue
		} else {
			trabalhos = append(trabalhos, trabalho)
		}
	}
	return trabalhos, nil
}
