package my_stack

type MyStack struct {
	els []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.els = append(this.els, x)
}

func (this *MyStack) Pop() int {
	length := len(this.els)

	el := this.els[length-1]
	this.els = this.els[:length-1]

	return el
}

func (this *MyStack) Top() int {
	return this.els[len(this.els)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.els) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
