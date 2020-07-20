package main

import (
	"fmt"
	"math/rand"
	"time"
	)







func bubble_sort(ulist [] int, usize int) {



	swapd := false

	for i := 0; i < usize; i++ {

		swapd = false

		for j := 0; j <usize -1; j++ {

			if ulist[j] > ulist[j+1] {

				ulist[j], ulist[j+1] = ulist[j+1], ulist[j]
				swapd = true

			}



		}
		if swapd == false {
			return 

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

	//fmt.Println(unsorted_slice)

	start_time := time.Now()

	bubble_sort(unsorted_slice, arr_size)

	duration := time.Since(start_time)

	//fmt.Println(unsorted_slice)

	fmt.Println("----------------Execution Time --->", duration)





}



/*

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


*/
