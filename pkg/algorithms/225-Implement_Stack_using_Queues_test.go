package algorithms

import "testing"

func TestMyStack(t *testing.T) {
    stack := Constructor()
    stack.Push(1)
    stack.Push(2)
    
    if stack.Top() != 2 {
        t.Fail()
    }

    if stack.Pop() != 2 {
        t.Fail()
    }

    if stack.Empty() != false {
        t.Fail()
    }
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

