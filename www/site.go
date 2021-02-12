package main

import (
  "fmt"
  "net/http"
  "html/template"
  "database/sql"
  "github.com/gorilla/mux"
_ "github.com/go-sql-driver/mysql"
)

type Article struct {
  Id uint16
  Title, Anons, FullText string
}

func index(w http.ResponseWriter, r *http.Request) {
  t, err := template.ParseFiles("templates/index.html", "templates/header.html")

  if err != nil {
    fmt.Fprintf(w, err.Error())
  }

  db, err := sql.Open("mysql", "root:0901ya1998@tcp(127.0.0.1:3306)/golang")
  if err != nil {
    panic(err)
  }

  defer db.Close()

  res, err := db.Query("SELECT `id`, `title`, `anons`, `full_text` FROM `article2`")
  if err != nil {
    panic(err)
  }

  var posts = []Article{}
  for res.Next() {
    var post Article
    err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.FullText)
    if err != nil {
      panic(err)
    }

    posts = append(posts, post)
  }

  // динамическое подключение
  t.ExecuteTemplate(w, "index", posts)
}

func create(w http.ResponseWriter, r *http.Request) {
  t, err := template.ParseFiles("templates/create.html", "templates/header.html")
  if err != nil {
    panic(err)
  }
  t.ExecuteTemplate(w, "create", nil)
}

func save_article(w http.ResponseWriter, r *http.Request) {
  // get data from form
  title := r.FormValue("title")
  anons := r.FormValue("anons")
  full_text := r.FormValue("full_text")

  if title == "" || anons == "" || full_text == "" {
    // можно выводить целый шаблон html
    fmt.Fprintf(w, "Не все данные заполнены")
  }

  db, err := sql.Open("mysql", "root:0901ya1998@tcp(127.0.0.1:3306)/golang")
  if err != nil {
    panic(err)
  }

  defer db.Close()

  insert, err := db.Query(fmt.Sprintf("INSERT INTO `article2` (`title`, `anons`, `full_text`) VALUES('%s', '%s', '%s')", title, anons, full_text))
  if err != nil {
    panic(err)
  }
  defer insert.Close()

  // переадрисация пользователя с кодом 301
  http.Redirect(w, r, "/", http.StatusSeeOther)
}

func show_post(w http.ResponseWriter, r *http.Request) {

  t, err := template.ParseFiles("templates/show.html", "templates/header.html")

  if err != nil {
    fmt.Fprintf(w, err.Error())
  }

  vars := mux.Vars(r)
  //fmt.Fprintf(w, "Id: %v\n", vars["id"])

  // формируем sql запрос для данной статьи
  db, err := sql.Open("mysql", "root:0901ya1998@tcp(127.0.0.1:3306)/golang")
  if err != nil {
    panic(err)
  }

  defer db.Close()

  res, err := db.Query(fmt.Sprintf("SELECT * FROM `article2` WHERE `id` = '%s'", vars["id"]))
  if err != nil {
    panic(err)
  }

  var showPost = Article{}
  for res.Next() {
    var post Article
    err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.FullText)
    if err != nil {
      panic(err)
    }

    showPost = post
  }

  // динамическое подключение
  t.ExecuteTemplate(w, "show", showPost)

}

func handleFunc() {
  router := mux.NewRouter()
  router.HandleFunc("/create", create).Methods("GET")
  router.HandleFunc("/save_article", save_article).Methods("POST")
  router.HandleFunc("/", index).Methods("GET")
  router.HandleFunc("/post/{id:[0-9]+}", show_post).Methods("GET")

  http.Handle("/", router) // обработка всех url адресов через router
  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
  http.ListenAndServe(":8000", nil)
}

func main() {
  handleFunc()
}
