package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func f1(w http.ResponseWriter, r *http.Request) {
	k := func(name string) (string, error) {
		return name + "年前", nil
	}
	t := template.New("f.tmpl")
	t.Funcs(template.FuncMap{
		"kua99": k,
	})
	_, err := t.ParseFiles("./lesson06/f.tmpl")
	if err != nil {
		fmt.Println("err:%v\n", err)
		return
	}
	name := "张万明"
	t.Execute(w, name)
}
func main() {
	http.HandleFunc("/", f1)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("HTTP server start failed,err:%v", err)
		return
	}
}
