package utils

func Min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

//PrintMap generic method for printing ant map
//func PrintMap[K comparable, V any](m map[K]V) {
//	for k, v := range m {
//		fmt.Printf("%s: %d\n", k, v)
//	}
//}
