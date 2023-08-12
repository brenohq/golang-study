package main

import (
	"fmt"
	"sync"
	"time"
)

var p1 = fmt.Println

type Account struct {
	balance int
	lock    sync.Mutex
}

func (a *Account) GetBalance() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.balance
}

func (a *Account) WithdrawMoney(v int) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if v > a.balance {
		p1("Not enough money.")
	} else {
		fmt.Printf("Withdrawn: %d Balance: %d\n", v, a.balance)
		a.balance -= v
	}
}

func main() {
	var acct Account
	acct.balance = 100

	p1("Balance:", acct.GetBalance())

	for i := 0; i < 12; i++ {
		go acct.WithdrawMoney(10)
	}

	time.Sleep(2 * time.Second)
}
