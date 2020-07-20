package main

import(
	"fmt"
	"math/rand"
	"time"
	)





func odd_even_sort(ulist [] int, usize int) {

	sorted := true


	for sorted == true {

		sorted = false
		for i := 0; i < usize-1; i += 2 {

			if ulist[i] > ulist[i+1] {

				ulist[i], ulist[i+1] = ulist[i+1], ulist[i]
				sorted = true
			}

		}
		for j := 1; j < usize-1; j += 2 {

			if ulist[j] > ulist[j+1] {
				ulist[j], ulist[j+1] = ulist[j+1], ulist[j]
				sorted = true

			}
		}
	}
}






func main() {

        arr_size := 100000

        unsorted_slice := make([] int, arr_size, arr_size)

        rand.Seed(time.Now().UnixNano())


        for i := 0; i < arr_size; i++ {

                unsorted_slice[i] = rand.Intn(100)

        }

        fmt.Println(unsorted_slice)

        start_time := time.Now()

        odd_even_sort(unsorted_slice, arr_size)

        duration := time.Since(start_time)

        fmt.Println(unsorted_slice)

        fmt.Println("----------------Execution Time --->", duration)





}




