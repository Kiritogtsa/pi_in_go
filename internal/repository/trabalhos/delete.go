package trabalhos

type Delete interface {
	Delete(int) error
}

func (tr *trabalhos) Delete(int) error {
	return nil
}
