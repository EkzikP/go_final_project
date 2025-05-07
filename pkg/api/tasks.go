package api

import (
	"go1f/pkg/db"
	"net/http"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, _ *http.Request) {
	tasks, err := db.Tasks(50) // в параметре максимальное количество записей
	if err != nil {
		output := "ошибка запроса к базе данных"
		writeJson(w, Out{Error: output})
		return
	}
	writeJson(w, TasksResp{Tasks: tasks})
}
