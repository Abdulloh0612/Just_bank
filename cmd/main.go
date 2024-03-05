package main

import (
	"Go_Projects/Just_bank/Just_bank/Registration"
	"Go_Projects/Just_bank/Just_bank/UserWindow"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID       int
	Name     string
	Surname  string
	Password string
	Balance  int
}

func main() {
	entry := Registration.FirstWindow()
	UserWindow.MainWindow(entry)
}
