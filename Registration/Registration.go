package Registration

import (
	"Go_Projects/Just_bank/Log_in"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Person2 struct {
	ID       int
	Name     string
	Surname  string
	Address  string
	Password string
	Balance  int
}

type Person struct {
	Name     string
	Surname  string
	Address  string
	Password string
}

func FirstWindow() int {
	var choice int
	var entry int
	fmt.Println("--------------------------")
	fmt.Println("1. Log in")
	fmt.Println("2. Register")
	fmt.Println("0. Exit")
	fmt.Println("--------------------------")
	fmt.Print("Your choice is ")
	fmt.Scan(&choice)
	switch choice {
	case 0:
		os.Exit(0)
	case 1:
		entry = Log_in.Log_in()
	case 2:
		entry = Register()
	default:
		fmt.Println("Такой команды нету!")
		FirstWindow()
	}
	return entry
}

func Register() int {
	var newUser Person
	fmt.Println("--------------------------")
	fmt.Println("Registration")
	fmt.Println("")
	fmt.Println("0. Exist")
	fmt.Println("")

	var Name, Surname, Address, Password string

	fmt.Print("Введите Имя: ")
	fmt.Scanln(&Name)
	if Name == "0" {
		os.Exit(1)
	} else if len(Name) < 1 {
		fmt.Println("Ошибка ввода имени")
		Register()
	}
	newUser.Name = Name

	fmt.Print("Введите Фамилию: ")
	fmt.Scanln(&Surname)
	if Surname == "0" {
		os.Exit(1)
	} else if len(Surname) < 1 {
		fmt.Println("Ошибка ввода фамилии")
		Register()
	}
	newUser.Surname = Surname

	fmt.Print("Введите свой номер или почту: ")
	fmt.Scanln(&Address)
	if Address == "0" {
		os.Exit(1)
	} else if len(Address) < 1 {
		fmt.Println("Ошибка ввода адреса")
		Register()
	}
	newUser.Address = Address

	fmt.Print("Введите Пароль: ")
	fmt.Scanln(&Password)
	if Password == "0" {
		os.Exit(1)
	} else if len(Password) <= 7 {
		fmt.Println("Пароль должен содержать не менее 8 символов")
		Register()
	}
	newUser.Password = Password

	entry := user_registration(newUser.Name, newUser.Surname, newUser.Address, newUser.Password)
	return entry
}

func user_registration(name, surname, number, password string) int {
	db, err := sql.Open("mysql", "abdulloh:member1206@tcp(localhost:3306)/pet_Registration")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM Users WHERE name = ? AND surname = ?", name, surname).Scan(&count)
	if err != nil {
		panic(err)
	}

	if count == 0 {
		_, err = db.Exec("INSERT INTO Users (name, surname, address, password) VALUES (?, ?, ?, ?)", name, surname, number, password)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("Пользователь уже существует")
		Register()
	}
	entry := Check(name, surname, number, password)
	return entry
}
func Check(name, surname, address, password string) int {
	db, err := sql.Open("mysql", "abdulloh:member1206@tcp(localhost:3306)/pet_Registration")
	if err != nil {
		log.Fatal("Ошибка при открытии соединения с базой данных:", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, surname, address, password, balance FROM Users WHERE name = ? and surname = ? and address = ? and password = ?", name, surname, address, password)
	if err != nil {
		log.Fatal("Ошибка при выполнении SQL-запроса:", err)
	}
	defer rows.Close()

	var users []Person2
	var user Person2

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.Surname, &user.Address, &user.Password, &user.Balance)
		if err != nil {
			log.Fatal("Ошибка при сканировании результата:", err)
		}
		users = append(users, user)
	}
	return user.ID
}
