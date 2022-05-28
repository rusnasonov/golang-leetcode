package algorithms

type MyStack struct {
    list []int
}


func Constructor() MyStack {
   return MyStack{
        list: []int{},
    }
}


func (this *MyStack) Push(x int)  {
    this.list = append(this.list, x)
}


func (this *MyStack) Pop() int {
    head := this.list[len(this.list) - 1]
    this.list = this.list[:len(this.list) - 1]
    return head
}


func (this *MyStack) Top() int {
   return this.list[len(this.list) - 1]
}


func (this *MyStack) Empty() bool {
    return len(this.list) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

