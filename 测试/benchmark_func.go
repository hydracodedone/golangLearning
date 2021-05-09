package funcTest

func insertElementAtIndexThird(newElement int, slice *[]int) {
	index := 3
	if slice == nil {
		*slice = make([]int, 4)
		(*slice)[index] = newElement
		return
	}
	if len(*slice) < 4 {
		length := len(*slice)
		missingEle := make([]int, index-length+1)
		missingEle[index-length] = newElement
		*slice = append(*slice, missingEle...)
	} else {
		*slice = append(*slice, newElement) //扩容
		copy((*slice)[index+1:], (*slice)[index:])
		(*slice)[index] = newElement
	}
}
