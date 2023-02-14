package main

import (
	"fmt"
	"math"
)

func main(){
	fmt.Println(narcissistic(153))
	fmt.Println(narcissistic(111))
	fmt.Println(narcissistic(1634))
	fmt.Println(outlier(2, 4, 0, 100, 4, 11, 2602, 36))
	fmt.Println(outlier(160, 3, 1719, 19, 11, 13, -21))
	fmt.Println(outlier(11, 13, 15, 19, 9, 13, -21))
	i := []string{"red", "blue", "yellow", "black", "grey"}
	fmt.Println(haystack(i, "blue"))
	j := []int{1,5,5,5,5,3}
	fmt.Println(blueOcean(j, 5))
	k := []int{1,2,3,4,5,6,10}
	fmt.Println(blueOcean(k, 1))

}

func narcissistic(n int) bool{
	var length int
	var i = n
	var x = n
	for i > 0 {
		i = i/10
		length++
	}

	slc := []int{}
	for n > 0 {
		k := math.Pow(float64(n%10), float64(length))
		slc = append(slc, int(k))
		n = n / 10
	}
	
	m := 0

	for _, l := range slc{
		m = m+l
	}

	if m == x{
		return true
	}else{
		return false
	}
}

func outlier(n ...int) (int, bool){
	var odds = []int{}
	var even = []int{}

	for _, m := range n{
		if m%2 != 0 {
			odds = append(odds, m)
		}else{
			even = append(even, m)
		}
	}

	if len(odds)==1{
		return odds[0], true
	}else if len(even)==1{
		return even[0], true
	}else{
		return 0, false
	}
}

func haystack(i []string, j string) string{
	var h map[string]string
	h = map[string]string{}

	for _, k := range i{
		h[k] = k
	}

	var value, isExist = h[j]

	if isExist {
		return value
	} else {
		return "item is not exists"
	}
}

func blueOcean(i []int, j int) []int{
	var h map[int]int
	h = map[int]int{}

	for _, k := range i{
		h[k] = k
	}
	
	delete(h, j)

    l := []int{}
    for _, m := range h {
        l = append(l, m)
    }

	return l
}