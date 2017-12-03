package daryl_db

func Get(t, match string, dest, src interface{}) error {
	q, err := GetQuery(t, match, dest)
	if err != nil {
		return err
	}
	if stmt, err := db.PrepareNamed(q); err != nil {
		return err
	} else {
		if err := stmt.Get(dest, src); err != nil {
			return err
		}
		return nil
	}
}

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

func Update(t, idf string, s interface{}) error {
	q, err := UpdateQuery(t, idf, s)
	if err != nil {
		return err
	}
	_, err = db.NamedExec(q, s)
	if err != nil {
		return err
	}
	return nil
}
