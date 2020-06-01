package arraylist

type ArrayList struct {
	size     int
	elements []int
}

func (list *ArrayList) Add(val int) {
	//grow and ensure capacity
	if list.size >= len(list.elements) {
		newCapacity := len(list.elements) * 2
		if newCapacity == 0 {
			newCapacity = 1
		}
		newElements := make([]int, newCapacity)
		copy(newElements, list.elements)
		list.elements = newElements
	}
	list.elements[list.size] = val
	list.size++
}

func (list *ArrayList) Remove(val int) {
	for i := 0; i < list.size; i++ {
		if list.elements[i] == val {
			copy(list.elements[i:list.size-1], list.elements[i+1:list.size])
			list.elements[list.size-1] = 0
			list.size--
		}
	}
}
