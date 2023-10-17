package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree),
// создать для нее собственное множество.

type set struct {
	m map[any]struct{}
}

func NewSet() *set {
	return &set{make(map[any]struct{})}
}

func (s *set) Add(value any) {
	if _, ok := s.m[value]; ok {
		return
	}
	s.m[value] = struct{}{}
}

func (s *set) GetAll() []any {
	tmp := make([]any, 0, len(s.m))
	for value := range s.m {
		tmp = append(tmp, value)
	}
	return tmp
}

func (s *set) IsExist(value any) bool {
	if _, ok := s.m[value]; ok {
		return true
	}
	return false
}

func main() {
	strSet := NewSet()
	for _, val := range []string {"cat", "cat", "dog", "cat", "tree"} {
		strSet.Add(val)
	}
	fmt.Println(strSet.GetAll()...) 
}