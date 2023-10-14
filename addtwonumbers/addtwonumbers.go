package addtwonumbers

import "strconv"

type listNode struct {
	val  int
	next *listNode
}

func (l *listNode) equals(other *listNode) bool {
	first := l
	for {
		if first.next == nil && other.next == nil {
			return true
		}
		if first.next == nil || other.next == nil {
			return false
		}
		if first.val != other.val {
			return false
		}
		first = first.next
		other = other.next
	}
}

func (l *listNode) String() string {
	var digits []byte
	current := l
	for {
		digits = append(digits, '0'+byte(current.val))
		current = current.next
		if current == nil {
			break
		}
	}
	return string(digits)
}

func newListNode(n int64) *listNode {
	str := strconv.FormatInt(n, 10)
	node := &listNode{val: int(str[0] - '0')}
	previous := node
	for _, c := range str[1:] {
		current := &listNode{}
		previous.next = current
		val, err := strconv.ParseInt(string(c), 10, 64)
		if err != nil {
			return nil
		}
		current.val = int(val)
		previous = current
	}
	return node
}

func addTwoNumbers(n, m *listNode) *listNode {
	carry := 0
	sum := &listNode{}
	current := sum
	for {
		current.val = carry
		if n != nil {
			current.val += n.val
			n = n.next
		}
		if m != nil {
			current.val += m.val
			m = m.next
		}
		carry = current.val / 10
		current.val = current.val % 10
		if n == nil && m == nil {
			break
		}
		current.next = &listNode{}
		current = current.next
	}
	if carry != 0 {
		current.next = &listNode{val: carry}
	}
	return sum
}
