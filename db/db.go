package db

import (
	"Go_project/models"
	"Go_project/services"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDataBase() {
	db, err := sql.Open("mysql", "Almuker:Almuker@tcp(127.0.0.1:3306)/progsus")
	services.ErrorsHandler(err)
	
	defer db.Close()
	fmt.Println("Подключение к БД")


	// Данные по users
	res , err := db.Query("SELECT * FROM `users`")
	services.ErrorsHandler(err)

	for res.Next() {
		var user models.User
		err = res.Scan(&user.Id, &user.Name, &user.Age)
		services.ErrorsHandler(err)
		fmt.Println("User " + user.Name)
	}
}

// Установка данных
// insert, err := db.Query("INSERT INTO `users` (`name`, `age`) VALUES ('Arman', 28)")
// if err != nil {
// 	fmt.Println(err)
// 	panic(err)
// }

// defer insert.Close()