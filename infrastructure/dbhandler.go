package infrastructure

import (
	"github.com/go-xorm/xorm"
)

//外部接続するハンドラを生成
//外部パッケージ（xorm）を利用するので、structでパッケージ化

//DBHandler handler
type DBHandler struct {
	connection *xorm.Engine
}

//NewHandler create handler
func NewHandler() *DBHandler {
	const d = "mysql"
	const sn = "test:test@tcp(db:3306)/test"

	engine, err := xorm.NewEngine(d, sn)
	if err != nil {
		panic(err.Error)
	}

	handler := new(DBHandler)
	handler.connection = engine
	return handler
}
