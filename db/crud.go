package daryl_db

import (
	"database/sql"
)

func Insert(t string, s interface{}) error {
	q, err := InsertQuery(t, s)
	if err != nil {
		return err
	}
	if stmt, err := db.PrepareNamed(q); err != nil {
		return err
	} else {
		var id string
		if err := stmt.Get(&id, s); err != nil {
			return err
		}
		return SetModelStringField(s, "Id", id)
	}
}

func Update(t, idf string, s interface{}) (sql.Result, error) {
	q, err := UpdateQuery(t, idf, s)
	if err != nil {
		return nil, err
	}
	return db.NamedExec(q, s)
}
