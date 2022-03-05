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
