package slicelist

import (
	"errors"
	"fmt"
)

type SliceList[T comparable] struct {
	slice []T
}

func NewSliceList[T comparable]() *SliceList[T] { // O(1)
	return &SliceList[T]{}
}

// Agrega un elemento 'item' al final de la lista. // O(1)
func (s *SliceList[T]) Append(item T) {
	s.slice = append(s.slice, item)
}

// Agrega un elemento 'item' al principio de la lista. // O(n)
func (s *SliceList[T]) Prepend(item T) {
	s.slice = append([]T{item}, s.slice...)
}

// Agrega un elemento 'item' en la posición 'posicion' de la lista. // O(n)
func (s *SliceList[T]) InsertAt(posicion int, item T) {
	s.slice = append(s.slice[:posicion], append([]T{item}, s.slice[posicion:]...)...)
}

// Elimina la primera ocurrencia del elemento item de la lista. // O(n)
// Devuelve true si el elemento fue eliminado con éxito, y false en caso contrario.
func (s *SliceList[T]) Remove(item T) bool {
	for i := 0; i < len(s.slice); i++ {
		if s.slice[i] == item {
			s.slice = append(s.slice[:i], s.slice[i+1:]...)
			return true
		}
	}
	return false
}

// Devuelve el elemento en la posición 'posicion' de la lista. // O(1)
// Si la posición es inválida o la lista está vacía, devuelve un error.
func (l *SliceList[T]) Get(posicion int) (T, error) {
	if posicion < 0 || posicion >= len(l.slice) {
		var t T
		return t, errors.New("Lista vacía")
	}
	return l.slice[posicion], nil
}

// Devuelve el número de elementos en la lista. // O(1)
func (l *SliceList[T]) Size() int {
	return len(l.slice)
}

// Devuelve una representación en forma de cadena de la lista. O(n)
func (l *SliceList[T]) String() string {
	if l.Size() == 0 {
		return "[]"
	}
	result := "["
	for i := 0; i < l.Size(); i++ {
		obtenido, err := l.Get(i)
		_, errorProxPos := l.Get(i + 1)
		if err == nil && errorProxPos == nil {
			result += fmt.Sprintf("%v,", obtenido)
		} else {
			result += fmt.Sprintf("%v", obtenido)
		}
	}
	result += "]"
	return result
}
