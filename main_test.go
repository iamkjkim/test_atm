package main

import (
	"testing"
)

func TestVerifyPIN(t *testing.T) {
	db := dummyDB
	cases := []struct {
		card    *Card
		pin     int16
		wantErr bool
	}{
		{&Card{"John Doe", "321"}, 111, false},
		{&Card{"John Doe", "321"}, 222, true},
	}

	for _, c := range cases {
		t.Run(c.card.Name, func(t *testing.T) {
			accounts, err := db.VerifyPIN(c.card, c.pin)
			if err == nil && c.wantErr {
				t.Fatalf("expected error; however value returned: %+v", accounts)
			} else if err != nil && !c.wantErr {
				t.Fatalf("expected values; however error returned: %+v", err)
			}
		})
	}
}

func TestBalance(t *testing.T) {
	db := dummyDB
	cases := []struct {
		card            *Card
		pin             int16
		expectedBalance int64
	}{
		{&Card{"John Doe", "321"}, 111, 45678},
		{&Card{"Jane Doe", "543"}, 222, 67890},
	}

	for _, c := range cases {
		t.Run(c.card.Name, func(t *testing.T) {
			accounts, err := db.VerifyPIN(c.card, c.pin)
			if err != nil {
				t.Fatalf("error %v", err)
			}

			tr := NewTransaction(accounts[0], TransactionBalance, 0) // select the first account for now
			bal, err := db.Process(tr)
			if err != nil {
				t.Fatalf("error %v", err)
			}
			if bal.Amount != c.expectedBalance {
				t.Fatalf("expected %d; however %d returned", c.expectedBalance, bal.Amount)
			}
		})
	}
}

func TestDeposit(t *testing.T) {
	db := dummyDB
	cases := []struct {
		card            *Card
		pin             int16
		expectedBalance int64
	}{
		{&Card{"John Doe", "321"}, 111, 95678},
		{&Card{"Jane Doe", "543"}, 222, 117890},
	}

	for _, c := range cases {
		t.Run(c.card.Name, func(t *testing.T) {
			accounts, err := db.VerifyPIN(c.card, c.pin)
			if err != nil {
				t.Fatalf("error %v", err)
			}

			tr := NewTransaction(accounts[0], TransactionDeposit, 50000) // select the first account for now
			bal, err := db.Process(tr)
			if err != nil {
				t.Fatalf("error %v", err)
			}
			if bal.Amount != c.expectedBalance {
				t.Fatalf("expected %d; however %d returned", c.expectedBalance, bal.Amount)
			}
		})
	}
}

func TestWidthdraw(t *testing.T) {
	db := dummyDB
	cases := []struct {
		card            *Card
		pin             int16
		expectedBalance int64
	}{
		{&Card{"John Doe", "321"}, 111, 95678},
		{&Card{"Jane Doe", "543"}, 222, 117890},
	}

	for _, c := range cases {
		t.Run(c.card.Name, func(t *testing.T) {
			accounts, err := db.VerifyPIN(c.card, c.pin)
			if err != nil {
				t.Fatalf("error %v", err)
			}

			tr := NewTransaction(accounts[0], TransactionDeposit, 50000) // select the first account for now
			bal, err := db.Process(tr)
			if err != nil {
				t.Fatalf("error %v", err)
			}
			if bal.Amount != c.expectedBalance {
				t.Fatalf("expected %d; however %d returned", c.expectedBalance, bal.Amount)
			}
		})
	}
}
