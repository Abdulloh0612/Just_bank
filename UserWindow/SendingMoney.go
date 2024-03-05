package UserWindow

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func SendingMoney(Id int, Password string) {
	db, err := sql.Open("mysql", "abdulloh:member1206@tcp(localhost:3306)/pet_Registration")
	if err != nil {
		log.Fatal("Ошибка при открытии соединения с базой данных:", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT balance FROM Users WHERE id = ? and password = ?", Id, Password)
	if err != nil {
		log.Fatal("Ошибка при выполнении SQL-запроса:", err)
	}
	defer rows.Close()

	var user Person

	for rows.Next() {
		err := rows.Scan(&user.Balance)
		if err != nil {
			log.Fatal("Error scanning result:", err)
		}
	}
	var money int
	var receiver int
	fmt.Println("--------------------------")
	fmt.Println("Sending money/Отправка денег")
	fmt.Printf("Ваш баланс: %d\n", user.Balance)
	fmt.Print("Введите ID получателя: ")
	fmt.Scanln(&receiver)
	fmt.Print("Введите сумму: ")
	fmt.Scanln(&money)
	if money == 0 {
		MainWindow(Id)
	} else if money > user.Balance {
		fmt.Println("Вы ввели больше чем у вас есть!")
		SendingMoney(Id, Password)
	} else {
		rows, err := db.Query("SELECT balance FROM Users WHERE id = ?", receiver)
		if err != nil {
			log.Fatal("Ошибка при выполнении SQL-запроса:", err)
		}
		defer rows.Close()

		var user2 Person

		for rows.Next() {
			err := rows.Scan(&user2.Balance)
			if err != nil {
				log.Fatal("Error scanning result:", err)
			}
		}
		_, err = db.Exec("UPDATE Users SET balance = ? WHERE id = ?", money+user2.Balance, receiver)
		if err != nil {
			log.Fatal("Ошибка при выполнении SQL-запроса:", err)
		}
		money = user.Balance - money
		_, err = db.Exec("UPDATE Users SET balance = ? WHERE id = ?", money, Id)
		if err != nil {
			log.Fatal("Ошибка при выполнении SQL-запроса:", err)
		}

		fmt.Println("Данные успешно изменены.")
		MainWindow(Id)
	}

}
