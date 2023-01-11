package stack

type Stack struct {
	contents []rune
}

func New() *Stack {
	s := new(Stack)
	s.contents = make([]rune, 0)
	return s
}

func (s *Stack) Push(r rune) {
	s.contents = append(s.contents, r)
}

func (s *Stack) Pop() rune {
	lastIdx := len(s.contents) - 1
	r := s.contents[lastIdx]
	s.contents = s.contents[:lastIdx] // may be memory inefficient
	return r
}

func (s *Stack) Peek() rune {
	if len(s.contents) > 0 {
		return s.contents[len(s.contents)-1]
	}
	return 0x0000
}

func (s *Stack) Size() int {
	return len(s.contents)
}

func (s *Stack) Contents() []rune {
	return s.contents
}
