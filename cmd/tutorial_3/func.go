package main

import "fmt"

func main() {
	// fixed length, same type
	var arr [3]int32
	fmt.Println(arr)
	arr[1] = 123
	fmt.Println(arr[0])
	fmt.Println(arr[1:3])

	// contiguous in memory,
	// doesn't need to store location of each index in memory
	// every 4 byte (int 32) for each index
	// & used to get memory location
	fmt.Println(&arr[0])
	fmt.Println(&arr[1])

	secArray := [3]int32{1, 2, 3}
	fmt.Println(secArray)

	//slices aka dynamic array
	slice := []int32{1, 2, 3}
	fmt.Printf("Old length is %v with capacity %v\n", len(slice), cap(slice))
	slice = append(slice, 4)
	fmt.Println(slice)
	fmt.Printf("New length is %v with capacity %v\n", len(slice), cap(slice))
	secSlice := []int32{5, 6}
	slice = append(slice, secSlice...)
	fmt.Printf("Newer length is %v with capacity %v\n", len(slice), cap(slice))

	// use make for maps, slices e.t.c
	thirdSlice := make([]int32, 3, 8) // 3 is len, 8 is cap
	fmt.Println(thirdSlice)

	dict := map[string]uint8{"a": 255}
	fmt.Println(dict["b"]) //map always returns something so be careful
	val, ok := dict["c"]
	if ok {
		fmt.Println(val)
	} else {
		fmt.Print("does not exist in map\n")
	}

	for str, num := range dict {
		fmt.Println(str, num)
	}

	for idx, val := range arr {
		fmt.Println(idx, val)
	}

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
}
