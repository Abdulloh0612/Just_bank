package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name"`
	Age  uint16 `json:"age"`
}

func main() {
	// Установка параметров подключения
	connectionString := "abdulloh:member1206@tcp(127.0.0.1:3306)/golang"

	// Открытие соединения с базой данных
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// insert, err := db.Query("insert into users (name,age) values('Bob',35)")
	// if err != nil {
	// 	panic(err)
	// }
	// defer insert.Close()
	// fmt.Println("Класссс")

	res, _ := db.Query("select name, age from users")

	for res.Next() {
		var user User
		err := res.Scan(&user.Name, &user.Age)
		if err != nil {
			panic(err)
		}

		fmt.Printf("User: %s with age %d \n", user.Name, user.Age)
	}

}
