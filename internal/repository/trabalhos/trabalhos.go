package trabalhos

import (
	"github.com/Kiritogtsa/pi_in_go/conf"
)

// criar arquivos para cada metodos e interfaces qu eu precise, 1 arquivo\3 metodos\1 interfaces select, 1 arquivo\3 metodos\1 interfaces  save,  1 arquivo\1 metodo/1 interface delete

type trabalhos struct {
	conn conf.Db
}

func Newtradb() Trit {
	conf, err := conf.NewConf()
	if err != nil {
		return nil
	}
	conn, err := conf.Newconn()
	if err != nil {
		return nil
	}
	return &trabalhos{conn: *conn}
}

// metodos base do crud, insert, update e persit

// metodos do crud mais chatos
