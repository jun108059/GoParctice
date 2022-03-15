# 1. Method 알아보기

잔액을 증가시키는 함수를 작성하려고 한다.

그런데, Go 에서 function 과 method 는 다르다. 자세히 알아보자.

## 1-1. function vs. method

`function` 표현식

```bash
fucn 함수이름(파라미터 파라미터타입) 리턴타입 {
}
```

`method` 표현식

```bash
func (리시버 리시버타입) 함수이름(파라미터 파라미터타입) 리턴타입 {
}
```

func 예약어와 함수이름 사이에 **리시버**만 추가하면 method 가 된다.

따라서, 잔액을 증가시키는 함수는 아래와 같이 쓰면 된다.

```go
package _02_accounts

// Deposit amount 만큼 계좌에 입금한다.
func (a Account) Deposit(amount int) {
	// 리시버는 struct 맨 앞글자를 소문자로 적으면 된다.
	a.balance += amount
}
```

## 1-2. 리시버(Receiver) 알아보기

### 리시버 작성 규칙

1. 항상 func <-> 함수이름 사이에 작성
2. 변수는 struct 의 맨 앞 글자를 소문자로 하나만 적는다.

리시버는 `struct`이고 method 를 만들어준다.


## 1-3. method 정의할 때 주의할 점

잔액을 출력하는 method도 작성해보자.

```go
// 계좌의 잔액을 조회 
func (a Account) GetBalance() int {
    return a.balance
}
```

main 함수에서 다음과 같이 호출해서 사용하면 된다.

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
}
```

출력 결과를 보면 계좌의 잔액은 `0`으로 나온다.

왜 그럴까?

### 1-3-1. 디커플링(리시버와 포인터)

리시버는 결국 struct로 구성되어있고, method가 호출되면 struct의 복사본을 만들게 된다.

디커플링은 인터페이스 등을 활용하여 모듈 간의 의존성을 최소화하는 방법으로 Go에서는 struct를 복사하여 이를 구현하였다.

코드에서 보면 `main()`에서 생성한 `account`를  `Deposit()`메소드에서 `a`로 복사한다는 의미이다.

우리가 설계한대로 Deposit() 메소드가 실행되기 위해서, 포인터를 통해 얕은 복사(주소값 복사)를 하면 된다.

```go
// Deposit amount 만큼 계좌에 입금한다.
func (a *Account) Deposit(amount int) {
    a.balance += amount
}
```

GetBalance() 메소드는 정상적으로 동작하는 것을 알 수 있는데, 깊은 복사로 struct를 복사해와서 필드인 balance의 데이터를 조회만 하기 때문에 상관 없다.

더 자세한 내용은 아래 블로그를 참고하면 좋을 것 같다.

https://velog.io/@cafefarm-johnny/Golang-%EB%94%94%EC%BB%A4%ED%94%8C%EB%A7%81-%EB%A9%94%EC%86%8C%EB%93%9C-Method


