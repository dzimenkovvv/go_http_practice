package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

var money int = 1000
var wallet int
var mtx sync.Mutex

func PayHandler(w http.ResponseWriter, r *http.Request) {
	moneyCount, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		msg := "Error to read body HTTP:" + err.Error()
		w.Write([]byte(msg))
		fmt.Println(msg)
		return
	}
	strMoneyCount := string(moneyCount)

	intMoneyCount, err := strconv.Atoi(strMoneyCount)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		msg := "Error convert money to int:" + err.Error()
		w.Write([]byte(msg))
		fmt.Println(msg)
		return
	}

	mtx.Lock()
	if money-intMoneyCount >= 0 {
		money -= intMoneyCount

		msg := "Покупка на " + strconv.Itoa(intMoneyCount) + " успешно совершена! \nОстаток баланса: " + strconv.Itoa(money) + " USD!"
		w.Write([]byte(msg))
		fmt.Println("Покупка успешно совершена!")
	} else {
		w.WriteHeader(http.StatusBadRequest)

		msg := "Недостаточно средств!"
		w.Write([]byte(msg))
		fmt.Println(msg)
	}
	mtx.Unlock()
}

func SaveHandler(w http.ResponseWriter, r *http.Request) {
	moneyCount, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		msg := "Error to read body HTTP:" + err.Error()
		w.Write([]byte(msg))
		fmt.Println(msg)
		return
	}

	strMoneyCount := string(moneyCount)

	intMoneyCount, err := strconv.Atoi(strMoneyCount)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		msg := "Error convert money to int:" + err.Error()
		w.Write([]byte(msg))
		fmt.Println(msg)
		return
	}

	mtx.Lock()
	if money >= intMoneyCount {
		money -= intMoneyCount

		wallet += intMoneyCount

		msg := "Успешно положено в кошелек " + strconv.Itoa(intMoneyCount) + " USD! \nОстаток на балансе: " + strconv.Itoa(money) + " USD \nВ кошельке: " + strconv.Itoa(wallet) + " USD"
		w.Write([]byte(msg))
		fmt.Println("Деньги успено перенесены в кошелек!")
	} else {
		msg := "Недостаточно средств для переноса в кошелек!"
		w.Write([]byte(msg))
		fmt.Println(msg)
	}
	mtx.Unlock()
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	moneyCount, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		msg := "Error to read body HTTP:" + err.Error()
		w.Write([]byte(msg))
		fmt.Println(msg)

		return
	}

	strMoneyCount := string(moneyCount)

	intMoneyCount, err := strconv.Atoi(strMoneyCount)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		msg := "Error convert money to int:" + err.Error()
		w.Write([]byte(msg))
		fmt.Println(msg)
		return
	}

	mtx.Lock()
	money += intMoneyCount

	msg := "Деньги успешно зачислены на баланс! \nБаланс: " + strconv.Itoa(money) + " USD"
	w.Write([]byte(msg))
	fmt.Println("Деньги успешно зачислены на баланс!")
	mtx.Unlock()
}

func ReturnBalanceHandler(w http.ResponseWriter, r *http.Request) {
	msg := "Текущий баланс: " + strconv.Itoa(money) + "\nВ кошельке: " + strconv.Itoa(wallet)
	w.Write([]byte(msg))
	fmt.Println("Текущий баланс отправлен пользователю!")
}

func main() {
	http.HandleFunc("/pay", PayHandler)
	http.HandleFunc("/save", SaveHandler)
	http.HandleFunc("/add", AddHandler)
	http.HandleFunc("/balance", ReturnBalanceHandler)

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Failed to listen HTTP request:", err)
	}

}
