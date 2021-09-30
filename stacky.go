package stacky


type Stack struct {
	items []int
}
type Queue struct {
	items []int
}
func (s *Stack) Push(i int)  {
	s.items=append(s.items,i)
}
func (q *Queue) Enqueue(i int)  {
	q.items=append(q.items, i)
}
func (s *Stack) Pop() int {
	removable:=s.items[len(s.items)-1]
	s.items=s.items[:len(s.items)-1]
	return removable
}
func (q *Queue) Dequeue()  {
	q.items= q.items[1:]
}
func (s *Stack) Peek() int {
	removable:=s.items[len(s.items)-1]
	return removable
}
func (s *Stack) IsEmpty() bool  {
	if len(s.items)==0 {
		return true
	}
	return false
}