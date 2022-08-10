package models

import (
	"database/sql"
)

type Task struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Pic   string `json:"pic"`
	Deadline   string `json:"deadline"`
	Status int    `json:"status"`
}

type TaskCollection struct {
	Tasks []Task `json:"items"`
}

func GetTasks(db *sql.DB) TaskCollection {
	sql := "SELECT id, name, pic, deadline, status FROM tasks"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	result := TaskCollection{}
	for rows.Next() {
		task := Task{}
		err2 := rows.Scan(&task.ID, &task.Name, &task.Pic, &task.Deadline, &task.Status)
		if err2 != nil {
			panic(err2)
		}
		result.Tasks = append(result.Tasks, task)
	}
	return result
}

func PutTask(db *sql.DB, name string, pic string, deadline string, status int) (int64, error) {
	sql := "INSERT INTO tasks(name, pic, deadline, status) VALUES(?,?,?,?)"

	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	result, err2 := stmt.Exec(name, pic, deadline, status)
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

func EditTask(db *sql.DB, taskId int, name string, pic string, deadline string, status int) (int64, error) {
	sql := "UPDATE tasks set name = ?, pic = ?, deadline = ?, status = ? WHERE id = ?"

	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	result, err2 := stmt.Exec(name, pic, deadline, status, taskId)

	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}

func DeleteTask(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM tasks WHERE id = ?"

	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	result, err2 := stmt.Exec(id)
	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}
