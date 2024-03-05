package Log_in

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Person struct {
	ID       int
	Name     string
	Surname  string
	Address  string
	Password string
	Balance  int
}

func Log_in() int {
	fmt.Println("--------------------------")
	fmt.Println("Log_in ")
	fmt.Println("0. Exist ")
	fmt.Println("")

	var Id, Password string
	var entering int

	// * The form
	fmt.Print("Введите ID: ")
	fmt.Scanln(&Id)
	if Id == "0" {
		os.Exit(0)
	} else {
		fmt.Print("Введите Пароль: ")
		fmt.Scanln(&Password)
		if Password == "0" {
			os.Exit(0)
		} else if len(Password) <= 7 {
			fmt.Println("Пароль должен содержать не менее 8 символов")
			Log_in()
		} else {
			entering = Check(strings.TrimSpace(Id), strings.TrimSpace(Password))
			return entering
		}
	}
	return entering
}

func Check(id, password string) int {
	db, err := sql.Open("mysql", "abdulloh:member1206@tcp(localhost:3306)/pet_Registration")
	if err != nil {
		log.Fatal("Ошибка при открытии соединения с базой данных:", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, surname, address, password, balance FROM Users WHERE id = ? AND password = ?", id, password)
	if err != nil {
		log.Fatal("Ошибка при выполнении SQL-запроса:", err)
	}
	defer rows.Close()

	var users []Person
	var user Person

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.Surname, &user.Address, &user.Password, &user.Balance)
		if err != nil {
			log.Fatal("Ошибка при сканировании результата:", err)
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		fmt.Println("Вы ввели неправельны пароль!")
		fmt.Println("")
		Log_in()
	}
	return user.ID
}
