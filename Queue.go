/************************************************************
*  Name         : Avinash Narasimhan                        *
*  University   : National University of Singapore          *
*  Discipline   : Electrical Engineering                    *
*  E-Mail       : e0178378@u.nus.edu                        *
*  Title        : Implementation of Queue using Two Stacks  *
*  Submitted to : A-STAR (IHPC)                             *
************************************************************/

package main

//import fmt package
import (
	"fmt"
)

/* *** STACK IMPLEMENTATION **** */

// creates a new structure (STACK)
func NewStack() *Stack {
	return &Stack{}
}

// Stack Structure
type Stack struct {
      arr []int
}

// A Push function for a stack variable
func (s *Stack) stackPush(value int) {
	s.arr = append(s.arr, value)
}

// Pop function for a Stack variable
func (s *Stack) stackPop() (int){
    length := len(s.arr)
    if length == 0 {
        return 0
    }

    res := s.arr[length-1]

    // Array length is reduced by 1
    s.arr = s.arr[:length-1] 
    return res
}

// Finding the topmost element of a stack
func (s *Stack) stackTop() (int) {
    length := len(s.arr)
    if length == 0 {
        return 0
    }

    top := s.arr[length-1]

    return top

}

// Check if the stack is empty and return a boolean value
func (s *Stack) stackEmpty() (bool) {
    length := len(s.arr)
    if length == 0 {
        return true
    }
    return false
}

/* ********QUEUE IMPLEMENTATION USING TWO STACK VARIABLES******/

// Creates a new Queue structure
func NewQueue() *Queue {

	return &Queue{}
}

// Queue Structure with two stack variables
type Queue struct {
	s1 Stack
	s2 Stack
}

// Pushes the element into the Queue
// elements of stack s1 are popped to stack s2
// new element is pushed into stack s1
// elements of stack s2 is popped to stack s1 
func (q *Queue) queuePush(value int) {

	for !(q.s1.stackEmpty()) {
		q.s2.stackPush(q.s1.stackTop())
		q.s1.stackPop()
	}

        q.s1.stackPush(value)

 	for !(q.s2.stackEmpty()) {
		q.s1.stackPush(q.s2.stackTop())
		q.s2.stackPop()
	}
	
}

// Queue pop operation
// The first element of the Stack s1 is the front element of 
// the queue
func (q *Queue) queuePop() (int) {


	return q.s1.stackPop()

}

// Front most element of the queue is the topmost element of 
// Stack s1
func (q *Queue) queueFront() (int) {

        value := q.s1.stackTop()

    	return value
}

func (q *Queue) queueEmpty() {

        length := len(q.s1.arr)  
			  if length == 0{
			      fmt.Println("Queue is empty")
			  } else{
            fmt.Println("Queue is not empty")
            fmt.Println("Queue elements are:",q.s1.arr)
        }

}


func main() {
// new queue variable
	q := NewQueue()

	var choice int
	cont := 1
	
	for cont != 0 {
		fmt.Println("What do you want to do?\n1. Push an element\n
		2. Pop elements\n3. Look at the Top element\n4. Is the Queue Empty?")

    	fmt.Println("Enter your choice:")
		fmt.Scan(&choice)
		var num int
		
		switch choice{
			case 1:
        		i := 1
        		for i != 0{
					fmt.Println(" Enter an element")
					fmt.Scan(&num)
					q.queuePush(num)
         			fmt.Println(" Push more elements? (Yes: 1, No: 0)")
          			fmt.Scan(&i)
        		}
        		fmt.Println("Final Queue: ",q.s1)
			case 2:
				fmt.Println("How many elements would you like to pop:")
				fmt.Scan(&num)
				length := len(q.s1.arr)
				
				if length == 0 {
				    fmt.Println(" Queue is empty")
				} else if num > length {
					fmt.Println(" The queue does not contain so many elements")
				} else {
					i := 0
					fmt.Println("The following are the elements:")
					for i < num {
						fmt.Println(q.queuePop())
						i++
					}
				}
			case 3:
			  	length := len(q.s1.arr)  
			  	if length == 0{
			  	    fmt.Println("Queue is empty")
			  	} else {
				    fmt.Println("The topmost element of the queue: ",q.queueFront())
			  	}
      		case 4:
          		q.queueEmpty()
		}
		fmt.Println("Do you want to continue? (Yes: 1, No: 0)")
		fmt.Scan(&cont)
	}				
}