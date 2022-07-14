package main

import (
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type Category uint

const (
	DEPOSIT Category = iota
	WITHDRAWAL
	INTEREST
	FEE
	TRANSFER
)

type Ledger struct {
	Reconciliations []Reconciliation
	Transactions    []Transaction
	// ClosingBalance is the result of tabulating operations in
	// Transactions after the most recent Reconciliation
}

type Reconciliation struct {
	Before         time.Time
	Transactions   []Transaction
	ClosingBalance money.Money
}

type Transaction struct {
	ID          uuid.UUID
	Description string
	DateTime    time.Time
	Amount      money.Money
}
