package entries

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

type User struct {
	ID       int    `gorm:"column:ID" json:"id"`
	Name     string `gorm:"column:NOME" json:"name"`
	Cpf      string `gorm:"column:CPF" json:"cpf"`
	Email    string `gorm:"column:EMAIL" json:"email"`
	Senha    string `gorm:"column:SENHA" json:"senha"`
	DataNas  string `gorm:"column:DATA_NASCIMENTO" json:"data_nas"`
	DataAd   string `gorm:"column:DATA_ADMISSAO" json:"data_ad"`
	Telefone int64  `gorm:"column:TELEFONE" json:"telefone"`
	Sexo     string `gorm:"column:SEXO" json:"sexo"`
	Grupo    string `gorm:"column:GRUPO" json:"grupo"`
	DeleteAt string `gorm:"column:DELETE_AT" json:"delete_at"`
	TrID     int    `gorm:"column:TR_ID" json:"tr_id"`
}

func Newuser(nome string, cpf string, email string, senha string, data_nas string, data_ad string, telefone int64, sexo string, grupo string) (*User, error) {
	if Validarcpf(cpf) != true {
		return nil, errors.New("nao foi possivel criar um cpf")
	}

	return &User{
		Name:    nome,
		Cpf:     cpf,
		Email:   email,
		Senha:   senha,
		DataNas: data_nas,
		DataAd:  data_ad,
		Sexo:    sexo,
		Grupo:   grupo,
	}, nil
}
func Validarcpf(cpf string) bool {
	cpf = strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		}
		return -1
	}, cpf)
	if len(cpf) != 11 {
		return false
	}
	if alldigitsequal(cpf) {
		return false
	}

	if !validarDigits(cpf[:9], cpf[9]) {
		return false
	}
	if !validarDigits(cpf[:10], cpf[10]) {
		return false
	}
	return true
}

func alldigitsequal(cpf string) bool {
	for i := 0; i < len(cpf); i++ {
		if cpf[i] != cpf[i-1] && i > 0 {
			return false
		}
	}
	return true
}

func validarDigits(base string, digit byte) bool {
	sum := 0
	for i := 0; i < len(base); i++ {
		num, _ := strconv.Atoi(string(base[i]))
		sum += num * (len(base) + 1 - i)
	}
	result := sum % 11
	var expected byte

	if result < 2 {
		expected = '0'
	} else {
		expected = byte(11 - result + '0')
	}
	return expected == digit
}
