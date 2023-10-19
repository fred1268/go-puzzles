package addtwonumbers

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	t.Parallel()
	testData := []struct {
		name string
		val1 *listNode
		val2 *listNode
		sum  *listNode
	}{
		{
			name: "123+234",
			val1: newListNode(123),
			val2: newListNode(234),
			sum:  newListNode(357),
		},
		{
			name: "12345+24680",
			val1: newListNode(12345),
			val2: newListNode(24680),
			sum:  newListNode(36926),
		},
		{
			name: "234+564",
			val1: newListNode(243),
			val2: newListNode(564),
			sum:  newListNode(708),
		},
		{
			name: "0+0",
			val1: newListNode(0),
			val2: newListNode(0),
			sum:  newListNode(0),
		},
		{
			name: "9999999+9999",
			val1: newListNode(9999999),
			val2: newListNode(9999),
			sum:  newListNode(89990001),
		},
	}
	for _, tt := range testData {
		name := tt.name
		val1 := tt.val1
		val2 := tt.val2
		sum := tt.sum
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			result := addTwoNumbers(val1, val2)
			if !sum.equals(result) {
				t.Errorf("wanted: '%s', got '%s'", sum, result)
			}
		})
	}
}
