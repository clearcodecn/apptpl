package mysql

import (
	"github.com/go-xorm/xorm"
)

type MysqlStorage struct {
	engine *xorm.Engine
}

func NewMysqlStorage(engine *xorm.Engine) *MysqlStorage {
	return &MysqlStorage{}
}

func (ms *MysqlStorage) NewSession() *xorm.Session {
	return ms.engine.NewSession()
}
