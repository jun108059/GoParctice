# Error 예외처리

Java에서는 Exception, try-catch로 처리하면 될 것 같은데 Go는 그런거 없다.

Go에서 예외처리하는 방법을 정리해보자.

## 1. 예외처리 예제

계좌에서 출금할 때 잔액이 부족하면 예외를 발생시키는 코드이다.

```go
package _02_accounts

import "errors"

// Account struct
type Account struct {
	owner   string // 계좌 주인 (private)
	balance int    // 잔액 (private)
}

// 예외처리를 위한 errors 생성
var errNoMoney = errors.New("can't withdraw you are poor")

// Withdraw 계좌에서 돈을 출금
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}
```

main문에서 아래와 같이 호출하면 예외가 발생한다.

```go
package main

import (
	"fmt"
	"github.com/jun108059/learngo/_02_accounts"
)

func main() {
	account := _02_accounts.NewAccount("yj park")
	account.Deposit(100)
	fmt.Println(account.GetBalance())
	// 예외 처리
	err := account.Withdraw(200)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.GetBalance())
}
```

```bash
// 결과
100
can't withdraw you are poor
100
```

## 2. error 타입

에러 타입은 다음과 같이 정의되어있다.

```go
type error interface {
    Error() string
}
```

Error() 메서드 한개를 가진 인터페이스다. 이 인터페이스를 구현한 어떤 타입이든 에러가 될 수 있다. 그냥 정수형도 사용자 정의 타입으로 만들면 에러로 사용 가능하고 멤버가 없는 구조체도 Error() 메서드만 구현하면 에러로 사용할 수 있다.

```go
// 정수를 사용자 정의 타입으로 정의
type intAsError int

// Error() 메서드 구현
func (i intAsError) Error() string {
    return fmt.Sprintf("Error Code(%d)", i)
}

// 빈 구조체
type EmptyError struct {
}

// Error() 메서드 구현
func (ee EmptyError) Error() string {
    return "속빈 에러"
}
```

errors 패키지의 New 메서드로 에러 오브젝트를 생성하는 것이 더 편리하고 직관적이다.