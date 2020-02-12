/*
singly linked list
1. position starts from 0, so Append(position:0) means add new head
*/


package main

import (
    "fmt"
)

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct{
    Head *Node
}

func (ll *LinkedList) Append(value int){
    node := Node{value, nil}

    if ll.Head == nil {
        // linked list is empty
        ll.Head = &node
        return
    }

    currentNode := ll.Head
    for {
        if currentNode.Next == nil {
            currentNode.Next = &node
            return
        }
        currentNode = currentNode.Next

    }
    return
}

func (ll *LinkedList) Traverse(){
    if ll.Head == nil{
        return
    }

    for currentNode := ll.Head ;;currentNode=currentNode.Next{
		fmt.Printf("%d->", currentNode.Value)
        if currentNode.Next == nil{
		    fmt.Println("")
            break
        }
    }
    return
}

func (ll *LinkedList) Size() int{
    if ll.Head == nil{
        return 0
    }

    var size int
    for currentNode := ll.Head ;;currentNode=currentNode.Next{
        size ++
        if currentNode.Next == nil{
            break
        }
    }

    return size
}

//Insert insert the node at position
// if position less than 0 or greater than size of ll, return error
// if position is 0, replace the head with new node
// if postion greater than 0, find the previous node prevNode,
//   insert the node after prevNode
func (ll *LinkedList) Insert(position, value int) error{
    size := ll.Size()
    if position<0 || position> size {
        return fmt.Errorf("Linked list out of range!!!")
    }


    node := Node{value, nil}

    // insert the node at front
    if position == 0{
        node.Next = ll.Head
        ll.Head = &node
        return nil
    }

    // insert the node at middle or tail
    prevNode := ll.Head //the node before position
    for i:=0;i<position-1;i++{
        prevNode = prevNode.Next
    }
    node.Next = prevNode.Next
    prevNode.Next = &node
    return nil
}

//Remove remove node at position
// if position less than 0 or greater than size of ll, return error
// if position is 0, replace the head with second node
// if postion greater than 0, find the previous node prevNode,
//   insert the node after prevNode
func (ll *LinkedList) Remove(position int) error{
    size := ll.Size()
    if position<0 || position>=size {
        return fmt.Errorf("Linked list out of range!!!")
    }


    var node,prevNode *Node //the node to be removde and previous node
    // insert the node at front
    if position == 0{
        node = ll.Head
        ll.Head = node.Next
        return nil
    }

    // remove the node at middle or tail
    prevNode = ll.Head
    for i:=0;i<position-1;i++{
        prevNode = prevNode.Next
    }
    node = prevNode.Next
    prevNode.Next = node.Next
    return nil
}

//Reverse reverse the linked list
//leetcode#206 https://leetcode.com/problems/reverse-linked-list/
func (ll *LinkedList) Reverse(){

    size := ll.Size()

    if size < 2{
        return
    }
    if size == 2{
        head := ll.Head
        ll.Head = head.Next
        ll.Head.Next = head
        head.Next = nil
        return
    }

    var head, prev, current, later *Node
    head = ll.Head
    prev = head
    current = prev.Next
    later = current.Next
    head.Next = nil //head becomes tail
    for {
        current.Next = prev

        if later.Next == nil{
            // later is the tail
            ll.Head = later
            later.Next = current
            break
        }
        prev = current
        current = later
        later = later.Next
    }
}

func main() {

    ll := new(LinkedList)

    ll.Append(6)
    ll.Append(3)
    ll.Append(5) // ll: 6->3->5
    ll.Traverse()

    ll.Insert(2, 11) // ll: 6->3->11->5
    ll.Traverse()
    if err:=ll.Remove(1); err!=nil{ // ll: 6->11->5
        fmt.Println(err)
    }
    ll.Traverse()

    ll.Append(9) // ll 6->11->5->9
    ll.Traverse()
    ll.Reverse() // ll 9->5->11->6
    ll.Traverse()

    ll.Remove(1) // 9->11->6
    ll.Remove(0) // 11->6
    ll.Reverse() // 6->11
    ll.Traverse()
    ll.Remove(1) // 6
    ll.Reverse() // 6
    ll.Traverse()
}
