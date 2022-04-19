package main

import (
	"fmt"
	"log"
	"net/http"
)

type TestHandler struct {
	str string
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	log.Printf("HandleFunc")
	w.Write([]byte(string("HandleFunc")))
}

//ServeHTTP方法，绑定TestHandler
func (th *TestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handle")
	w.Write([]byte(string("Handle")))
}

func PostTest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(string("post")))
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println(r.Form["name"])
	fmt.Println(r.Form["age"])
}

func main() {
	http.Handle("/", &TestHandler{"Hi"}) //根路由
	http.HandleFunc("/test", SayHello)   //test路由
	http.HandleFunc("/post", PostTest)   //test路由
	http.ListenAndServe("127.0.0.1:9999", nil)
}
