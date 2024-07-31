package internal

import "fmt"

func NewCacheCreation() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}

	head.Right = tail
	tail.Left = head

	return Queue{Head: head, Tail: tail}
}

func (c *Cache) Check(str string) {
	var node *Node

	if val, ok := c.Hash[str]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{Val: str}
	}
	c.Add(node)
	c.Hash[str] = node
}

func (c *Cache) Remove(n *Node) *Node {
	left := n.Left
	right := n.Right

	right.Left = left
	left.Right = right
	c.Queue.Length -= 1

	delete(c.Hash, n.Val)
	return n
}

func (c *Cache) Add(n *Node) {

	temp := c.Queue.Head.Right

	c.Queue.Head.Right = n

	n.Left = c.Queue.Head
	n.Right = temp
	temp.Left = n

	c.Queue.Length++

	if c.Queue.Length > Size {
		c.Remove(c.Queue.Tail.Left)
	}

}

func (c *Cache) Display() {
	(c.Queue.Display())
}

func (q *Queue) Display() {
	node := q.Head.Right
	fmt.Printf("%d - [", q.Length)
	for i := 0; i < q.Length; i++ {
		fmt.Printf("{%s}", node.Val)

		if i < q.Length-1 {
			fmt.Printf("<-->")
		}
		node = node.Right
	}
	fmt.Println("]")
}
