package arrays

func Contains(ns []int, n int) bool {
	for i := range ns {
		if ns[i] == n {
			return true
		}
	}
	return false
}
