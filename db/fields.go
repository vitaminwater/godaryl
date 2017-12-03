package daryl_db

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	sq "github.com/Masterminds/squirrel"
)

func ToExpr(fields ...string) []interface{} {
	v := make([]interface{}, len(fields))
	for i, f := range fields {
		v[i] = sq.Expr(fmt.Sprintf(":%s", f))
	}
	return v
}

func ListField(s interface{}, access string) ([]string, []string, error) {
	t := reflect.TypeOf(s)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return nil, nil, errors.New("first argument should be struct or pointer to struct")
	}

	fields := make([]string, 0, t.NumField())
	dbFields := make([]string, 0, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		if dbTag, ok := t.Field(i).Tag.Lookup("db"); ok == true {
			if dbTag == "-" {
				continue
			}
			if accessTag, ok := t.Field(i).Tag.Lookup("access"); ok == true {
				if !strings.Contains(accessTag, access) {
					continue
				}
			}
			dbFields = append(dbFields, dbTag)
			/*if _, ok := t.Field(i).Tag.Lookup("jsonb"); ok == true {
				fields = append(fields, fmt.Sprintf("%s::jsonb", dbTag))
			} else {*/
			fields = append(fields, dbTag)
			//}
		}
	}
	return fields, dbFields, nil
}

func SetModelStringField(s interface{}, field, value string) error {
	t := reflect.TypeOf(s)

	if t.Kind() != reflect.Ptr {
		return errors.New("SetModelIntField expects a pointer argument")
	}

	ps := reflect.ValueOf(s).Elem()

	f := ps.FieldByName(field)
	if !f.IsValid() || !f.CanSet() || f.Kind() != reflect.String {
		return errors.New("SetModelIntField Invalid field or cannot set or not Int")
	}
	f.SetString(value)
	return nil
}
