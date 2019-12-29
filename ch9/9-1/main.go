package main

type WdData struct {
	ch chan bool
	amount int
}

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withDraws = make(chan WdData)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	ch := make(chan bool)
	wd := WdData{ch, amount}
	withDraws <- wd
	return <- ch
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case wdData := <- withDraws:
			if balance < wdData.amount {
				wdData.ch <- false
			} else {
				balance -= wdData.amount
				wdData.ch <- true
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
