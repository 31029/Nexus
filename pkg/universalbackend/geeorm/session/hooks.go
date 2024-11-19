package session

import (
	"reflect"

	//"github.com/31029/nexus/pkg/universalbackend/geeorm/log"
)

// Hooks constants
const (
	BeforeQuery  = "BeforeQuery"
	AfterQuery   = "AfterQuery"
	BeforeUpdate = "BeforeUpdate"
	AfterUpdate  = "AfterUpdate"
	BeforeDelete = "BeforeDelete"
	AfterDelete  = "AfterDelete"
	BeforeInsert = "BeforeInsert"
	AfterInsert  = "AfterInsert"
)

type IAfterQuery interface {
    AfterQuery(s *Session) error
}

type IBeforeInsert interface {
    BeforeInsert(s *Session) error
}

// CallMethod calls the registered hooks
// 钩子机制同样是通过反射来实现的，s.RefTable().Model 或 value 即当前会话正在操作的对象，使用 MethodByName 方法反射得到该对象的方法。
func (s *Session) CallMethod(method string, value interface{}) {
	param := reflect.ValueOf(value)
    switch method {
    case AfterQuery:
        if i, ok := param.Interface().(IAfterQuery); ok {
            i.AfterQuery(s)
        }
    case BeforeInsert:
        if i, ok := param.Interface().(IBeforeInsert); ok {
            i.BeforeInsert(s)
        }
    default:
        panic("unsupported hook method")
    }
	return
}