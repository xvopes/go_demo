//package main
//
//import (
//	"fmt"
//	"io/ioutil"
//	"net/http"
//)
//
//func sayHello(w http.ResponseWriter, r *http.Request){
//	b, _ := ioutil.ReadFile("./hello.txt")
//	_, _ = fmt.Fprint(w, string(b))
//}
//
//func main() {
//	http.HandleFunc("/hello", sayHello)
//	err := http.ListenAndServe(":9091", nil)
//	if err != nil {
//		fmt.Print("http sever failed, err:%v\n", err)
//		return
//	}
//}

package main

//
//import (
//	"github.com/gin-gonic/gin"
//)
//
//func sayHello(ctx *gin.Context) {
//	ctx.JSON(200, gin.H{
//		"message": "Hello,golang",
//	})
//}
//func main() {
//	r := gin.Default() //返回默认的路由引擎
//
//	r.GET("/hello", sayHello)
//
//	r.GET("/book")
//
//	r.Run(":9091")
//}

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("Parse template failed,err%v", err)
		return
	}

	name := "小王子"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("render template failed,err:%v", err)
		return
	}
}
func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9100", nil)
	if err != nil {
		fmt.Println("Http server start failed, err:%v", err)
		return
	}
}
