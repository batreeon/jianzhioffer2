type MinStack struct {
    stack []int
    min int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(x int)  {
    if len(this.stack) == 0 {
        this.min = x
    }
    if x < this.min {
        this.min = x
    }
    this.stack = append(this.stack,x)
}


func (this *MinStack) Pop()  {
    pop := this.stack[len(this.stack)-1]
    
    if pop == this.min {
        this.min = this.stack[0]
        for i := 1 ; i < len(this.stack)-1 ; i++ {
            if this.stack[i] < this.min {
                this.min = this.stack[i]
            }
        }
    }//若要pop的值是最小值，则在剩余元素中找最小值
    
    this.stack = this.stack[0:len(this.stack)-1]
}


func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
    return this.min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */