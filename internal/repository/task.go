package repository

import (
	"database/sql"

	"main/internal/model"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) GetAll() ([]model.Task, error) {
	rows, err := r.db.Query("SELECT * FROM task")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.Id, &task.UserID, &task.Title)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *TaskRepository) Create(task model.Task) error {
	_, err := r.db.Exec("INSERT INTO task (id, userId, title) VALUES (?, ?, ?)", task.Id, task.UserID, task.Title)
	return err
}

func (r *TaskRepository) Update(task model.Task) error {
	_, err := r.db.Exec("UPDATE task SET title = ? WHERE id = ?", task.Title, task.Id)
	return err
}

func (r *TaskRepository) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM task WHERE id = ?", id)
	return err
}