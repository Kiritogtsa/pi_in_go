package trabalhos

import "github.com/Kiritogtsa/pi_in_go/internal/domains/entries"

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

func (t *trabalhos) insert(*entries.Trabalhos) error {
	return nil
}

func (t *trabalhos) update(*entries.Trabalhos) error {
	return nil
}
func (t *trabalhos) Persit(tr *entries.Trabalhos) error {
	if tr.ID == 0 {
		return t.insert(tr)
	} else {
		return t.update(tr)
	}
}
