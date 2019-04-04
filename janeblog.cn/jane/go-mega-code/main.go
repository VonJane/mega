package main

import (
    "html/template"
    "net/http"
)

// User struct
type User struct {
    Username string
}

type Post struct{
    User User
    Body string
}

type IndexViewModel struct{
    Title string
    User User
    Posts []Post
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        u1 := User{Username: "jane"}
        u2 := User{Username: "wang"}

        posts := []Post{
            Post{User:u1,Body:"I'm tired!"},
            Post{User:u2,Body:"I just want rest!"},
        }
        v := IndexViewModel{Title : "HomePage",User : u1,Posts:posts}
        tpl, _ := template.ParseFiles("templates/index.html")
        tpl.Execute(w, &v)
    })
    http.ListenAndServe(":8888", nil)
}