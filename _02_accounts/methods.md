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
2. 변수는 struct 의 맨 앞 글자를 소문자로 하나만 적는다

