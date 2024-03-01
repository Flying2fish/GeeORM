package dialect

import "reflect"

var dialectsMap = map[string]Dialect{}

type Dialect interface { //给不同数据库实现dialect接口
	DataTypeOf(typ reflect.Value) string
	TableExistSQL(tableName string) (string, []interface{})
}

func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

// 通过driver来获取对应数据库的dialect
func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return
}
