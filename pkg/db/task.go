package db

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func AddTask(task *Task) (int64, error) {
	var id int64
	// определите запрос
	query := `INSERT INTO scheduler (date, title, comment, repeat) VALUES (?, ?, ?, ?)`
	res, err := DB.Exec(query, task.Date, task.Title, task.Comment, task.Repeat)
	if err == nil {
		id, err = res.LastInsertId()
	}
	return id, err
}

func Tasks(limit int) ([]*Task, error) {
	query := `SELECT * FROM scheduler ORDER BY date LIMIT ?`
	rows, err := DB.Query(query, limit)
	if err != nil {
		return []*Task{}, err
	}
	defer rows.Close()

	tasks := []*Task{}
	for rows.Next() {
		t := &Task{}
		err := rows.Scan(&t.ID, &t.Date, &t.Title, &t.Comment, &t.Repeat)
		if err != nil {
			return []*Task{}, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}
