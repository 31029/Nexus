package geeorm

import (
	"database/sql"

	"github.com/31029/nexus/pkg/universalbackend/geeorm/dialect"
	"github.com/31029/nexus/pkg/universalbackend/geeorm/log"
	"github.com/31029/nexus/pkg/universalbackend/geeorm/session"
)

// Session 负责与数据库的交互
// 而交互前的准备工作（比如连接/测试数据库），交互后的收尾工作（关闭连接）等就交给 Engine 来负责了。
// Engine 是 GeeORM 与用户交互的入口
type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}
	// Send a ping to make sure the database connection is alive.
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	// make sure the specific dialect exists
	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("dialect %s Not Found", driver)
		return
	}
	e = &Engine{db: db, dialect: dial}
	log.Info("Connect database success")
	return
}

func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db, engine.dialect)
}

func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error("Failed to close database")
	}
	log.Info("Close database success")
}