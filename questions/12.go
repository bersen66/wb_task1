package questions

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
// собственное множество.

type node struct {
	left   *node
	right  *node
	parent *node
	value  string
}

func (n *node) isLeaf() bool {
	return n.left == nil && n.right == nil
}

func (n *node) isRoot() bool {
	return n.parent == nil
}

type Iterator struct {
	data *node
}

func newIterator(node *node) Iterator {
	return Iterator{data: node}
}

func (it *Iterator) MoveNext() bool {
	c := it.data
	if c.isLeaf() {
		switch {
		case c.isRoot():
			it.data = nil
			return false
		case c.parent.left == c:
			it.data = c.parent
			return true
		case c.parent.right == c:
			for {
				if c.isRoot() {
					it.data = nil
					return false
				}
				c = c.parent
				if c.value > it.data.value {
					it.data = c
					return true
				}
			}
		}
	} else if c.right != nil {
		c = c.right
		for c.left != nil {
			c = c.left
		}
		it.data = c
		return true
	} else if c.right == nil {
		for {
			if c.isRoot() {
				it.data = nil
				return false
			}
			c = c.parent
			if c.value > it.data.value {
				it.data = c
				return true
			}
		}
	}

	return false
}

func (it *Iterator) Value() string {
	return it.data.value
}

func (it *Iterator) IsValid() bool {
	return it.data != nil
}

type Set struct {
	root *node
	size int
}

func NewSet() *Set {
	return &Set{
		root: nil,
		size: 0,
	}
}

func (s *Set) Empty() bool {
	return s.size == 0 && s.root == nil
}

func (s *Set) Size() int {
	return s.size
}

func (s *Set) Insert(v string) bool {
	if s.Empty() {
		s.root = &node{
			left:   nil,
			right:  nil,
			parent: nil,
			value:  v,
		}
		s.size++
		return true
	}

	for cur := s.root; ; {
		if cur.value == v {
			return false
		}
		if cur.value < v {
			if cur.right == nil {
				cur.right = &node{
					left:   nil,
					right:  nil,
					parent: cur,
					value:  v,
				}
				s.size++
				return true
			}
			cur = cur.right
			continue

		}
		if cur.value > v {
			if cur.left == nil {
				cur.left = &node{
					left:   nil,
					right:  nil,
					parent: cur,
					value:  v,
				}
				s.size++
				return true
			}
			cur = cur.left
			continue
		}
	}
}

func (s *Set) Iterator() Iterator {
	curr := s.root
	for curr.left != nil {
		curr = curr.left
	}
	return newIterator(curr)
}
