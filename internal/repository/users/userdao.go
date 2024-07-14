package users

import (
	"github.com/Kiritogtsa/pi_in_go/conf"
)

type userre struct {
	conn conf.Db
}

func Nuserre() Userif {
	conf, err := conf.NewConf()
	if err != nil {
		return nil
	}
	conn, err := conf.Newconn()
	if err != nil {
		return nil
	}
	return &userre{conn: *conn}
}

// metodos de insert e update privados fora do pacote, o metodo persert publico

// metodos que o precisa ser feito para a tabela usuario
