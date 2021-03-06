package main

import(
	"fmt"
	"math/rand"
	"time"
	)



func get_sort(sort_ar []int , ar_size int) {

	start :=  time.Now()

	for n:=1; n<ar_size; n++ {
		j := n
		for  j > 0 {
			if sort_ar[j] < sort_ar[j-1] {

				sort_ar[j], sort_ar[j-1] = sort_ar[j-1], sort_ar[j]
			}
			j = j -1
	           }
		}
	duration := time.Since(start)

	fmt.Println(sort_ar)
	fmt.Printf("\n\n ------------ Execution Time ---> %v \n", duration)
	}

func main() {

	const arr_size = 10000 
	unsorted_arr := make( []int, arr_size, arr_size)

	rand.Seed(time.Now().UnixNano())

	for i:= 0; i < arr_size; i++ {
 
		unsorted_arr[i] =  rand.Intn(100)
		//fmt.Println(unsorted_arr[i])
	}


	fmt.Println(unsorted_arr)

	get_sort(unsorted_arr, arr_size)

}
