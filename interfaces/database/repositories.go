package database

//interfaces for infrastructure

type DbHandler interface {
	//Excute(statement string)
	Excute(statement string) (Result, error)
	Query(statement string) (Row, error)
}

type Row interface {
	Scan(d ...interface{}) error
	Next() bool
	Close() error
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}
