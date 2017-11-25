package daryl_db

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
	log "github.com/sirupsen/logrus"
)

func InsertQuery(t string, s interface{}) (string, error) {
	fields, _ := ListField(s, "i")
	qi, _, err := sq.Insert(t).Columns(fields...).Values(ToExpr(fields...)...).Suffix("RETURNING id").ToSql()
	if err != nil {
		return "", err
	}
	log.Info(qi)
	return qi, nil
}

func UpdateQuery(t, idf string, s interface{}) (string, error) {
	fields, _ := ListField(s, "u")
	ub := sq.Update(t)
	for _, field := range fields {
		ub = ub.Set(field, sq.Expr(fmt.Sprintf(":%s", field)))
	}
	qu, _, err := ub.Where(fmt.Sprintf("%[1]s = :%[1]s", idf)).ToSql()
	if err != nil {
		return "", err
	}
	log.Info(qu)
	return qu, nil
}
