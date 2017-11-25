package daryl_db

import (
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"reflect"
	"strings"
)

func ToExpr(fields ...string) []interface{} {
	v := make([]interface{}, len(fields))
	for i, f := range fields {
		v[i] = sq.Expr(fmt.Sprintf(":%s", f))
	}
	return v
}

func ListField(s interface{}, access string) ([]string, error) {
	t := reflect.TypeOf(s)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return nil, errors.New("first argument should be struct of pointer to struct")
	}

	fields := make([]string, 0, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		if tg, ok := t.Field(i).Tag.Lookup("db"); ok == true {
			if a, ok := t.Field(i).Tag.Lookup("access"); ok == true {
				if !strings.Contains(a, access) {
					continue
				}
			}
			fields = append(fields, tg)
		}
	}
	return fields, nil
}

func SetModelIntField(s interface{}, field string, value int) error {
	t := reflect.TypeOf(s)

	if t.Kind() != reflect.Ptr {
		return errors.New("SetModelIntField expects a pointer argument")
	}

	ps := reflect.ValueOf(s).Elem()

	f := ps.FieldByName(field)
	if !f.IsValid() || !f.CanSet() || f.Kind() != reflect.Int {
		return errors.New("SetModelIntField Invalid field or cannot set or not Int")
	}
	f.SetInt(int64(value))
	return nil
}
