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
			return errors.New("An affected row was expected")
		}
	} else {
		_, err := db.Exec(query)

		if err != nil {
			return err
		}
	}

	return nil
}

/* Only 1 result */
/* thing, err := database.Get("SELECT thing FROM dbthings WHERE thingId = ? ", 123456) */
func Get(query string, key ...interface{}) (data interface{}, err error) {
	db = GetConnection()

	err = db.QueryRow(query, key...).Scan(&data)

	if err != nil {
		return nil, err
	}

	return data, nil
}
