package users

type Delete interface {
	Delete(id int) error
}

func (u *userre) Delete(id int) error {
	return nil
}
