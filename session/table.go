package session

import (
	"GeeORM/log"
	"GeeORM/schema"
	"fmt"
	"reflect"
	"strings"
)

// 通过结构体返回含有结构体表结构的Session
func (s *Session) Model(value interface{}) *Session {
	if s.refTable == nil || reflect.TypeOf(value) != reflect.TypeOf(s.refTable.Model) {
		s.refTable = schema.Parse(value, s.dialect)
	}
	return s
}

// 获取Schema管理的表的结构
func (s *Session) RefTable() *schema.Schema {
	if s.refTable == nil {
		log.Error("Model is not set")
	}
	return s.refTable
}

func (s *Session) CreateTable() error {
	table := s.RefTable()
	var columns []string
	for _, field := range table.Fields {
		columns = append(columns, fmt.Sprintf("%s %s %s", field.Name, field.Type, field.Tag))
	}
	desc := strings.Join(columns, ",")
	_, err := s.Raw(fmt.Sprintf("CREATE TABLE %s (%s);", table.Name, desc)).Exec()
	return err
}

func (s *Session) DropTable() error {
	_, err := s.Raw(fmt.Sprintf("DROP TABLE IF EXISTS %s", s.RefTable().Name)).Exec()
	return err
}

func (s *Session) HasTable() bool {
	sql, values := s.dialect.TableExistSQL(s.RefTable().Name)
	row := s.Raw(sql, values...).QueryRow()
	var tmp string
	_ = row.Scan(&tmp)
	return tmp == s.RefTable().Name
}
