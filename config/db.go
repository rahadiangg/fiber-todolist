package config

import (
	"fmt"
	"github.com/rahadiangg/fiber-todolist/todo"
	"gorm.io/driver/mysql"
	gorm "gorm.io/gorm"
	"log"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:mysql-local@tcp(127.0.0.1)/fiber_todolist?parseTime=True"))
	if err != nil {
		log.Fatal(fmt.Sprintf("Can't init db: %s", err.Error()))
	}

	err = db.AutoMigrate(todo.ToDoModel{})
	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
