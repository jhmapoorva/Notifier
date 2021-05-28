package DataStructures

import "errors"

type Set struct {
	size int
	m    map[string]struct{}
}

func NewSet() Set {
	s := Set{}
	s.size = 0
	s.m = make(map[string]struct{})
	return s
}

func (s *Set) Add(value string) error {
	if s.Contains(value) {
		return errors.New("value already exists")
	}
	s.m[value] = struct{}{}
	s.size += 1
	return nil
}

func (s *Set) Remove(value string) error {
	if !s.Contains(value) {
		return errors.New("value does not exist to delete")
	}
	delete(s.m, value)
	s.size -= 1
	return nil
}

func (s *Set) Contains(value string) bool {
	_, c := s.m[value]
	return c
}

func (s *Set) Size() int {
	return s.size
}
