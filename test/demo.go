package main

import (
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

// ServeHTTP方法，绑定TestHandler
func (th *TestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handle")
	w.Write([]byte(string("Handle")))
}

func main() {
	http.Handle("/", &TestHandler{"Hi"}) //根路由
	http.HandleFunc("/test", SayHello)   //test路由
	http.ListenAndServe("127.0.0.1:4096", nil)
}
