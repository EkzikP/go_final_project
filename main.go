package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"go1f/pkg/db"
	"go1f/pkg/server"
	"log"
	"os"
)

func init() {
	// Загрузка переменных окружения
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

	DB, err := sql.Open("sqlite", TODO_DBFILE)
	if err != nil {
		fmt.Println(err)
		return
	}
	db.DB = DB
	defer DB.Close()

	err = db.Init(TODO_DBFILE)
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
