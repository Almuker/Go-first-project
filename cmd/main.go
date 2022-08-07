package main

import (
	"Go_project/db"
	"Go_project/services"
)

func main() {
	// Подключение к БД
	db.ConnectToDataBase()
	// Запросы
	services.HandleRequest()
}