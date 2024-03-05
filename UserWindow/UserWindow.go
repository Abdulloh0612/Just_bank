package UserWindow

import (
	"database/sql"
	"fmt"
	"log"
	"os"

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

func MainWindow(ID int) {
	var choice int
	Id, Name, Surname, Address, Password, Balance := Check(ID)
	fmt.Println("--------------------------")
	fmt.Printf("User:%d %s's Account \n", Id, Name)
	fmt.Println("")
	fmt.Println("1. Check account")
	fmt.Println("2. Edit account")
	fmt.Println("3. Put money")
	fmt.Println("4. Withdraw money")
	fmt.Println("5. Sending money")
	fmt.Println("6. Delete account")
	fmt.Println("0. Exit")
	fmt.Println("--------------------------")
	fmt.Print("Your choice is ")
	fmt.Scan(&choice)

	switch choice {
	case 0:
		os.Exit(0)
	case 1:
		fmt.Println("--------------------------")
		fmt.Printf("Your ID: %d\n", Id)
		fmt.Printf("Your Name: %s\n", Name)
		fmt.Printf("Your Surname: %s\n", Surname)
		fmt.Printf("Your numder/mail: %s\n", Address)
		fmt.Printf("Your Balance: %d\n", Balance)
		fmt.Println("--------------------------")
		MainWindow(Id)
	case 2:
		Edit(ID, Password)
	case 3:
		PutMoney(ID, Password)
	case 4:
		WithdrawMoney(ID, Password)
	case 5:
		SendingMoney(ID, Password)
	case 6:
		DeleteAccount(ID, Password)
	default:
		fmt.Println("Вы нажали неправельную номер!")
	}

}
func Check(id int) (Id int, Name, Surname, Address, Password string, Balance int) {
	db, err := sql.Open("mysql", "abdulloh:member1206@tcp(localhost:3306)/pet_Registration")
	if err != nil {
		log.Fatal("Ошибка при открытии соединения с базой данных:", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, surname, address, password, balance FROM Users WHERE id = ?", id)
	if err != nil {
		log.Fatal("Ошибка при выполнении SQL-запроса:", err)
	}
	defer rows.Close()

	var user Person

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.Surname, &user.Address, &user.Password, &user.Balance)
		if err != nil {
			log.Fatal("Error scanning result:", err)
		}
	}
	return user.ID, user.Name, user.Surname, user.Address, user.Password, user.Balance
}
