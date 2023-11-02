package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct{
	Name   string
	Gender string
	Age    int
}
func sayHello(w http.ResponseWriter,r *http.Request){
	t,err:=template.ParseFiles("./lesson04/.idea/hello.tmpl")
	if err!=nil{
		fmt.Println("parse template failed,err:%v",err)
		return
	}
	u1:=User{
		Name:   "张万明",
		Gender: "女",
		Age:    18,
	}
	t.Execute(w,u1)
}
func main(){
	http.HandleFunc("/",sayHello)
	err:=http.ListenAndServe(":9000",nil)
	if err!=nil{
		fmt.Println("HTTP server start failed,err:%v",err)
		return
	}
}