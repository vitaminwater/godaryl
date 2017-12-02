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

func ListField(s interface{}, prefix, access string) ([]string, error) {
	t := reflect.TypeOf(s)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return nil, errors.New("first argument should be struct or pointer to struct")
	}

	fields := make([]string, 0, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		if tg, ok := t.Field(i).Tag.Lookup("db"); ok == true {
			if tg == "-" {
				continue
			}
			abs := tg
			if prefix != "" {
				abs = fmt.Sprintf("%s.%s", prefix, tg)
			}
			if a, ok := t.Field(i).Tag.Lookup("access"); ok == true {
				if !strings.Contains(a, access) {
					continue
				}
			}
			if t.Field(i).Type.Kind() == reflect.Struct || (t.Field(i).Type.Kind() == reflect.Ptr && t.Field(i).Type.Elem().Kind() == reflect.Struct) {
				v := reflect.ValueOf(s)
				if reflect.TypeOf(s).Kind() == reflect.Ptr {
					v = v.Elem()
				}
				f, err := ListField(v.Field(i).Interface(), tg, access)
				if err != nil {
					return fields, err
				}
				if len(f) == 0 {
					fields = append(fields, tg)
				} else {
					fields = append(fields, f...)
				}
			} else {
				fields = append(fields, abs)
			}
		}
	}
	return unduplicate(fields), nil
}

func unduplicate(fs []string) []string {
	dbNames := DBNames(fs)
	added := map[string]string{}
	res := []string{}
	for i, f := range fs {
		d := dbNames[i]
		_, ok := added[d]
		if ok == false {
			res = append(res, f)
			added[d] = f
		}
	}
	return res
}

func DBNames(fs []string) []string {
	res := []string{}
	for _, f := range fs {
		if strings.Contains(f, ".") {
			sp := strings.Split(f, ".")
			res = append(res, sp[len(sp)-1])
		} else {
			res = append(res, f)
		}
	}
	return res
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
