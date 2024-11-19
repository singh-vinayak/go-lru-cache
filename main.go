package main

import "fmt"

const SIZE = 5

type Node struct {
    Val string
    Left *Node
    Right *Node
}

type Queue struct {
    Head *Node
    Tail *Node
    Length int
}

type Cache struct {
    Queue Queue
    Hash Hash
}

type Hash map[string]*Node

func NewCache() Cache {
    return Cache{
        Queue: NewQueue(),
        Hash: Hash{},
    }
}

func NewQueue() Queue {
    head := &Node{}
    tail := &Node{}
    head.Right = tail
    tail.Left = head
    return Queue{
        Head: head,
        Tail: tail,
    }
}

func (c *Cache) Check(str string) {
    node := &Node{}
    if Val, ok := c.Hash[str]; ok {
        node = c.Remove(Val)
    } else {
        node = &Node{Val: str}
    }
    c.Add(node)
    c.Hash[str] = node
}

func (c *Cache) Add(node *Node){
    fmt.Printf("adding %s\n", node.Val)
    temp := c.Queue.Head.Right
    c.Queue.Head.Right = node
    node.Left = c.Queue.Head
    node.Right = temp
    temp.Left = node
    c.Queue.Length +=1

    if c.Queue.Length > SIZE {
        c.Remove(c.Queue.Tail.Left)
    }

    c.Hash[node.Val] = node
} 

func (c *Cache) Remove(node *Node) *Node{
    fmt.Printf("remove %s\n", node.Val)
    left :=  node.Left
    right := node.Right
    left.Right = right
    right.Left = left
    c.Queue.Length -=1
    delete(c.Hash, node.Val)
    return node
}

func (c *Cache) Display() {
    c.Queue.Display()
}

func (q *Queue) Display() {
    node := q.Head.Right
    fmt.Printf("%d - [", q.Length)
    for i:=0;i<q.Length;i++ {
        fmt.Printf("{%s}",node.Val)
        if i<q.Length-1 {
            fmt.Printf("<--->")
        }
        node = node.Right
    }
    fmt.Printf("]\n")
}

func main() {
    fmt.Println("Start Cache");
    cache := NewCache()
    for _, word := range []string{"parrot", "avocado", "dragonfruit", "mango","tomato", "dragonfruit", "tree"} {
        cache.Check(word)
        cache.Display()
    }
}