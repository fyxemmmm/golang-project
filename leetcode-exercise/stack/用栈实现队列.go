package stack



type MyQueue struct {
	pushStack []int
	popStack []int
}


/** Initialize your data structure here. */
func ConstructorMyQueue() MyQueue {
	return MyQueue{}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.pushStack = append(this.pushStack, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.popStack) > 0 {
		res := this.popStack[len(this.popStack)-1]
		this.popStack = this.popStack[:len(this.popStack)-1]
		return res
	}
	this.move()
	res := this.popStack[len(this.popStack)-1]
	this.popStack = this.popStack[:len(this.popStack)-1]
	return res
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.popStack) > 0 {
		return this.popStack[len(this.popStack)-1]
	}

	this.move()
	return this.popStack[len(this.popStack)-1]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.pushStack) == 0 && len(this.popStack) == 0
}


func (this *MyQueue) move() {
	for len(this.pushStack) > 0 {
		res := this.pushStack[len(this.pushStack)-1]
		this.popStack = append(this.popStack, res)
		this.pushStack = this.pushStack[:len(this.pushStack)-1]
	}
}
