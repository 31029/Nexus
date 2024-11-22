package session

import "github.com/31029/nexus/pkg/universalbackend/geeorm/log"

// 新建文件 session/transaction.go 封装事务的 Begin、Commit 和 Rollback 三个接口。
// 封装的另一个目的是统一打印日志，方便定位问题。
func (s *Session) Begin() (err error) {
	log.Info("transaction begin")
	if s.tx, err = s.db.Begin(); err != nil {
		log.Error(err)
		return
	}
	return
}

func (s *Session) Commit() (err error) {
	log.Info("transaction commit")
	if err = s.tx.Commit(); err != nil {
		log.Error(err)
	}
	return
}

func (s *Session) Rollback() (err error) {
	log.Info("transaction rollback")
	if err = s.tx.Rollback(); err != nil {
		log.Error(err)
	}
	return
}