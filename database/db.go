package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// ORM serve como ponte de comunicação entre o banco de dados e a aplicação,sem exigir a linguagem do banco
	// Gorm é um ORM que facilita e acelera o desenvolvimento de software com Go.
	// Ele auxilia na realização de todas as operações do CRUD (criar, ler, atualizar, excluir registros)
	// e permite diferentes tipos de associações entre os modelos.
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	StringConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(StringConexao))
	if err != nil {
		log.Panic("Erro ao Conectar no banco!!")
	}
}
