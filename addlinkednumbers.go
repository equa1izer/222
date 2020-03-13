package main

func AddLinkedNumbers(num1, num2 *NodeAddL) *NodeAddL {
	n1 := listToInt(num1)
	n2 := listToInt(num2)
	sum := n1 + n2

	return intToList(sum)
}

func intToList(x int) *NodeAddL {

	var res *NodeAddL

	for x != 0 {
		res = PushFront(res, x%10)
		x /= 10
	}

	return res
}

func listToInt(list *NodeAddL) int {
	res := 0

	power := power(10, listLen(list))

	for it := list; it != nil; it = it.Next {
		power /= 10
		res = res + it.Num*power
	}

	return res
}

func power(a, b int) int {
	if b < 0 {
		return 0
	}
	if b == 0 {
		return 1
	}
	return a * power(a, b-1)
}

func listLen(l *NodeAddL) int {
	res := 0
	for it := l; it != nil; it = it.Next {
		res++
	}
	return res
}
