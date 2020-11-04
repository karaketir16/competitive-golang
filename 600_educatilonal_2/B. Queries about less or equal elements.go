package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"sort"
)

//N _
const (
	N = 2*1e5 + 5
)

func upperBound(array []int, target int) int {
	low, high, mid := 0, len(array)-1, 0

	for low <= high {
		mid = (low + high) / 2
		if array[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

func inputter(inputReader *bufio.Reader,strs chan *string){
	for {
		input , _ := inputReader.ReadString('\n')
		input = strings.TrimSpace(input)
		if len(input) > 0 {
			inputs := strings.Split(input, " ")
			for i := range inputs{
				strs <- &inputs[i]
			}
		}
	}
}

func main(){
	rs := bufio.NewReader(os.Stdin)
	ws := bufio.NewWriter(os.Stdout)

	strs := make(chan *string, 100)

	go inputter(rs, strs)

	n , _ := strconv.ParseInt(*<-strs, 10, 32)
	m , _ := strconv.ParseInt(*<-strs, 10, 32)

	a := make([]int, 0, N)

	for i := 0; i < int(n); i++ {
		x, _ := strconv.ParseInt(*<-strs, 10, 32)
		a = append(a, int(x))
	}

	sort.Ints(a)

	for i := 0; i < int(m); i++ {
		x, _ := strconv.ParseInt(*<-strs, 10, 32)
		ws.WriteString(fmt.Sprint(upperBound(a, int(x))) + " ")
	}
	
	ws.Flush()
}


