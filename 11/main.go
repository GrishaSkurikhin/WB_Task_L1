package main

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств.

type set struct {
	m map[any]struct{} // используем ключи мапы для хранения элементов
}

// Конструктор типа "множество"
func NewSet() *set {
	return &set{make(map[any]struct{})}
}

// Добавление в множество
func (s *set) Add(value any) {
	if _, ok := s.m[value]; ok {
		return
	}
	s.m[value] = struct{}{}
}

// Получение всех элементов множества в виде слайса
func (s *set) GetAll() []any {
	tmp := make([]any, 0, len(s.m))
	for value := range s.m {
		tmp = append(tmp, value)
	}
	return tmp
}

// Проверка на существования элемента в множестве
func (s *set) IsExist(value any) bool {
	if _, ok := s.m[value]; ok {
		return true
	}
	return false
}

// Пересечение множеств
func intersection(set1, set2 *set) *set {
	resultSet := NewSet()
	for _, value := range set1.GetAll() {
		if set2.IsExist(value) {
			resultSet.Add(value)
		}
	}
	return resultSet
}

func main() {
	set1 := NewSet()
	for _, val := range []int {1, 2, 3} {
		set1.Add(val)
	}
	set2 := NewSet()
	for _, val := range []int {0, 2, 3, 6} {
		set2.Add(val)
	}
	set3 := intersection(set1, set2)
	fmt.Println(set3.GetAll()...)
}