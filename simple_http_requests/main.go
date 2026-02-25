package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	str := "Hello world!"

	b := []byte(str)
	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Во время записи HTTP ответа возникла ошибка:", error.Error(err))
	} else {
		fmt.Println("Запрос успешно выполнен!")
	}
}

func handlerPayCancel(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Оплата успешно отменена!"))
	if err != nil {
		fmt.Println("Во время записи HTTP ответа возникла ошибка:", error.Error(err))
	} else {
		fmt.Println("Я успешно отменил оплату!")
	}
}

func handlerPay(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Оплата успешно произведена!"))
	if err != nil {
		fmt.Println("Во время записи HTTP ответа возникла ошибка:", error.Error(err))
	} else {
		fmt.Println("Я успешно произвел оплату!")
	}
}

func main() {
	http.HandleFunc("/default", handler)
	http.HandleFunc("/cancel", handlerPayCancel)
	http.HandleFunc("/pay", handlerPay)

	fmt.Println("Начал слушать...")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Возникла ошибка чтения http запроса!")
	}
}
