package daryl_db

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	log "github.com/sirupsen/logrus"
)

func GetQuery(t, match string, s interface{}) (string, error) {
	_, dbFields, _ := ListField(s, "s")
	q := sq.Select(dbFields...).From(t)
	if match != "" {
		q = q.Where(fmt.Sprintf("%[1]s = :%[1]s", match))
	}
	qg, _, err := q.ToSql()
	if err != nil {
		return "", err
	}
	return qg, nil
}

func GettPaginatedQuery(t, match string, s interface{}, from, to int32) (string, error) {
	_, dbFields, _ := ListField(s, "s")
	q := sq.Select(dbFields...).From(t)
	if match != "" {
		q = q.Where(fmt.Sprintf("%[1]s = :%[1]s", match))
	}
	q = q.Offset(uint64(from)).Limit(uint64(to - from))
	qg, _, err := q.ToSql()
	if err != nil {
		return "", err
	}
	return qg, nil
}

func InsertQuery(t string, s interface{}) (string, error) {
	fields, dbFields, _ := ListField(s, "i")
	qi, _, err := sq.Insert(t).Columns(dbFields...).Values(ToExpr(fields...)...).Suffix("RETURNING id").ToSql()
	if err != nil {
		return "", err
	}
	log.Info(qi)
	return qi, nil
}

func UpdateQuery(t, idf string, s interface{}) (string, error) {
	fields, dbFields, _ := ListField(s, "u")
	ub := sq.Update(t)
	for i, field := range fields {
		ub = ub.Set(dbFields[i], sq.Expr(fmt.Sprintf(":%s", field)))
	}
	qu, _, err := ub.Where(fmt.Sprintf("%[1]s = :%[1]s", idf)).ToSql()
	if err != nil {
		return "", err
	}
	log.Info(qu)
	return qu, nil
}
