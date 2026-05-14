package main

import (
	"fmt"
	"log"
)

type Card struct {
	Name   string // name is trivial for now
	Number string // card number
}

type Account struct {
	Number string // account number
}

// assume the user has a card associated with multiple accounts
type UserInfo struct {
	PIN      int16
	Accounts []*Account
}

type TransactionType int

const (
	TransactionDeposit = iota
	TransactionWithdraw
	TransactionBalance
)

type Transaction struct {
	Account *Account
	Type    TransactionType
	Amount  int64
}

type Balance struct {
	Account *Account
	Amount  int64
}

type DB struct {
	// Handle *DBConn
	CardHolder    map[string]*UserInfo
	AccountAmount map[string]int64
}

// a customer inserts his/her card and retrieves his/her associated accounts for selection
func (d *DB) VerifyPIN(card *Card, pin int16) ([]*Account, error) {
	if c, ok := d.CardHolder[card.Number]; !ok {
		return nil, fmt.Errorf("card not found")
	} else {
		if pin == c.PIN {
			return c.Accounts, nil
		} else {
			return nil, fmt.Errorf("unauthorized")
		}
	}
}

func (d *DB) Process(transaction *Transaction) (*Balance, error) {
	if a, ok := d.AccountAmount[transaction.Account.Number]; !ok {
		return nil, fmt.Errorf("account not found")
	} else {
		switch transaction.Type {
		case TransactionDeposit:
			a += transaction.Amount
		case TransactionWithdraw:
			a -= transaction.Amount
		default: // case TransactionBalance:
		}
		return &Balance{transaction.Account, a}, nil
	}
}

func NewTransaction(account *Account, trType TransactionType, amount int64) *Transaction {
	return &Transaction{account, trType, amount}
}

var dummyDB *DB

// package initializer
func init() {
	// dummy DB to ease unit tests
	dummyDB = &DB{
		CardHolder: map[string]*UserInfo{
			"321": {
				PIN:      111,
				Accounts: []*Account{{Number: "123"}, {Number: "234"}},
			},
			"543": {
				PIN:      222,
				Accounts: []*Account{{Number: "345"}, {Number: "456"}},
			},
			"765": {
				PIN:      333,
				Accounts: []*Account{{Number: "567"}},
			},
		},
		AccountAmount: map[string]int64{
			"123": 45678,
			"234": 56789,
			"345": 67890,
			"456": 789012,
			"567": 8901234,
		},
	}
}

func main() {
	db := dummyDB
	card := &Card{"John Doe", "321"}

	accounts, err := db.VerifyPIN(card, 111)
	if err != nil {
		log.Fatal(err)
	}

	tr := NewTransaction(accounts[0], TransactionBalance, 0)
	bal, err := db.Process(tr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Balance: %d\n", bal.Amount)

	tr.Type = TransactionDeposit
	tr.Amount = 10000
	bal, err = db.Process(tr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Balance: %d\n", bal.Amount)
}
