package persistence

import (
	"sync"
	"fmt"
	"ameicosmeticos/app/contracts"
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlHandler struct {
	Conn *sql.DB
}

var (
	mysqlHandler	*MysqlHandler
	dbOnce			sync.Once
)

func NewDBHandler(config DBConnectionConfig) *MysqlHandler {
	if mysqlHandler == nil {
		dbOnce.Do(func() {
			println("> [ infrastructure ] Creating database handler ...")
			connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.Database)
			sqlConn, _ := sql.Open("mysql", connectionString)

			if sqlConn.Ping() != nil {
				log.SetFlags(0)
				log.Fatal("> [ infrastructure ] Não foi possivel estabelecer conexão com o BD")
			}

			sqlConn.SetMaxOpenConns(config.MaxOpenConns)
			sqlConn.SetMaxIdleConns(config.MaxIdleConns)
			sqlConn.SetConnMaxLifetime(config.ConnMaxLifetime)
			mysqlHandler = &MysqlHandler{sqlConn}
			println("> [ infrastructure ] Database handler created")
		})
	}

	return mysqlHandler
}

func (this *MysqlHandler) Execute(statement string) {
	this.Conn.Exec(statement)
}

func (this *MysqlHandler) Query(statement string) (contracts.IRow, error) {	
	rows, err := this.Conn.Query(statement)

	if err != nil {
		fmt.Println(err)
		return new(MysqlRow), err
	}

	row := new(MysqlRow)
	row.Rows = rows
	return row, nil
}

type MysqlRow struct {
	Rows *sql.Rows
}

func (r MysqlRow) Scan(dest ...interface{}) error {
	err := r.Rows.Scan(dest...)

	if err != nil {
		return err
	}

	return  nil
}

func (r MysqlRow) Next() bool {
	return r.Rows.Next()
}

func (r MysqlRow) Close() error {
	return r.Rows.Close()
}