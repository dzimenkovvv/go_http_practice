package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
)

var money atomic.Int64
var wallet atomic.Int64
var mtx sync.Mutex

func PayHandler(w http.ResponseWriter, r *http.Request) {
	moneyCount, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error to read body HTTP:", err)
		return
	}
	strMoneyCount := string(moneyCount)

	intMoneyCount, err := strconv.Atoi(strMoneyCount)
	if err != nil {
		fmt.Println("Error convert money to int:", err)
		return
	}

	mtx.Lock()
	if money.Load()-int64(intMoneyCount) >= 0 {
		money.Add(int64(-intMoneyCount))

		fmt.Println("Покупка на", intMoneyCount, "успешно совершена!")
		fmt.Println("Остаток баланса:", money.Load())
	} else {
		fmt.Println("Недостаточно средств для совершения покупки!")
	}
	mtx.Unlock()
}

func SaveHandler(w http.ResponseWriter, r *http.Request) {
	moneyCount, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error to read body HTTP:")
		return
	}

	strMoneyCount := string(moneyCount)

	intMoneyCount, err := strconv.Atoi(strMoneyCount)
	if err != nil {
		fmt.Println("Error convert money to int:", err)
		return
	}

	mtx.Lock()
	if money.Load() >= int64(intMoneyCount) {
		money.Add(int64(-intMoneyCount))

		wallet.Add(int64(intMoneyCount))

		fmt.Println("Успешно положено в кошелек", intMoneyCount, "USD!")
		fmt.Println("Остаток на балансе:", money.Load(), "USD")
		fmt.Println("В кошельке:", wallet.Load(), "USD")
	} else {
		fmt.Println("Недостаточно средств для переноса в кошелек!")
	}
	mtx.Unlock()
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	moneyCount, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error to read body HTTP:", err)
		return
	}

	strMoneyCount := string(moneyCount)

	intMoneyCount, err := strconv.Atoi(strMoneyCount)
	if err != nil {
		fmt.Println("Error convert money to int:", err)
		return
	}

	mtx.Lock()
	money.Add(int64(intMoneyCount))
	fmt.Println("Деньги успешно зачислены на баланс!")
	fmt.Println("Баланс:", money.Load(), "USD")
	mtx.Unlock()
}

func main() {
	money.Add(1000)

	http.HandleFunc("/pay", PayHandler)
	http.HandleFunc("/save", SaveHandler)
	http.HandleFunc("/add", AddHandler)

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Failed to listen HTTP request:", err)
	}

}
