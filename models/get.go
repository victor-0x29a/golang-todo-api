package models

import "golang-api/db"

func Get(id int64) (toDo ToDo, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM todos WHERE id = $1`, id)

	err = row.Scan(
		&toDo.ID,
		&toDo.Title,
		&toDo.Description,
		&toDo.Done)

	return
}
