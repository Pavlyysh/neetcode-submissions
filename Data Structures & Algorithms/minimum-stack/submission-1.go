type MinStack struct {
	stack []int
}

func Constructor() MinStack {
	return MinStack{stack:[]int{}}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	if len(this.stack)==0 {
		return
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	min := this.Top()
	for _, num := range this.stack {
		if num < min {
			min = num
		}
	}

	return min
}
