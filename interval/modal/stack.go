package modal

type Stack struct {
	s      []string
	top    string
	length int
}
func NewStack() *Stack {
	stack := &Stack{
		s: make([]string, 0, 10),
	}
	return stack
}
func (s *Stack) Push(item string) {
	s.s = append(s.s, item)
	s.top = item
	s.length += 1
}
func (s *Stack) Empty() bool {
	return s.length == 0
}
func (s *Stack) Pop() string {
	if s.Empty() {
		return ""
	}
	item := s.s[s.length-1]
	s.s = s.s[:s.length-1]
	s.length -= 1
	return item
}
func (s *Stack) Len() int {
	return s.length
}
func (s *Stack) Top() string {
	if s.Empty() {
		return ""
	} else {
		return s.top
	}
}