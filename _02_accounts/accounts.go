package _02_accounts

// Account struct
type Account struct {
	owner   string // 계좌 주인 (private)
	balance int    // 잔액 (private)
}

// NewAccount Account 생성 (생성자로 동작)
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit 계좌에 돈 입금하는 method
func (a Account) Deposit(amount int) {
	a.balance += amount
}

// GetBalance 계좌 잔액 불러오기
func (a Account) GetBalance() int {
	return a.balance
}
