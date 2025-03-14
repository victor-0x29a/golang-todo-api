package models

import "golang-api/db"

func Insert(toDo ToDo) (id int64, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	var sql string = `INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING id`

	err = conn.QueryRow(sql, toDo.Title, toDo.Description, toDo.Done).Scan(&id)

	return
}
