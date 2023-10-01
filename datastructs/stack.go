package parsing

type Stack[T any] []T

func (s *Stack[T]) Empty() bool {
	return len(*s) == 0
}

func (s *Stack[T]) Push(val T) {
	*s = append(*s, val)
}

func (s *Stack[T]) Top() (T, bool) {
	if s.Empty() {
		var zero T
		return zero, false
	}
	return (*s)[len(*s)-1], true
}

func (s *Stack[T]) Pop() (T, bool) {
	top, ok := s.Top()
	if ok {
		*s = (*s)[:len(*s)-1]
	}
	return top, ok
}

func (s *Stack[T]) ToSlice() []T {
	return *s
}
