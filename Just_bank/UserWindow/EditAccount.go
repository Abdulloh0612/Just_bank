package UserWindow

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func Edit(Id int, Password string) {
	db, err := sql.Open("mysql", "abdulloh:member1206@tcp(localhost:3306)/pet_Registration")
	if err != nil {
		log.Fatal("Ошибка при открытии соединения с базой данных:", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, surname, address, password FROM Users WHERE id = ?", Id)
	if err != nil {
		log.Fatal("Ошибка при выполнении SQL-запроса:", err)
	}
	defer rows.Close()

	var user Person

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.Surname, &user.Address, &user.Password)
		if err != nil {
			log.Fatal("Error scanning result:", err)
		}
	}
	var parameter, value string
	fmt.Println("--------------------------")
	fmt.Println("Edit Account")
	fmt.Println("")
	fmt.Printf("Your name:        %s\n", user.Name)
	fmt.Printf("Your surname:     %s\n", user.Surname)
	fmt.Printf("Your address: %s\n", user.Address)
	fmt.Printf("Your password:    %s\n", user.Password)
	fmt.Println("")
	fmt.Print("Введите что вы хотите изменить: ")
	fmt.Scanln(&parameter)
	parameter = strings.TrimSpace(strings.ToLower(parameter))
	if parameter == "0" {
		MainWindow(Id)
	} else if parameter == "id" || parameter == "name" || parameter == "surname" || parameter == "address" || parameter == "password" {
		fmt.Print("Введите данные для этого параметра: ")
		fmt.Scanln(&value)
		query := fmt.Sprintf("UPDATE Users SET %s = ? WHERE id = ?", parameter)
		_, err = db.Exec(query, value, Id)
		if err != nil {
			log.Fatal("Ошибка при выполнении SQL-запроса:", err)
		}

		fmt.Println("Данные успешно изменены.")
	} else {
		fmt.Println("Посмотрите выше и напишите правельно!")
		Edit(Id, Password)
	}
	MainWindow(Id)
}
