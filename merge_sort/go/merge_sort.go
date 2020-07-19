package main

import (
	"fmt"
	"time"
	"math/rand"
	)



func merge_sort(ulist []int) {


	if (len(ulist) > 1) {

		mid := len(ulist) / 2
		left :=  make([] int, mid)
		right := make([] int, len(ulist)-mid)

		copy(left, ulist[:mid])
		copy(right, ulist[mid:])

		merge_sort(left)
		merge_sort(right)

		i,j,k := 0,0,0

		for (i < len(left) && j < len(right)) {

			if left[i] < right[j] {

				ulist[k] = left[i]
				i = i + 1
				k = k + 1

			}else {

				ulist[k] = right[j]
				j = j + 1
				k = k + 1

			}

		}

		for i < len(left) {

			ulist[k] = left[i]
			i = i + 1
			k = k + 1

		}

		for j < len(right) {

			ulist[k] = right[j]
			j = j + 1
			k = k + 1

		}


	}

	//fmt.Println(ulist)
}



func main() {

        const arr_size = 10000000
        unsorted_arr := make( []int, arr_size, arr_size)

        rand.Seed(time.Now().UnixNano())

        for i:= 0; i < arr_size; i++ {
 
                unsorted_arr[i] =  rand.Intn(100)
                //fmt.Println(unsorted_arr[i])
        }


        //fmt.Println(unsorted_arr)

	start_time := time.Now()

        merge_sort(unsorted_arr)

	duration := time.Since(start_time)

	//fmt.Println(unsorted_arr)

	fmt.Println("--------------- Execution Time -->", duration)

}

