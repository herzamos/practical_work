package tests

type Sorter interface {
	// @ requires acc(arr, 1)
	// @ ensures acc(arr, 1)
	// @ ensures forall i, j int :: { i, j } 0 <= i && i <= j && j < len(arr) ==> arr[i] <= arr[j]
	Sort(arr []int)
}

type InsertionSorter struct{}