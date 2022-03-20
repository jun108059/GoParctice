package _03_dict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound = errors.New("not Found")
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
