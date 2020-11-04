package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	// "sort"
	// "math"
)

//N _
const (
	N = 2*1e5 + 5
)

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

func abs(a int64) int64 {
	if a > 0{
		return a
	}
	return 0 - a
}

func main(){
	rs := bufio.NewReader(os.Stdin)
	ws := bufio.NewWriter(os.Stdout)

	strs := make(chan *string, 100)

	go inputter(rs, strs)

	n , _ := strconv.ParseInt(*<-strs, 10, 64)

	a := make([]int64, n + 1)

	for i := int64(1); i < n + 1; i++ {
		x, _ := strconv.ParseInt(*<-strs, 10, 64)
		a[x] = i
	}

	tot := int64(0)

	for i := int64(1); i < n; i++ {
		tot += abs(a[i] - a[i + 1])
	}

	ws.WriteString(fmt.Sprint(tot))

	ws.Flush()
}


