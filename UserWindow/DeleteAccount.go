package UserWindow

import (
	"Go_Projects/Just_bank/Registration"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func DeleteAccount(Id int, Password string) {
	db, err := sql.Open("mysql", "abdulloh:member1206@tcp(localhost:3306)/pet_Registration")
	if err != nil {
		log.Fatal("Ошибка при открытии соединения с базой данных:", err)
	}
	defer db.Close()
	var pass string
	fmt.Println("--------------------------")
	fmt.Println("Удоления Аккаента")
	fmt.Print("Чтобы удолить введите пароль: ")
	fmt.Scanln(&pass)
	if pass == "0" {
		MainWindow(Id)
	} else if pass == Password {
		_, err = db.Exec("DELETE FROM Users WHERE id = ? and password = ?", Id, Password)
		if err != nil {
			log.Fatal("Ошибка при выполнении SQL-запроса:", err)
		}
		fmt.Println("Аккаунт успешно удален")
		Registration.FirstWindow()
	} else {
		fmt.Println("Вы ввели не правельный пароль!")
		DeleteAccount(Id, Password)
	}

}
