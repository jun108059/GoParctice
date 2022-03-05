# 1. 생성자 잘 만들기

## 1-1. 요구사항

계좌를 생성하는 코드를 작성하기.

1. Account struct 만들기
2. 계좌는 `계좌주인이름`과 `잔액`을 가지고 있어야 함.
3. 최초 계좌 생성 시 `잔액`은 `0` 으로 고정
4. `계좌주인이름`과 `잔액`은 클라이언트에서 임의로 바꿀 수 없어야 함.

# 2. 무지성으로 기능 개발하기

`account.go`

```go
package _02_accounts

// Account struct
type Account struct {
    owner   string // 계좌 주인 (private)
    balance int    // 잔액 (private)
}
```

`main.go`

```go
package main

import "github.com/jun108059/learngo/_02_accounts"

func main() {
  account := _02_accounts.Account{owner: "yj_park", balance: 20000}
}
```

## 2-1. 문제점

Account struct 내에 owner 와 balance 가 소문자로 작성되어 private이기 때문에 접근할 수 없는 것을 알 수 있다.

## 2-1. 해결방안

owner, balance 필드 public 으로 선언 (대문자로 적어주기)

`account.go`

```go
package _02_accounts

// Account struct
type Account struct {
    Owner   string // 계좌 주인 (public)
    Balance int    // 잔액 (public)
}
```

`main.go`

```go
package main

import "github.com/jun108059/learngo/_02_accounts"

func main() {
  account := _02_accounts.Account{Owner: "yj_park", Balance: 20000}
}
```

## 2-2. 문제점

계좌 주인 이름과 잔액을 둘다 클라이언트에서 설정하여 보낼 수 있기 때문에 잔액을 마음대로 설정하여 계좌를 생성할 수 있는 문제가 있음.

만약 Java, Python, Javascript 였다면 생성자를 통해서 제어할 수 있지만, Go는 생성자를 만들지 않는다.

계좌를 생성할 때 Owner(계좌 주인 이름)만 전달 받아서 생성하려면 어떻게 해야할 까?

## 2-2. 해결방안

1. balance(잔액)만 private 으로 변경하기

`account.go`

```go
package _02_accounts

// Account struct
type Account struct {
    Owner   string // 계좌 주인 (public)
    balance int    // 잔액 (private)
}
```

`main.go`

```go
package main

import "github.com/jun108059/learngo/_02_accounts"

func main() {
  account := _02_accounts.Account{Owner: "yj_park"}
}
```

이제 생성된 계좌의 잔액은 항상 0 으로 초기화되는 것을 알 수 있다.

## 2-3. 문제점

Owner 가 public 이기 때문에 클라이언트에서 계좌를 생성한 이후에 이름을 마음대로 재할당할 수 있는 문제가 있다.

`main.go`

```go
package main

import "github.com/jun108059/learngo/_02_accounts"

func main() {
  account := _02_accounts.Account{Owner: "yj_park"}
  account.Owner = "mandoo"
}
```

그렇다고 Owner 를 private 으로 만들면 Account 자체를 생성할 수 없게 되는 첫번째 문제 상황으로 돌아가게 된다.


# 3. Go 스타일로 생성자 만들기

construct(생성) 하거나 struct 를 만들도록 function 을 작성해보자.

> struct 자체는 private 으로 감추고, public function 을 통해 생성자를 만들어주는 방식이다.

`account.go`

```go
package _02_accounts

// Account struct
type Account struct {
    Owner   string // 계좌 주인 (public)
    balance int    // 잔액 (private)
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}
```

return 시 복사한 새로운 object 를 만들지 않고 **포인터**를 통해 object 의 address 를 return 하도록 function 을 작성했다.

> 복사해서 새로운 object 를 return 하게 되면 성능에 영향을 미친다.

`main.go`

```go
package main

import "github.com/jun108059/learngo/_02_accounts"

func main() {
	account := _02_accounts.NewAccount("yj park")
}
```

이제 `계좌주인이름`과 `잔액`은 임의로 클라이언트에서 변경할 수 없다.
