package infrastructure

import (
	"database/sql"

	"github.com/wakashiyo/charisha_dev-server/interfaces/database"
)

//DBHandler handler
type DBHandler struct {
	conn *sql.DB
}

//NewDbHandler create handler
func NewDbHandler() database.DbHandler {
	const d = "mysql"
	const sn = "test:test@tcp(db:3306)/test"

	db, err := sql.Open(d, sn)
	if err != nil {
		panic(err.Error)
	}

	handler := new(DBHandler)
	handler.conn = db
	return handler
}

func (db *DBHandler) Excute(statement string) (database.Result, error) {
	res := SqlResult{}
	result, err := db.conn.Exec(statement)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, err
}

func (db *DBHandler) Query(statment string) (database.Row, error) {
	rows, err := db.conn.Query(statment)
	if err != nil {
		return rows, err
	}
	row := new(SqlRow)
	row.Rows = rows
	return row, err
}

type SqlResult struct {
	Result sql.Result
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

type SqlRow struct {
	Rows *sql.Rows
}

func (r SqlRow) Scan(d ...interface{}) error {
	return r.Rows.Scan(d...)
}

func (r SqlRow) Next() bool {
	return r.Rows.Next()
}

func (r SqlRow) Close() error {
	return r.Rows.Close()
}
