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
	digit, err := strconv.ParseInt(str[0:1], 10, 64)
	if err != nil {
		return nil
	}
	node := &listNode{val: int(digit)}
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
		current.val = (n.val + m.val + carry) % 10
		carry = (n.val + m.val) / 10
		n = n.next
		m = m.next
		if n == nil || m == nil {
			break
		}
		current.next = &listNode{}
		current = current.next
	}
	if n == nil && m == nil {
		return sum
	}
	var remain *listNode
	if n == nil {
		remain = m
	}
	if m == nil {
		remain = n
	}
	current.next = &listNode{}
	current = current.next
	for {
		current.val = (remain.val + carry) % 10
		carry = (remain.val + carry) / 10
		remain = remain.next
		if remain == nil {
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
