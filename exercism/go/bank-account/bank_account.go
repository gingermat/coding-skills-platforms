package account

import (
	"sync"
)

type BankAccount struct {
	sync.Mutex
	amount int
}

func (b *BankAccount) closed() bool {
	return b.amount < 0
}

func (b *BankAccount) Balance() (int, bool) {
	if b.closed() {
		return 0, false
	}

	return b.amount, true
}

func (b *BankAccount) Close() (int, bool) {
	b.Lock()
	defer b.Unlock()

	if b.closed() {
		return 0, false
	}

	amount := b.amount
	b.amount = -1

	return amount, true
}

func (b *BankAccount) Deposit(amount int) (int, bool) {
	b.Lock()
	defer b.Unlock()

	if b.closed() {
		return 0, false
	}

	if b.amount+amount < 0 {
		return b.amount, false
	}

	b.amount += amount

	return b.amount, true
}

func Open(amount int) *BankAccount {
	if amount < 0 {
		return nil
	}

	return &BankAccount{
		amount: amount,
	}
}
