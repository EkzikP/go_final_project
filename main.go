package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"go1f/pkg/db"
	"go1f/pkg/server"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Не найден файл с переменными окружения")
	}
}

func main() {
	// Инициализация БД
	TODO_DBFILE, exists := os.LookupEnv("TODO_DBFILE")
	if !exists {
		TODO_DBFILE = "scheduler.db"
	}
	_, err := initDb(TODO_DBFILE)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Запуск WEB сервера
	TODO_PORT, exists := os.LookupEnv("TODO_PORT")
	if !exists {
		TODO_PORT = "7540"
	}
	str := startServer(TODO_PORT)
	if str != "" {
		fmt.Println(str)
	}

}

func startServer(port string) string {
	err := server.StartServer(port)
	if err != nil {
		return fmt.Sprintf("Ошибка при запуске сервера: %s", err.Error())
	}
	return ""
}

func initDb(path string) (*db.TasksStore, error) {
	store, err := sql.Open("sqlite", path)
	if err != nil {
		return &db.TasksStore{}, err
	}
	defer store.Close()

	base := db.New(store)
	if err := base.Initialize(); err != nil {
		return &db.TasksStore{}, err
	}
	return base, nil
}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	date, err := time.Parse("20060102", dstart)
	if err != nil {
		return "", errors.New("Неправильный формат даты")
	}

	sliceRepeat, err := split(repeat)
	if err != nil {
		return "", err
	}

	switch sliceRepeat[0] {
	case "d":
		if len(sliceRepeat) != 2 {
			return "", errors.New("не указан интервал в днях")
		}

		interval, err := strconv.Atoi(sliceRepeat[1])
		if err != nil {
			return "", errors.New("не верный интервал")
		}

		if interval < 1 || interval > 400 {
			return "", errors.New("превышен максимально допустимый интервал")
		}

		for {
			date = date.AddDate(0, 0, interval)
			if afterNow(date, now) {
				break
			}
		}
		return date.Format("20060102"), nil
	case "y":
		for {
			date = date.AddDate(1, 0, 0)
			if afterNow(date, now) {
				break
			}
		}
		return date.Format("20060102"), nil
	case "w":

	case "m":

	default:
		return "", errors.New("недопустимый символ")
	}

}

func afterNow(date time.Time, now time.Time) bool {
	return date.After(now)
}

func split(repeat string) ([]string, error) {
	if len(repeat) < 1 {
		return []string{}, errors.New("неверный формат правила повторения")
	}
	slice := strings.Split(repeat, " ")
	return slice, nil
}
