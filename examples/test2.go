package tests

const N = 1000

// @ requires forall i, j int :: {arr[i], arr[j]} 0 <= i && i < j && j < N ==> arr[i] <= arr[j]
// @ ensures !found ==> forall i int :: {arr[i]} 0 <= i && i < len(arr) ==> arr[i] != target
// @ ensures found ==> 0 <= idx && idx < len(arr) && arr[idx] == target
func BinarySearchArr(arr [N]int, target int) (idx int, found bool) {
	low := 0
	high := len(arr)
	mid := 0
	for low < high {
		mid = (low + high) / 2
		if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low, low < len(arr) && arr[low] == target
}