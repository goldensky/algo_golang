package main

func ChunkFunction(arr []int, chunkSize int) [][]int {
	resultSize := (len(arr) + chunkSize - 1) / chunkSize
	result := make([][]int, 0, resultSize)

	for i := 0; i < resultSize; i++ {
		start := i * chunkSize
		end := min(start+chunkSize, len(arr))
		result = append(result, arr[start:end])
	}

	return result // [][]int
}

//arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//result := ChunkFunction(arr, 3)
//fmt.Println(result)
