package entries

import "errors"

type Trabalhos struct {
	ID        int    `gorm:"column:ID" json:"id"`
	Nome      string `gorm:"column:NOME" json:"nome"`
	Descricao string `gorm:"column:DESCRICAO" json:"descricao"`
}

func Newtra(nome string, descricao string) (*Trabalhos, error) {
	if nome != "" && descricao != "" {
		return &Trabalhos{
			Nome:      nome,
			Descricao: descricao,
		}, nil
	}
	return nil, errors.New("erro ao criar o trabalho")
}
