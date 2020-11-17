package database

import (
	"errors"
)

func Run(query string, params ...interface{}) error {
	db = GetConnection()

	if len(params) > 0 {
		stmt, err := db.Prepare(query)

		if err != nil {
			return err
		}

		defer stmt.Close()

		row, err := stmt.Exec(params...)

		if err != nil {
			return err
		}

		if i, err := row.RowsAffected(); err != nil || i != 1 {
			return errors.New("ERROR: An affected row was expected")
		}
	} else {
		_, err := db.Exec(query)

		if err != nil {
			return err
		}
	}

	return nil
}
