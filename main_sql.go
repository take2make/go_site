package main

import (
  "fmt"
  "database/sql"
_ "github.com/go-sql-driver/mysql"
  )

type User struct {
  Name string `json:"name"`
  Age uint16 `json:"age"`
}

func main() {

  fmt.Println("work with MySql")

  db, err := sql.Open("mysql", "root:0901ya1998@tcp(127.0.0.1:3306)/golang")
  if err != nil {
    panic(err)
  }
  // Установка данных

  //insert, err := db.Query("INSERT INTO `users` (`name`, `age`) VALUES ('Bob', 28) ")
  //if err != nil {
  //  panic(err)
  //}
  //defer insert.Close()

  // Выборка данных

  res, err := db.Query("SELECT `name`, `age` FROM `users`")
  if err != nil {
    panic(err)
  }

  for res.Next() {
    var user User
    err = res.Scan(&user.Name, &user.Age)
    if err != nil {
      panic(err)
    }
    fmt.Printf("User: %s with age %d\n", user.Name, user.Age)
  }

  defer db.Close()
  fmt.Println("worked with sql")
}
