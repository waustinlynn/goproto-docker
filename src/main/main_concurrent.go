package main

import (
	"fmt"
	"time"
	"math/rand"
	concurrent "../concurrent"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func syncSum(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func randArr(size int) []int {
	var ret []int
	for i := 0; i < size; i++ {
		ret = append(ret, rand.Intn(100))
	}	
	return ret
}

func viaChannel(s []int) int {
	multiS := splitArray(s, len(s) /4)

	c := make(chan int)

	for _, val := range multiS {
		go sum(val, c)
	}

	sum := 0
	for i := 0; i < len(multiS); i++ {
		sum += <- c
	}
	return sum
}

func main() {
	s := randArr(500000000)
	start := time.Now()
	asyncSum := viaChannel(s)
	fmt.Println(time.Since(start).Seconds())
	fmt.Println(asyncSum)
	start = time.Now()
	syncSum := syncSum(s)
	fmt.Println(time.Since(start).Seconds())
	fmt.Println(syncSum)
	concurrent.Say("something")
	
}

func splitArray(ar []int, by int) [][]int {
	soFar := 0
	listSize := len(ar)
	var newArr [][]int
	for soFar < listSize {
		newArr = append(newArr, ar[soFar:soFar + by])
		soFar += by
	}

	return newArr
}
