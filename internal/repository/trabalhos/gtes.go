package trabalhos

import "github.com/Kiritogtsa/pi_in_go/internal/domains/entries"

type Gets interface {
	Getall() ([]entries.Trabalhos, error)
}

func (t *trabalhos) Getall() ([]entries.Trabalhos, error) {
	return nil, nil
}
