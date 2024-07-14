package conf

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql" // Importar o driver MySQL
	"github.com/gorilla/sessions"
)

const (
	maxWaitTime  = 60 * time.Second
	pingInterval = 5 * time.Second
)

var Store = sessions.NewCookieStore([]byte("top-secret"))

// Estruturas para a configuraçã o
type db struct {
	User   string `json:"user"`
	Port   string `json:"port"`
	Senha  string `json:"passowd"`
	Host   string `json:"host"`
	Dbname string `json:"db_name"`
}

type Confjson struct {
	Db  db `json:"db"`
	App struct {
		Port string `json:"port"`
	} `json:"app"`
}

// Função para ler a configuração do JSON
func NewConf() (*Confjson, error) {
	// Determina o caminho absoluto para conf.json
	// Abre o arquivo conf.json
	file, err := os.Open("config.json")
	if err != nil {
		return nil, fmt.Errorf("não existe o arquivo conf.json: %w", err)
	}
	defer file.Close()

	var conf Confjson
	if err := json.NewDecoder(file).Decode(&conf); err != nil {
		return nil, fmt.Errorf("erro ao decodificar conf.json: %w", err)
	}
	fmt.Println(conf)

	// Valida as configurações
	if conf.Db.User == "" || conf.Db.Host == "" || conf.Db.Dbname == "" || conf.Db.Port == "" {
		return nil, errors.New("configurações de banco de dados inválidas")
	}

	return &conf, nil
}

// Função para criar uma nova conexão
func (c *Confjson) Newconn() (*Db, error) {
	return newConn(c.Db)
}

// Estrutura para o banco de dados
type Db struct {
	Conndb *sql.DB
}

// Função para criar uma nova conexão de banco de dados
func newConn(d db) (*Db, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		d.User, d.Senha, d.Host, d.Port, d.Dbname)
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao banco de dados: %w", err)
	}

	// Verifica a conexão com waiterdb
	if err := waiterdb(conn); err != nil {
		return nil, err
	}

	return &Db{
		Conndb: conn,
	}, nil
}

// Função para aguardar o banco de dados estar disponível
func waiterdb(db *sql.DB) error {
	deadline := time.Now().Add(maxWaitTime)
	for {
		err := db.Ping()
		if err == nil {
			log.Println("Banco de dados conectado com sucesso.")
			return nil
		}
		if time.Now().After(deadline) {
			return fmt.Errorf("tempo de conexão expirou: %w", err)
		}
		log.Printf("Aguardando pelo banco de dados... %v", err)
		time.Sleep(pingInterval)
	}
}
