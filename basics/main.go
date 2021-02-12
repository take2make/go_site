package main

import ("fmt"; "net/http"; "html/template")
// импортируем дополнительные пакеты

type User struct {
  Name string
  Age uint16
  Money int16
  Avg_grades, Happiness float64
  Hobbies []string
}



func (u User) getAllInfo() string {
  return fmt.Sprintf("username is: %s. He is %d and he has money: %d", u.Name, u.Age, u.Money)
}

func (u *User) SetNewName(newName string) {
  // чтобы изменить саму струтуру НУЖНА ссылка
  u.Name = newName
}

func home_page2(w http.ResponseWriter, r *http.Request){
  // Request - запрос, который всегда передается
  // -------------------------------------------- //
  fmt.Fprintf(w, "hello Go Го\n")
  // задаем форматированную строку
  // w - куда записываем
  // -------------------------------------------- //
  bob := User{"Bob", 25, -50, 4.2, 0.8, []string{"fotbool", "Skate", "dance"}}
  // bob.name = "Alex" - set new name

  bob.SetNewName("Alex")
  fmt.Fprintf(w, bob.getAllInfo())

}

func home_page(w http.ResponseWriter, r *http.Request) {
  // fmt.Fprintf(w, "<b>Main Text</b>")
  bob := User{"Bob", 25, -50, 4.2, 0.8, []string{"fotbool", "Skate", "dance"}}
  teml, _ := template.ParseFiles("templates/home_page.html")
  teml.Execute(w, bob)
}


func contacts_page(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "contact page")
}

func handleRequest() {
  http.HandleFunc("/", home_page)
  http.HandleFunc("/contacts/", contacts_page)
  http.ListenAndServe(":8080", nil)
}

func main() {
  // var bob User = ...
  // bob := User{name: "Bob", age: 25, money: -50, avg_grades: 4.2, happiness: 0.8}

  handleRequest()
}
