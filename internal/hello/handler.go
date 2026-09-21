package internal

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func NewHelloHandler(router *http.ServeMux) {
	handler := &HelloHandler{} // Создаем инстанс HelloHandler
	router.HandleFunc("/hello", handler.Hello())
}

// возращаем из функции -> функцию http.HandlerFunc
func (handler *HelloHandler) Hello() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Hello its me")
	}
}
