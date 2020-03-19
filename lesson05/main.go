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

func sayHello(w http.ResponseWriter, r *http.Request) {
	//定义模版
	//解析模版
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("Parse template failed,err%v", err)
		return
	}
	//渲染模版
	//u1 := User{
	//	Name:   "小王子",
	//	Gender: "男",
	//	Age:    23,
	//}

	m1 := map[string]interface{}{
		"Name":   "小王子",
		"Gender": "男",
		"Age":    23,
	}
	t.Execute(w, m1)
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9100", nil)
	if err != nil {
		fmt.Println("Http server start failed, err:%v", err)
		return
	}
}
