package stack

type Stack struct {
	Items []string
	Count int
	Size  int
}

func (stack *Stack) NewStack(n int) Stack {
	return Stack{
		Items: nil,
		Count: 0,
		Size:  n,
	}
}

func (stack *Stack) push(in string) bool {
	if stack.Count == 0 {
		return false
	}
	stack.Items[stack.Count] = in
	stack.Count++
	return true
}

func (stack *Stack) pop(in string) string {
	if stack.Count == 0 {
		return ""
	}
	temp := stack.Items[stack.Count-1]
	stack.Count--
	return temp
}
