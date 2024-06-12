package h

type Heap []int

func (h *Heap) Insert(val int) {
	*h = append(*h, val)
	hepifyUp(h, h.Len()-1)
}

func (h Heap) Len() int {
	return len(h)
}

func (h *Heap) Delete(val int) {
	if h.Len() == 0 {
		return
	}

	index := -1
	for i, v := range *h {
		if v == val {
			index = i
			break
		}
	}

	if index == -1 {
		return
	}

	(*h)[index], (*h)[len(*h)-1] = (*h)[len(*h)-1], (*h)[index]
	*h = (*h)[:len(*h)-1]

	hepifyDown(h, index)
}

func (h Heap) Peek() int {
	if h.Len() == 0 {
		return -1
	}

	return h[0]
}

func (h Heap) FindIndex(val int) int {
	for i, v := range h {
		if v == val {
			return i
		}
	}

	return -1
}

func (h *Heap) UpdateVal(index, val int) {
	if index >= h.Len() {
		return
	}

	tmp := (*h)[index]
	(*h)[index] = val
	if tmp > val {
		hepifyUp(h, index)
	} else {
		hepifyDown(h, index)
	}
}

func (h *Heap) IncreaseByValue(index, val int) {
	if index >= h.Len() {
		return
	}

	(*h)[index] += val
	hepifyDown(h, index)
}

func (h *Heap) DecreaseByValue(index, val int) {
	if index >= h.Len() {
		return
	}

	(*h)[index] -= val
	hepifyUp(h, index)
}

func hepifyUp(h *Heap, index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if (*h)[index] < (*h)[parent] {
			(*h)[index], (*h)[parent] = (*h)[parent], (*h)[index]
			index = parent
		} else {
			break
		}
	}
}

func hepifyDown(h *Heap, index int) {
	for {
		right_child := (index + 1) * 2
		left_child := right_child - 1

		smallest := index

		if left_child < len(*h) && (*h)[left_child] < (*h)[smallest] {
			smallest = left_child
		}

		if right_child < len(*h) && (*h)[right_child] < (*h)[smallest] {
			smallest = right_child
		}

		if smallest != index {
			(*h)[index], (*h)[smallest] = (*h)[smallest], (*h)[index]
			index = smallest
		} else {
			break
		}
	}
}
