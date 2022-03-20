package _03_dict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound     = errors.New("not Found")
	errCantUpdate   = errors.New("can not Update non-existing word")
	errAlreadyExist = errors.New("already Exist")
)

// Search 검색 key의 value 찾기
func (d Dictionary) Search(word string) (string, error) {
	// map 조회 시 값이 있는지 여부를 함께 받아올 수 있다! (쩐다!)
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

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

// Update map dictionary 데이터 업데이트
func (d Dictionary) Update(key string, value string) error {
	_, err := d.Search(key)
	switch err {
	case nil:
		d[key] = value
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete map dictionary 데이터 삭제
func (d Dictionary) Delete(key string) {
	delete(d, key)
}
