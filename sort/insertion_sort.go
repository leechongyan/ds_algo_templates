package sort

// InsertionSort operates similarly to sorting a hand of cards
// compare current index against all previous index values
// and when we find the correct index, we insert it there;
// in the worst case (when arr is sorted in reverse order), it takes O(N ^ 2) time
func InsertionSort(arr []int) (sorted []int) {
	sorted = make([]int, len(arr))
	_ = copy(sorted, arr)
	if len(sorted) < 2 {
		return sorted
	}
	// first item is always trivially sorted
	for i := 1; i < len(sorted); i++ {
		cur := i
		for cur > 0 && sorted[cur-1] > sorted[cur] {
			sorted[cur], sorted[cur-1] = sorted[cur-1], sorted[cur]
			cur--
		}
	}
	return sorted
}

// recursive version of the above
// 1 major change - this is modified IN PLACE
func recursive(arr []int) (sorted []int) {
	if len(arr) <= 1 {
		return arr
	}
	sorted, last := recursive(arr[:len(arr)-1]), arr[len(arr)-1]
	// find correct place to insert
	sorted = append(sorted, last)
	prev := len(sorted) - 2
	for prev >= 0 && last < sorted[prev] {
		sorted[prev], sorted[prev+1] = last, sorted[prev]
		prev--
	}
	return sorted
}
