# Dictionary And Map

Go에서는 struct를 쓰지 않고 type만을 쓰는 것도 가능하다.

## 1. Dictionary로 map 선언 및 호출

map 자료형을 dictionary 로 선언하면 다음과 같다.

```go
package _03_dict

// Dictionary type
type Dictionary map[string]string
```

이제 main 문에서 호출하려면 아래와 같이 생성하고, 값을 조회할 수 있다.

```go
package main

import (
	"fmt"
	"github.com/jun108059/GoPractice/_03_dict"
)

func main() {
	dictionary := _03_dict.Dictionary{"first": "Fist word"}
	fmt.Println(dictionary["first"])
}
```

## 2. Dictionary로 Method 사용하기

struct 처럼 Method 를 만들어서 사용할 수 있다.

### 2-1. map 조회 Method (Search)

map의 key를 조회하면 값을 반환하고, 없다면 error를 반환해주는 함수를 작성해보자.

```go
package _03_dict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("not Found")

func (d Dictionary) Search(word string) (string, error) {
	// map 조회 시 값이 있는지 여부를 함께 받아올 수 있다! (쩐다!)
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}
```

GoLang의 map 자료구조에서 엄청 편리한 조회 기능을 제공한다!

```go
value, exists := map[key]
```

map에 조회하고자 하는 key의 value와 함께 존재여부를 멀티 리턴해주기 때문에 에러처리가 엄청 쉬워진다. 

main 함수에서 호출하는 예제는 다음과 같다.

```go
package main

import (
	"fmt"
	"github.com/jun108059/GoPractice/_03_dict"
)

func main() {
	dictionary := _03_dict.Dictionary{"first": "First word"}
	searchWord, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(searchWord)
	}
}
```

```bash
// 호출 결과
not Found
```

### 2-2. map 추가 Method (Add)

map 자료구조에 데이터를 추가하는 method를 작성해보자.

고려할 점은 기존에 있는 데이터인지 Validation 검증을 선행하도록 구현한다. 

```go
package _03_dict

import "errors"

var errNotFound = errors.New("not Found")
var errAlreadyExist = errors.New("already Exist")

// 위 Search와 동일한 코드 생략

// Add map dictionary에 데이터 추가하기
func (d Dictionary) Add(key string, value string) error {
	_, err := d.Search(key)
	if err == errNotFound {
		// map 데이터 할당
		d[key] = value
	} else {
		return errAlreadyExist
	}
	return nil
}
```

작성해둔 Search Method를 활용해서 기존에 있는 값이라면 already Exist 에러를 알려주도록 검증했다.

main 에서 호출해보면 결과는 아래와 같다.

```go
package main

import (
	"fmt"
	"github.com/jun108059/GoPractice/_03_dict"
)

func main() {
	dictionary := _03_dict.Dictionary{}
	key := "FirstName"
	value := "YoungJun"
	errDict := dictionary.Add(key, value)
	if errDict != nil {
		fmt.Println(errDict)
	} else {
		fmt.Println("Add Complete")
	}
	searchValue, errDict2 := dictionary.Search(key)
	if errDict2 != nil {
		fmt.Println(errDict2)
	} else {
		fmt.Println("Search Value : ", searchValue)
	}
	errDict3 := dictionary.Add(key, value)
	if errDict3 != nil {
		fmt.Println(errDict3)
	} else {
		fmt.Println("Add Complete")
	}
}
```

```bash
// 호출 결과
Add Complete
Search Value :  YoungJun
already Exist
```
