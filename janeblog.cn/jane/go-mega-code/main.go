package main

import (
    "html/template"
    "net/http"
    "io/ioutil"
    "os"
)

// User struct
type User struct {
    Username string
}

type Post struct{
    User 
    Body string
}

type IndexViewModel struct{
    Title string
    User 
    Posts []Post
}

// PopulateTemplates func
// Create map template name to template.Template
func PopulateTemplates() map[string]*template.Template {
	const basePath = "templates"
	result := make(map[string]*template.Template)

	layout := template.Must(template.ParseFiles(basePath + "/_base.html"))
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}
	for _, fi := range fis {
		func() {
			f, err := os.Open(basePath + "/content/" + fi.Name())
			if err != nil {
				panic("Failed to open template '" + fi.Name() + "'")
			}
			defer f.Close()
			content, err := ioutil.ReadAll(f)
			if err != nil {
				panic("Failed to read content from file '" + fi.Name() + "'")
			}
			tmpl := template.Must(layout.Clone())
			_, err = tmpl.Parse(string(content))
			if err != nil {
				panic("Failed to parse contents of '" + fi.Name() + "' as template")
			}
			result[fi.Name()] = tmpl
		}()
	}
	return result
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
       templates := PopulateTemplates()
       templates["index.html"].Execute(w,&v)
    })
    http.ListenAndServe(":8888", nil)
}